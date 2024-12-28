package redisdb

import (
	"context"
	"time"

	"github.com/neghi14/starter/database"
	"github.com/neghi14/starter/internal/parser"
	"github.com/neghi14/starter/utils"
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
	index := []*redis.FieldSchema{}

	mo, err := cfg.parser.ParseToKeyValue(&model)
	if err != nil {
		return nil, err
	}

	for _, m := range mo {
		index = append(index, &redis.FieldSchema{
			FieldName: "$." + m.Key,
			As:        m.Key,
			FieldType: redis.SearchFieldTypeText,
		})
	}

	// Create Index
	if _, err := rdb.FTCreate(ctx, "idx:"+cfg.table, &redis.FTCreateOptions{
		OnJSON: true,
		Prefix: []interface{}{cfg.table + ":"},
	}, index...).Result(); err != nil {
		return nil, err
	}

	return &database.DatabaseAdapter[Model]{
		Name:  "redis-database",
		Count: func(ctx context.Context, filter database.Filter) (int64, error) { return 0, nil },
		FindOne: func(ctx context.Context, filter database.Filter) (*Model, error) {
			//TODO: create function that parses filter to redis query object
			_, err := rdb.FTSearch(ctx, "idx:"+cfg.table, "").Result()
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
		Find: func(ctx context.Context, filter database.Filter) ([]*Model, error) {
			_, err := rdb.FTSearch(ctx, "idx:"+cfg.table, "").Result()
			if err != nil {
				return nil, err
			}
			return nil, nil
		},
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
		UpdateOne: func(ctx context.Context, filter database.Filter, update Model) error {
			kv, err := cfg.parser.ParseToKeyValue(&update)
			if err != nil {
				return err
			}
			input := map[string]interface{}{}

			for _, k := range kv {
				input[k.Key] = k.Value
			}

			_, err = rdb.JSONSet(ctx, cfg.table+":"+filter.Param[0].Key, "*", input).Result()
			if err != nil {
				return err
			}
			return nil
		},
		Update: func(ctx context.Context, filter database.Filter, update Model) error {
			kv, err := cfg.parser.ParseToKeyValue(&update)
			if err != nil {
				return err
			}
			input := map[string]interface{}{}

			for _, k := range kv {
				input[k.Key] = k.Value
			}

			_, err = rdb.JSONSet(ctx, cfg.table+":"+filter.Param[0].Key, "*", input).Result()
			if err != nil {
				return err
			}
			return nil
		},
		DeleteOne: func(ctx context.Context, filter database.Filter) error {
			if _, err := rdb.JSONDel(ctx, cfg.table+":"+filter.Param[0].Key, "*").Result(); err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, filter database.Filter) error {
			if _, err := rdb.JSONDel(ctx, cfg.table+":"+filter.Param[0].Key, "*").Result(); err != nil {
				return err
			}
			return nil
		},
		Disconnet: func(ctx context.Context) error {
			return rdb.Close()
		},
	}, nil
}

func createKey(prefix string) string {
	return prefix + ":" + utils.Generate(12)
}