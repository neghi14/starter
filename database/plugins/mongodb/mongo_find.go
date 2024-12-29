package mongodb

import (
	"context"

	"github.com/neghi14/starter/database"
)

type mongoFind[T any] struct {
	filter []string
}

func (mgo *mongoFind[T]) One() *mongoFindOne[T] {
	return &mongoFindOne[T]{
		filter: mgo.filter,
	}
}

func (mgo *mongoFind[T]) Many() *mongoFindMany[T] {
	return &mongoFindMany[T]{
		filter: mgo.filter,
	}
}

type mongoFindOne[T any] struct {
	filter []string
	sort   []string
	column []string
}

func (mfo *mongoFindOne[T]) Sort(sort ...string) database.IFindOne[T] {
	mfo.sort = sort
	return mfo
}

func (mfo *mongoFindOne[T]) Column(column ...string) database.IFindOne[T] {
	mfo.column = column
	return mfo
}

func (mfo *mongoFindOne[T]) Exec(ctx context.Context) (*T, error) {
	return nil, nil
}

type mongoFindMany[T any] struct {
	filter []string
	skip   int64
	limit  int64
	sort   []string
	column []string
}

func (mfo *mongoFindMany[T]) Sort(sort ...string) database.IFindMany[T] {
	mfo.sort = sort
	return mfo
}

func (mfo *mongoFindMany[T]) Column(column ...string) database.IFindMany[T] {
	mfo.column = column
	return mfo
}

func (mfo *mongoFindMany[T]) Limit(limit int64) database.IFindMany[T] {
	mfo.limit = limit
	return mfo
}

func (mfo *mongoFindMany[T]) Skip(skip int64) database.IFindMany[T] {
	mfo.skip = skip
	return mfo
}

func (mfo *mongoFindMany[T]) Exec(ctx context.Context) (*T, error) {
	return nil, nil
}
