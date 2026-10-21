package clients

import (
	"context"
	"testing"
	"time"

	"github.com/PagerDuty/terraform-provider-pagerduty/pagerduty"
)

// Configuring the SDKv2 provider must happen once per credential set: the
// returned client config is reused across reconciles until the TTL expires,
// and different credentials never share a client.
func TestConfigureSDKCachesPerConfiguration(t *testing.T) {
	sdk := pagerduty.Provider(pagerduty.IsMuxed)
	cache := newConfigCache[any](time.Minute)
	now := time.Unix(0, 0)
	cache.now = func() time.Time { return now }

	cfgA := map[string]any{keyToken: "token-a", "skip_credentials_validation": true}
	cfgB := map[string]any{keyToken: "token-b", "skip_credentials_validation": true}

	a1, err := configureSDK(context.Background(), sdk, cache, cfgA)
	if err != nil {
		t.Fatalf("configureSDK(a): %v", err)
	}
	a2, _ := configureSDK(context.Background(), sdk, cache, cfgA)
	if a1 != a2 {
		t.Error("same configuration: want the cached client config, got a new one")
	}

	b, _ := configureSDK(context.Background(), sdk, cache, cfgB)
	if b == a1 {
		t.Error("different credentials: got the client config of another credential set")
	}
	if got := b.(*pagerduty.Config).Token; got != "token-b" {
		t.Errorf("client config token: got %q, want %q", got, "token-b")
	}

	now = now.Add(2 * time.Minute)
	a3, _ := configureSDK(context.Background(), sdk, cache, cfgA)
	if a3 == a1 {
		t.Error("after TTL: want a freshly configured client config, got the expired one")
	}
}
