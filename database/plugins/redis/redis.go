package redis

import (
	"context"

	"github.com/neghi14/starter/database"
)

type redisConf struct{}

func New[Model any](cfg *redisConf, model Model) *database.DatabaseAdapter[Model] {

	return &database.DatabaseAdapter[Model]{
		Name: "redis-database",
		FindOne: func(ctx context.Context, filter database.Filter, result Model) error {
			return nil
		},
		Find:      func(ctx context.Context, filter database.Filter, result []Model) error { return nil },
		Save:      func(ctx context.Context, data interface{}) error { return nil },
		UpdateOne: func(ctx context.Context, filter database.Filter, update Model) error { return nil },
		Update:    func(ctx context.Context, filter database.Filter, update []Model) error { return nil },
		DeleteOne: func(ctx context.Context, filter database.Filter) error { return nil },
		Delete:    func(ctx context.Context, filter database.Filter) error { return nil },
		Disconnet: func(ctx context.Context) error { return nil },
	}
}
