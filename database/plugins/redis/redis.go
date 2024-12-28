package redisdb

import (
	"context"
	"time"

	"github.com/neghi14/starter/database"
	"github.com/neghi14/starter/internal/parser"
	"github.com/redis/go-redis/v9"
)

type redisConf struct {
	database       int
	connection_url string
	ttl            time.Time
	table          string
	parser         parser.Parser
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

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.connection_url,
		DB:   cfg.database,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}

	// Create Index
	if _, err := rdb.FTCreate(ctx, "idx:"+cfg.table, &redis.FTCreateOptions{
		OnJSON: true,
		Prefix: []interface{}{cfg.table + ":"},
	}, &redis.FieldSchema{}).Result(); err != nil {
		return nil, err
	}

	return &database.DatabaseAdapter[Model]{
		Name: "redis-database",
		FindOne: func(ctx context.Context, filter database.Filter) (*Model, error) {
			//TODO: create function that parses filter to redis query object
			_, err := rdb.FTSearch(ctx, "idx:"+cfg.table, "").Result()
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
		Find: func(ctx context.Context, filter database.Filter) ([]*Model, error) { return nil, nil },
		Save: func(ctx context.Context, data Model) error {
			kv, err := cfg.parser.ParseToKeyValue(&data)
			if err != nil {
				return err
			}
			input := map[string]interface{}{}

			for _, k := range kv {
				input[k.Key] = k.Value
			}

			_, err = rdb.JSONSet(ctx, createKey(cfg.table), "*", input).Result()
			if err != nil {
				return err
			}
			return nil
		},
		UpdateOne: func(ctx context.Context, filter database.Filter, update Model) error { return nil },
		Update:    func(ctx context.Context, filter database.Filter, update []Model) error { return nil },
		DeleteOne: func(ctx context.Context, filter database.Filter) error { return nil },
		Delete:    func(ctx context.Context, filter database.Filter) error { return nil },
		Disconnet: func(ctx context.Context) error { return nil },
	}, nil
}

func createKey(prefix string) string {
	return prefix
}
