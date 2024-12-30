package mongodb

import (
	"context"

	"github.com/neghi14/starter"
	"github.com/neghi14/starter/database"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mongoFind[T any] struct {
	client *mongo.Collection
	filter bson.D
}

func (mgo *mongoFind[T]) One() *mongoFindOne[T] {
	return &mongoFindOne[T]{
		client: mgo.client,
		parser: starter.NewParser(),
		filter: mgo.filter,
	}
}

func (mgo *mongoFind[T]) Many() *mongoFindMany[T] {
	return &mongoFindMany[T]{
		filter: mgo.filter,
		parser: starter.NewParser(),
	}
}

type mongoFindOne[T any] struct {
	client *mongo.Collection
	parser *starter.Parser
	filter bson.D
	sort   bson.D
	column bson.D
}

func (mfo *mongoFindOne[T]) Sort(sortKeys ...database.Args) database.IFindOne[T] {
	parsedSortkey := bson.D{}
	for _, sk := range sortKeys {
		var orientation int
		val, ok := sk.Value().(database.SortKey)
		if !ok {
			panic("Invalid Value provided, expected Value of Type SortKey")
		}
		switch val {
		case database.ASC:
			orientation = -1
		case database.DESC:
			orientation = 1
		default:
			panic("Invalid Key Provided!")
		}

		parsedSortkey = append(parsedSortkey, bson.E{Key: sk.Key(), Value: orientation})
	}
	mfo.sort = parsedSortkey
	return mfo
}

func (mfo *mongoFindOne[T]) Column(column ...database.Args) database.IFindOne[T] {
	parsedColumns := bson.D{}
	for _, c := range column {
		parsedColumns = append(parsedColumns, bson.E{Key: c.Key(), Value: c.Value()})
	}
	mfo.column = parsedColumns
	return mfo
}

func (mfo *mongoFindOne[T]) Exec(ctx context.Context) (T, error) {
	var empty T
	var res T
	var single bson.D
	err := mfo.client.FindOne(ctx, mfo.filter, options.FindOne().
		SetSort(mfo.sort).SetProjection(mfo.column)).Decode(&single)
	if err != nil {
		return empty, err
	}
	d, err := mfo.parser.ConvertFromBson(single)
	if err != nil {
		return empty, err
	}
	if err = mfo.parser.ParseToStruct(&res, d); err != nil {
		return empty, err
	}
	return res, nil
}

type mongoFindMany[T any] struct {
	client *mongo.Collection
	parser *starter.Parser
	filter bson.D
	skip   int64
	limit  int64
	sort   bson.D
	column bson.D
}

func (mfo *mongoFindMany[T]) Sort(sortKeys ...database.Args) database.IFindMany[T] {
	parsedSortkey := bson.D{}
	for _, sk := range sortKeys {
		var orientation int
		val, ok := sk.Value().(database.SortKey)
		if !ok {
			panic("Invalid Value provided, expected Value of Type SortKey")
		}
		switch val {
		case database.ASC:
			orientation = -1
		case database.DESC:
			orientation = 1
		default:
			panic("Invalid Key Provided!")
		}

		parsedSortkey = append(parsedSortkey, bson.E{Key: sk.Key(), Value: orientation})
	}
	mfo.sort = parsedSortkey
	return mfo
}

func (mfo *mongoFindMany[T]) Column(column ...database.Args) database.IFindMany[T] {
	parsedColumns := bson.D{}
	for _, c := range column {
		parsedColumns = append(parsedColumns, bson.E{Key: c.Key(), Value: c.Value()})
	}
	mfo.column = parsedColumns
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

func (mfo *mongoFindMany[T]) Exec(ctx context.Context) ([]T, error) {
	var res []T
	result, err := mfo.client.Find(ctx, mfo.filter, options.Find().
		SetSort(mfo.sort).SetProjection(mfo.column).
		SetSkip(mfo.skip).SetLimit(mfo.limit))
	if err != nil {
		return nil, err
	}

	defer result.Close(ctx)
	for result.Next(ctx) {
		var single bson.D

		if err = result.Decode(&single); err != nil {
			return nil, err
		}
		d, err := mfo.parser.ConvertFromBson(single)
		if err != nil {
			return nil, err
		}
		var singleRes T
		if err = mfo.parser.ParseToStruct(&singleRes, d); err != nil {
			return nil, err
		}

		res = append(res, singleRes)
	}

	return res, nil
}
