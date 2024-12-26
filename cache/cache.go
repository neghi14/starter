package cache

import "context"

type CacheAdapter struct {
	Name string
	Get  func(ctx context.Context, key string, data interface{}) error
	Set  func(ctx context.Context, key string, data interface{}) error
}
