package singletonlist

import (
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/google/go-cmp/cmp"
)

const field = "config"

func TestListToObjectConvert(t *testing.T) {
	conv := listToObject{field}
	tests := map[string]struct {
		mode   config.Mode
		params map[string]any
		want   map[string]any
	}{
		"ToTerraformUnwrapsList": {
			mode:   config.ToTerraform,
			params: map[string]any{"name": "n", field: []any{map[string]any{"aggregate": "all"}}},
			want:   map[string]any{"name": "n", field: map[string]any{"aggregate": "all"}},
		},
		"ToTerraformEmptyListIsNull": {
			mode:   config.ToTerraform,
			params: map[string]any{field: []any{}},
			want:   map[string]any{field: nil},
		},
		"ToTerraformAbsentStaysAbsent": {
			mode:   config.ToTerraform,
			params: map[string]any{"name": "n"},
			want:   map[string]any{"name": "n"},
		},
		"FromTerraformWrapsObject": {
			mode:   config.FromTerraform,
			params: map[string]any{"id": "P1", field: map[string]any{"aggregate": "all"}},
			want:   map[string]any{"id": "P1", field: []any{map[string]any{"aggregate": "all"}}},
		},
		"FromTerraformNullStaysNull": {
			mode:   config.FromTerraform,
			params: map[string]any{field: nil},
			want:   map[string]any{field: nil},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := conv.Convert(tc.params, nil, tc.mode)
			if err != nil {
				t.Fatalf("Convert() error: %v", err)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("Convert() -want +got:\n%s", diff)
			}
		})
	}
}

func TestListToObjectRejectsMultiItemList(t *testing.T) {
	params := map[string]any{field: []any{map[string]any{}, map[string]any{}}}
	if _, err := (listToObject{field}).Convert(params, nil, config.ToTerraform); err == nil {
		t.Fatal("Convert() accepted a two-item list, want error")
	}
}
