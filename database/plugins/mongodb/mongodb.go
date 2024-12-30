package mongodb

import (
	"context"
	"sync"
	"time"

	"github.com/neghi14/starter"
	"github.com/neghi14/starter/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var mongo_sync_instance *sync.Once
var mongo_instance *mongo.Database

type MongoDB struct {
	*mongo.Database
}

func New(db_url string, db string) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if mongo_instance == nil {
		var client *mongo.Client
		var err error
		mongo_sync_instance.Do(func() {
			clientOpts := options.Client().ApplyURI(db_url)
			client, err = mongo.Connect(clientOpts)
			if err != nil {
				panic(err)
			}
		})
		if err = client.Ping(ctx, nil); err != nil {
			return nil, err
		}
		mongo_instance = client.Database(db)
	}

	return &MongoDB{mongo_instance}, nil
}

type MongoModel[T any] struct {
	client *mongo.Collection
	database.Model[T]
}

func Model[T any](db *MongoDB, model T, collection string) *MongoModel[T] {
	//set indexes if none

	//connect to collection
	client := db.Collection(collection)
	return &MongoModel[T]{
		client: client,
	}
}

func (mm *MongoModel[T]) Find(filter ...database.Args) *mongoFind[T] {
	parsedFilter := bson.D{}
	for _, fil := range filter {
		parsedFilter = append(parsedFilter, bson.E{Key: fil.Key(), Value: fil.Value()})
	}
	return &mongoFind[T]{
		filter: parsedFilter,
		client: mm.client,
	}
}

func (mm *MongoModel[T]) Save(body T) *mongoSave[T] {
	return &mongoSave[T]{
		body:   body,
		parser: starter.NewParser(),
	}
}

func (mm *MongoModel[T]) Update(filter ...database.Args) *mongoUpdate[T] {
	parsedFilter := bson.D{}
	for _, fil := range filter {
		parsedFilter = append(parsedFilter, bson.E{Key: fil.Key(), Value: fil.Value()})
	}
	return &mongoUpdate[T]{
		filter: parsedFilter,
		client: mm.client,
	}
}

func (mm *MongoModel[T]) Delete(filter ...database.Args) *mongoUpdate[T] {
	parsedFilter := bson.D{}
	for _, fil := range filter {
		parsedFilter = append(parsedFilter, bson.E{Key: fil.Key(), Value: fil.Value()})
	}
	return &mongoUpdate[T]{
		filter: parsedFilter,
	}
}
