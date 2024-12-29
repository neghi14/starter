package mongodb

import (
	"context"
)

type mongoSave[T any] struct {
	body T
}

func (mso *mongoSave[T]) Exec(ctx context.Context) (*T, error) {
	return nil, nil
}
