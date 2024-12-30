package database

import (
	"context"
)

// db.model().find({"id": "sdfsg"}).one().many().limit(1).column("email", "name").sort("created_at", asc).skip(10).exec()
// db.model().save().exec()
// db.model().update({"id": "dsfgk"}).one(//body).many(//body).exec()
// db.model().delete({"id": "dsfgk"}).one().many().exec()

type Model[T any] interface {
	Find(filter ...Args) Find[T]
	Save(body T) ISave
	Update(filter ...Args) Update
	Delete(filter ...Args) Delete
}

type Find[T any] interface {
	One() IFindOne[T]
	Many() IFindMany[T]
}

type Update interface {
	One() IUpdateOne
	Many() IUpdateMany
}

type Delete interface {
	One() IDeleteOne
	Many() IDeleteMany
}

type IFindOne[T any] interface {
	Sort(sort ...Args) IFindOne[T]
	Column(column ...Args) IFindOne[T]
	Exec(ctx context.Context) (T, error)
}

type IFindMany[T any] interface {
	Sort(sort ...Args) IFindMany[T]
	Limit(limit int64) IFindMany[T]
	Column(column ...Args) IFindMany[T]
	Skip(skip int64) IFindMany[T]
	Exec(ctx context.Context) ([]T, error)
}

type ISave interface {
	Exec(ctx context.Context) error
}

type IUpdateOne interface {
	Exec(ctx context.Context) error
}

type IUpdateMany interface {
	Exec(ctx context.Context) error
}

type IDeleteOne interface {
	Exec(ctx context.Context) error
}

type IDeleteMany interface {
	Exec(ctx context.Context) error
}

type Args struct {
	key   string
	value interface{}
}

type SortKey int

const (
	ASC SortKey = iota
	DESC
)

var sortKeyMap = map[SortKey]string{
	ASC:  "asc",
	DESC: "desc",
}

// func SortKey(key string, value sortKey) Args {
// 	a := Args{key: key, value: value}
// 	return a
// }

func (s SortKey) String() string {
	return sortKeyMap[s]
}

func (a Args) String(key, value string) Args {
	a.key = key
	a.value = value
	return a
}
func (a Args) Int(key string, value int) Args {
	a.key = key
	a.value = value
	return a
}

func (a Args) Sort(key string, value SortKey) Args {
	a.key = key
	a.value = value
	return a
}

func (a *Args) Key() string {
	return a.key
}
func (a *Args) Value() interface{} {
	return a.value
}
