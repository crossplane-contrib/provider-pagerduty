package clients

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

// configCache caches configured provider clients keyed by a hash of the
// provider configuration, evicting entries older than ttl.
type configCache[T any] struct {
	mu    sync.Mutex
	items map[string]cacheEntry[T]
	ttl   time.Duration
	now   func() time.Time
}

type cacheEntry[T any] struct {
	value     T
	createdAt time.Time
}

func newConfigCache[T any](ttl time.Duration) *configCache[T] {
	return &configCache[T]{items: map[string]cacheEntry[T]{}, ttl: ttl, now: time.Now}
}

func (c *configCache[T]) get(key string) (T, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	now := c.now()
	for k, e := range c.items {
		if now.Sub(e.createdAt) > c.ttl {
			delete(c.items, k)
		}
	}
	e, ok := c.items[key]
	return e.value, ok
}

func (c *configCache[T]) put(key string, v T) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = cacheEntry[T]{value: v, createdAt: c.now()}
}

func hashOf(b []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

type fwData struct {
	resourceData   any
	dataSourceData any
}

// cachingProvider wraps a Terraform Plugin Framework provider so Configure,
// which upjet calls on every reconcile, only builds a new PagerDuty client
// (and validates the credentials against the API) once per configuration.
type cachingProvider struct {
	provider.Provider
	cache *configCache[fwData]
}

func newCachingProvider(inner provider.Provider, ttl time.Duration) provider.Provider {
	return &cachingProvider{Provider: inner, cache: newConfigCache[fwData](ttl)}
}

func (p *cachingProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	h := hashOf([]byte(req.Config.Raw.String()))
	if d, ok := p.cache.get(h); ok {
		resp.ResourceData = d.resourceData
		resp.DataSourceData = d.dataSourceData
		return
	}

	p.Provider.Configure(ctx, req, resp)
	if resp.Diagnostics.HasError() || resp.ResourceData == nil {
		return
	}
	p.cache.put(h, fwData{resourceData: resp.ResourceData, dataSourceData: resp.DataSourceData})
}
