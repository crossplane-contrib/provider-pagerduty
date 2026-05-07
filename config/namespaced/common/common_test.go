package commmon

import (
	"context"
	"testing"
)

func TestGetFakeID(t *testing.T) {
	tests := map[string]struct {
		externalName string
		want         string
	}{
		"empty external name uses pagerduty-shaped sentinel": {
			externalName: "",
			want:         pagerDutyIDShapedSentinel,
		},
		"existing external name is unchanged": {
			externalName: "PABC123",
			want:         "PABC123",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetFakeID(context.Background(), tt.externalName, nil, nil)
			if err != nil {
				t.Fatalf("GetFakeID returned error: %v", err)
			}
			if got != tt.want {
				t.Fatalf("GetFakeID() = %q, want %q", got, tt.want)
			}
		})
	}
}
