package mongodb

import (
	"context"

	"github.com/neghi14/starter"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type mongoSave[T any] struct {
	client *mongo.Collection
	body   T
	parser *starter.Parser
}

func (mso *mongoSave[T]) Exec(ctx context.Context) error {
	kv, err := mso.parser.ParseToKeyValue(&mso.body)
	if err != nil {
		return err
	}

	body, err := mso.parser.ConvertToBson(kv)
	if err != nil {
		return err
	}

	_, err = mso.client.InsertOne(ctx, body)
	if err != nil {
		return err
	}
	return nil
}
