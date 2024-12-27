package redisdb

import (
	"context"
	"time"

	"github.com/neghi14/starter/database"
	"github.com/redis/go-redis/v9"
)

type redisConf struct {
	database       int
	connection_url string
	ttl            time.Time
	table          string
}

func Opts() *redisConf {
	return &redisConf{}
}

func (r *redisConf) SetDatabase(db int) *redisConf {
	r.database = db
	return r
}

func (r *redisConf) SetConnectionUrl(url string) *redisConf {
	r.connection_url = url
	return r
}

func (r *redisConf) SetTTL(ttl time.Time) *redisConf {
	r.ttl = ttl
	return r
}

func (r *redisConf) SetTable(table string) *redisConf {
	r.table = table
	return r
}

func New[Model any](cfg *redisConf, model Model) (*database.DatabaseAdapter[Model], error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr: cfg.connection_url,
		DB:   cfg.database,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	return &database.DatabaseAdapter[Model]{
		Name: "redis-database",
		FindOne: func(ctx context.Context, filter database.Filter) (*Model, error) {
			return nil, nil
		},
		Find:      func(ctx context.Context, filter database.Filter) ([]*Model, error) { return nil, nil },
		Save:      func(ctx context.Context, data interface{}) error { return nil },
		UpdateOne: func(ctx context.Context, filter database.Filter, update Model) error { return nil },
		Update:    func(ctx context.Context, filter database.Filter, update []Model) error { return nil },
		DeleteOne: func(ctx context.Context, filter database.Filter) error { return nil },
		Delete:    func(ctx context.Context, filter database.Filter) error { return nil },
		Disconnet: func(ctx context.Context) error { return nil },
	}, nil
}
