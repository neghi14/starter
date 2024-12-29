package redisdb

import (
	"reflect"
	"testing"

	"github.com/neghi14/starter/database"
	"github.com/neghi14/starter/internal/parser"
	"github.com/redis/go-redis/v9"
)

func TestNew(t *testing.T) {
	type UserModel struct{}

	type args[T any] struct {
		cfg   *redisConf
		model T
	}
	tests := []struct {
		name    string
		args    args[UserModel]
		want    *database.DatabaseAdapter[UserModel]
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.cfg, tt.args.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOpts(t *testing.T) {
	tests := []struct {
		name string
		want *redisConf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Opts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Opts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisConf_SetDatabase(t *testing.T) {
	type args struct {
		db int
	}
	tests := []struct {
		name string
		r    *redisConf
		args args
		want *redisConf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetDatabase(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redisConf.SetDatabase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisConf_SetConnectionUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		r    *redisConf
		args args
		want *redisConf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetConnectionUrl(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redisConf.SetConnectionUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisConf_SetTTL(t *testing.T) {
	type args struct {
		ttl int64
	}
	tests := []struct {
		name string
		r    *redisConf
		args args
		want *redisConf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetTTL(tt.args.ttl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redisConf.SetTTL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisConf_SetTable(t *testing.T) {
	type args struct {
		table string
	}
	tests := []struct {
		name string
		r    *redisConf
		args args
		want *redisConf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.SetTable(tt.args.table); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("redisConf.SetTable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createKey(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createKey(tt.args.prefix); got != tt.want {
				t.Errorf("createKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRedisSearchKeyType(t *testing.T) {
	type args struct {
		parserType parser.ParserValueType
	}
	tests := []struct {
		name string
		args args
		want redis.SearchFieldType
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRedisSearchKeyType(tt.args.parserType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRedisSearchKeyType() = %v, want %v", got, tt.want)
			}
		})
	}
}
