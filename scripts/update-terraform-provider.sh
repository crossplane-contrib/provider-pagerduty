#!/bin/sh
set -eu

repo=${1:?Terraform provider repository is required}
version=${2:?Terraform provider version is required}
go=${3:-go}
tag="refs/tags/v${version}"

# Keep ls-remote separate from the pipeline so a failed lookup cannot be hidden.
refs=$(git ls-remote --exit-code "$repo" "$tag" "${tag}^{}")
# Annotated tags need the peeled commit; lightweight tags point to it directly.
sha=$(printf '%s\n' "$refs" | awk -v tag="$tag" '
    $2 == tag { direct = $1 }
    $2 == tag "^{}" { peeled = $1 }
    END { print (peeled != "" ? peeled : direct) }
')
case "$sha" in
    ''|*[!0-9a-fA-F]*) echo "Cannot resolve ${tag} to a commit" >&2; exit 1 ;;
esac
if [ "${#sha}" -ne 40 ]; then
    echo "Invalid commit for ${tag}: ${sha}" >&2
    exit 1
fi

# Never pass an empty revision to Go: that would request latest instead.
"$go" get "github.com/PagerDuty/terraform-provider-pagerduty@${sha}"
"$go" mod tidy
