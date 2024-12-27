package database

import "context"

type DatabaseAdapter[T any] struct {
	Name      string
	Count     func(ctx context.Context, filter Filter) (int64, error)
	FindOne   func(ctx context.Context, filter Filter) (*T, error)
	Find      func(ctx context.Context, filter Filter) ([]*T, error)
	Save      func(ctx context.Context, data interface{}) error
	UpdateOne func(ctx context.Context, filter Filter, update T) error
	Update    func(ctx context.Context, filter Filter, update []T) error
	DeleteOne func(ctx context.Context, filter Filter) error
	Delete    func(ctx context.Context, filter Filter) error
	Disconnet func(ctx context.Context) error
}

type DatabaseAdapterOptions struct {
}

type DatabaseConfig func(*DatabaseAdapterOptions)

type Filter struct {
	Param []Param
	Skip  int64
	Limit int64
	Sort  Param
}

type Param struct {
	Key   string
	Value interface{}
}

func Opts() *Filter {
	return &Filter{
		Sort: Param{
			Key:   "_id",
			Value: 1,
		},
	}
}

func (f *Filter) Params(param ...Param) *Filter {
	f.Param = append(f.Param, param...)
	return f
}

func (f *Filter) SortBy(sort Param) *Filter {
	f.Sort = sort
	return f
}

func (f *Filter) LimitBy(value int64) *Filter {
	f.Limit = value
	return f
}

func (f *Filter) SkipBy(value int64) *Filter {
	f.Skip = value
	return f
}
