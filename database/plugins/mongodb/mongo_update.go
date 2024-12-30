package mongodb

import (
	"context"

	"github.com/neghi14/starter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type mongoUpdate[T any] struct {
	filter bson.D
	client *mongo.Collection
}

func (mgo *mongoUpdate[T]) One(body T) *mongoUpdateOne[T] {
	return &mongoUpdateOne[T]{
		client: mgo.client,
		filter: mgo.filter,
		body:   body,
		parser: starter.NewParser(),
	}
}

func (mgo *mongoUpdate[T]) Many(body T) *mongoUpdateMany[T] {
	return &mongoUpdateMany[T]{
		client: mgo.client,
		filter: mgo.filter,
		body:   body,
		parser: starter.NewParser(),
	}
}

type mongoUpdateOne[T any] struct {
	client *mongo.Collection
	filter bson.D
	parser *starter.Parser
	body   T
}

func (mgo *mongoUpdateOne[T]) Exec(ctx context.Context) error {
	kv, err := mgo.parser.ParseToKeyValue(&mgo.body)
	if err != nil {
		return err
	}

	body, err := mgo.parser.ConvertToBson(kv)
	if err != nil {
		return err
	}

	_, err = mgo.client.UpdateOne(ctx, mgo.filter, bson.D{{Key: "$set", Value: body}})
	if err != nil {
		return err
	}
	return nil
}

type mongoUpdateMany[T any] struct {
	client *mongo.Collection
	filter bson.D
	parser *starter.Parser
	body   T
}

func (mgo *mongoUpdateMany[T]) Exec(ctx context.Context) error {
	kv, err := mgo.parser.ParseToKeyValue(&mgo.body)
	if err != nil {
		return err
	}

	body, err := mgo.parser.ConvertToBson(kv)
	if err != nil {
		return err
	}

	_, err = mgo.client.UpdateMany(ctx, mgo.filter, bson.D{{Key: "$set", Value: body}})
	if err != nil {
		return err
	}
	return nil
}
