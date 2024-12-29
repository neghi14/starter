package database

import "context"

// db.model().find({"id": "sdfsg"}).one().many().limit(1).column("email", "name").sort("created_at", asc).skip(10).exec()
// db.model().save().exec()
// db.model().update({"id": "dsfgk"}).one(//body).many(//body).exec()
// db.model().delete({"id": "dsfgk"}).one().many().exec()

type Model[T any] interface {
	Find() Find[T]
	Save() ISave[T]
	Update() Update[T]
	Delete() Delete[T]
}

type Find[T any] interface {
	One() IFindOne[T]
	Many() IFindMany[T]
}

type Update[T any] interface {
	One() IUpdateOne[T]
	Many() IUpdateMany[T]
}

type Delete[T any] interface {
	One() IDeleteOne[T]
	Many() IDeleteMany[T]
}

type IFindOne[T any] interface {
	Sort(sort ...string) IFindOne[T]
	Column(column ...string) IFindOne[T]
	Exec[T]
}

type IFindMany[T any] interface {
	Sort(sort ...string) IFindMany[T]
	Limit(limit int64) IFindMany[T]
	Column(column ...string) IFindMany[T]
	Skip(skip int64) IFindMany[T]
	Exec[T]
}

type ISave[T any] interface {
	Exec[T]
}

type IUpdateOne[T any] interface {
	Exec[T]
}

type IUpdateMany[T any] interface {
	Exec[T]
}

type IDeleteOne[T any] interface {
	Exec[T]
}

type IDeleteMany[T any] interface {
	Exec[T]
}

type Exec[T any] interface {
	Exec(ctx context.Context) (*T, error)
}

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
