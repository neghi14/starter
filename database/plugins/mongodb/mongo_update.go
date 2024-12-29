package mongodb

import (
	"context"
)

type mongoUpdate[T any] struct {
	filter []string
}

func (mgo *mongoUpdate[T]) One(body T) *mongoUpdateOne[T] {
	return &mongoUpdateOne[T]{
		filter: mgo.filter,
		body:   body,
	}
}

func (mgo *mongoUpdate[T]) Many(body T) *mongoUpdateMany[T] {
	return &mongoUpdateMany[T]{
		filter: mgo.filter,
		body:   body,
	}
}

type mongoUpdateOne[T any] struct {
	filter []string
	body   T
}

func (mgo *mongoUpdateOne[T]) Exec(ctx context.Context) (*T, error) {

	return nil, nil
}

type mongoUpdateMany[T any] struct {
	filter []string
	body   T
}

func (mgo *mongoUpdateMany[T]) Exec(ctx context.Context) (*T, error) {

	return nil, nil
}
