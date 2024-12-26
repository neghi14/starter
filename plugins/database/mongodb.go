package database

import (
	"context"
	"time"

	"github.com/neghi14/starter"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mongoOptions struct {
	collection    string
	database_url  string
	database_name string
}

func NewMongoConfig() *mongoOptions {
	opts := &mongoOptions{}
	return opts
}

// SetCollection is used to set the/update the collection used
// by the mongodb plugin
func (mo *mongoOptions) SetCollection(col string) *mongoOptions {
	mo.collection = col
	return mo
}

// SetConnectionUrl sets the current url used for database connection
func (mo *mongoOptions) SetConnectionUrl(url string) *mongoOptions {
	mo.database_url = url
	return mo
}

// SetDatabaseName set the database name used by the current connection
func (mo *mongoOptions) SetDatabaseName(name string) *mongoOptions {
	mo.database_name = name
	return mo
}

func Mongo(cfg *mongoOptions) (*starter.DatabaseAdapter, error) {
	mo := new()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	client, err := mongo.Connect(options.Client().ApplyURI(cfg.database_url))
	if err != nil {
		return nil, err
	}
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}
	db := client.Database(cfg.database_name).Collection(cfg.collection)
	mo.getAttr()
	db.Indexes().CreateMany(ctx, []mongo.IndexModel{})
	return &starter.DatabaseAdapter{
		Name: "mongo-database",
		FindOne: func(ctx context.Context, filter, result interface{}) error {
			var data bson.D
			res := db.FindOne(ctx, filter)
			if err := res.Decode(&data); err != nil {
				return err
			}
			model, _ := mo.ConvertFromBson(data)
			return mo.parseToStruct(result, model)
		},
		Find: func(ctx context.Context, filter, result interface{}) error {
			var data []bson.D
			res, err := db.Find(ctx, filter)
			if err != nil {
				return err
			}
			if err := res.All(ctx, &data); err != nil {
				return err
			}
			return nil
		},
		Save: func(ctx context.Context, data interface{}) error {
			res, err := mo.parseToKeyValue(data)
			if err != nil {
				return err
			}
			input, err := mo.ConvertToBson(res)
			if err != nil {
				return err
			}
			_, err = db.InsertOne(ctx, input)
			if err != nil {
				return err
			}

			return nil
		},
		UpdateOne: func(ctx context.Context, filter, update interface{}) error {
			_, err := db.UpdateOne(ctx, filter, update)
			if err != nil {
				return err
			}
			return nil
		},
		Update: func(ctx context.Context, filter, update interface{}) error {
			_, err := db.UpdateMany(ctx, filter, update)
			if err != nil {
				return err
			}
			return nil
		},
		DeleteOne: func(ctx context.Context, filter interface{}) error {
			_, err := db.DeleteOne(ctx, filter)
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, filter interface{}) error {
			_, err := db.DeleteMany(ctx, filter)
			if err != nil {
				return err
			}
			return nil
		},
	}, nil
}
