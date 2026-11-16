#!/usr/bin/env python3
"""Exercise tag resolution without the network or changes to the real Go module."""

import os
from pathlib import Path
import subprocess
import tempfile
import unittest


SCRIPT = Path(__file__).with_name("update-terraform-provider.sh").resolve()
MODULE = "github.com/PagerDuty/terraform-provider-pagerduty"


class UpdateTerraformProviderTest(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory()
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.repo = self.root / "upstream"
        self.repo.mkdir()
        self.git("init", "-q")
        self.git("-c", "user.name=Test", "-c", "user.email=test@example.com",
                 "commit", "-q", "--allow-empty", "-m", "fixture")
        self.sha = self.git("rev-parse", "HEAD").strip()
        self.log = self.root / "go.log"
        self.go = self.root / "go"
        self.go.write_text('#!/bin/sh\nprintf "%s\\n" "$*" >> "$GO_LOG"\n'
                           '[ "${FAIL_GET:-}" != 1 ] || exit 1\n')
        self.go.chmod(0o755)

    def git(self, *args):
        return subprocess.check_output(["git", "-C", str(self.repo), *args], text=True)

    def update(self, version="1.2.3", repo=None, **env):
        return subprocess.run(
            ["sh", str(SCRIPT), str(repo or self.repo), version, str(self.go)],
            env={**os.environ, "GO_LOG": str(self.log), **env},
            capture_output=True, text=True)

    def assert_go_calls(self):
        self.assertEqual(self.log.read_text().splitlines(),
                         [f"get {MODULE}@{self.sha}", "mod tidy"])

    def test_lightweight_tag(self):
        self.git("tag", "v1.2.3")
        self.assertEqual(self.update().returncode, 0)
        self.assert_go_calls()

    def test_annotated_tag(self):
        self.git("-c", "user.name=Test", "-c", "user.email=test@example.com",
                 "tag", "-a", "v1.2.3", "-m", "release")
        self.assertNotEqual(self.git("rev-parse", "v1.2.3").strip(), self.sha)
        self.assertEqual(self.update().returncode, 0)
        self.assert_go_calls()

    def test_missing_tag_does_not_run_go(self):
        self.assertNotEqual(self.update().returncode, 0)
        self.assertFalse(self.log.exists())

    def test_failed_lookup_does_not_run_go(self):
        self.assertNotEqual(self.update(repo=self.root / "missing").returncode, 0)
        self.assertFalse(self.log.exists())

    def test_failed_go_get_does_not_tidy(self):
        self.git("tag", "v1.2.3")
        self.assertNotEqual(self.update(FAIL_GET="1").returncode, 0)
        self.assertEqual(self.log.read_text().splitlines(), [f"get {MODULE}@{self.sha}"])

    def test_malformed_or_empty_lookup_does_not_run_go(self):
        git = self.root / "git"
        git.write_text('#!/bin/sh\nprintf "%s" "$REFS"\n')
        git.chmod(0o755)
        for refs in ("", "bad\trefs/tags/v1.2.3\n",
                     "a" * 40 + "\trefs/tags/v9.9.9\n"):
            with self.subTest(refs=refs):
                result = self.update(PATH=f"{self.root}:{os.environ['PATH']}", REFS=refs)
                self.assertNotEqual(result.returncode, 0)
                self.assertFalse(self.log.exists())

    def test_parallel_generation_waits_for_dependency_update(self):
        self.git("tag", "v1.2.3")
        root = SCRIPT.parent.parent
        harness = self.root / "generation.mk"
        harness.write_text(
            f'include {root}/Makefile\n'
            'go.generate:\n\t@echo generate >> "$$GO_LOG"\n')
        result = subprocess.run(
            ["make", "-j4", "-f", str(harness),
             "-o", "config/schema.json", "-o", "pull-docs", "-o", "copy-examples",
             "generate", f"MAKE=make -f {harness} -o config/schema.json -o pull-docs -o copy-examples",
             "NPROCS=2",
             f"GO={self.go}", f"TERRAFORM_PROVIDER_REPO={self.repo}",
             "TERRAFORM_PROVIDER_VERSION=1.2.3"],
            cwd=root, env={**os.environ, "GO_LOG": str(self.log)},
            capture_output=True, text=True)
        self.assertEqual(result.returncode, 0, result.stdout + result.stderr)
        calls = [line for line in self.log.read_text().splitlines()
                 if line.startswith(("get ", "mod tidy", "generate"))]
        self.assertEqual(calls, [f"get {MODULE}@{self.sha}", "mod tidy", "generate"])


if __name__ == "__main__":
    unittest.main()
