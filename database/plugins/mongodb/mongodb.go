package mongodb

import (
	"context"
	"time"

	"github.com/neghi14/starter/database"
	"github.com/neghi14/starter/internal"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type mongoConf struct {
	collection    string
	database_url  string
	database_name string
	parser        internal.Parser
	timestamp     bool
}

func NewMongoConfig() *mongoConf {
	opts := &mongoConf{}
	return opts
}

// SetCollection is used to set the/update the collection used
// by the mongodb plugin
func (mo *mongoConf) SetCollection(col string) *mongoConf {
	mo.collection = col
	return mo
}

// SetConnectionUrl sets the current url used for database connection
func (mo *mongoConf) SetConnectionUrl(url string) *mongoConf {
	mo.database_url = url
	return mo
}

// SetDatabaseName set the database name used by the current connection
func (mo *mongoConf) SetDatabaseName(name string) *mongoConf {
	mo.database_name = name
	return mo
}

func New[Model any](cfg *mongoConf, model Model) (*database.DatabaseAdapter[Model], error) {

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

	db.Indexes().CreateMany(ctx, []mongo.IndexModel{})
	return &database.DatabaseAdapter[Model]{
		Name: "mongo-database",
		Count: func(ctx context.Context, filter database.Filter) (int64, error) {
			var fil bson.D

			for _, f := range filter.Param {
				fil = append(fil, bson.E{Key: f.Key, Value: f.Value})
			}
			size, err := db.CountDocuments(ctx, fil)
			if err != nil {
				return 0, err
			}
			return size, nil
		},
		FindOne: func(ctx context.Context, filter database.Filter) (*Model, error) {
			var result Model
			var data bson.D
			var fil bson.D

			for _, f := range filter.Param {
				fil = append(fil, bson.E{Key: f.Key, Value: f.Value})
			}

			res := db.FindOne(ctx, fil)
			if err := res.Decode(&data); err != nil {
				return nil, err
			}

			model, _ := cfg.parser.ConvertFromBson(data)
			err = cfg.parser.ParseToStruct(&result, model)
			if err != nil {
				return nil, err
			}
			return &result, nil
		},
		Find: func(ctx context.Context, filter database.Filter) ([]*Model, error) {
			var fil bson.D
			var result []*Model

			for _, f := range filter.Param {
				fil = append(fil, bson.E{Key: f.Key, Value: f.Value})
			}
			sort := bson.D{{Key: filter.Sort.Key, Value: filter.Sort.Value}}
			res, err := db.Find(ctx, fil, options.Find().
				SetSort(sort).
				SetLimit(filter.Limit).
				SetSkip(filter.Skip))
			if err != nil {
				return nil, err
			}
			defer res.Close(ctx)
			for res.Next(ctx) {
				var single bson.D
				if err := res.Decode(&single); err != nil {
					return nil, err
				}
				b, err := cfg.parser.ConvertFromBson(single)
				if err != nil {
					return nil, err
				}
				var mo Model
				err = cfg.parser.ParseToStruct(&mo, b)
				if err != nil {
					return nil, err
				}
				result = append(result, &mo)
			}
			return result, nil
		},
		Save: func(ctx context.Context, data interface{}) error {
			res, err := cfg.parser.ParseToKeyValue(data)
			if err != nil {
				return err
			}
			input, err := cfg.parser.ConvertToBson(res)
			if err != nil {
				return err
			}
			_, err = db.InsertOne(ctx, input)
			if err != nil {
				return err
			}

			return nil
		},
		UpdateOne: func(ctx context.Context, filter database.Filter, update Model) error {
			var fil bson.D

			for _, f := range filter.Param {
				fil = append(fil, bson.E{Key: f.Key, Value: f.Value})
			}
			parsed, err := cfg.parser.ParseToKeyValue(update)
			if err != nil {
				return err
			}
			data, err := cfg.parser.ConvertToBson(parsed)
			if err != nil {
				return err
			}
			_, err = db.UpdateOne(ctx, fil, data)
			if err != nil {
				return err
			}
			return nil
		},
		Update: func(ctx context.Context, filter database.Filter, update []Model) error {
			parsed, err := cfg.parser.ParseToKeyValue(update)
			if err != nil {
				return err
			}
			data, err := cfg.parser.ConvertToBson(parsed)
			if err != nil {
				return err
			}
			_, err = db.UpdateMany(ctx, filter, bson.D{{Key: "$set", Value: data}})
			if err != nil {
				return err
			}
			return nil
		},
		DeleteOne: func(ctx context.Context, filter database.Filter) error {
			var fil bson.D
			for _, f := range filter.Param {
				fil = append(fil, bson.E{Key: f.Key, Value: f.Value})
			}
			_, err := db.DeleteOne(ctx, fil)
			if err != nil {
				return err
			}
			return nil
		},
		Delete: func(ctx context.Context, filter database.Filter) error {
			var fil bson.D
			for _, f := range filter.Param {
				fil = append(fil, bson.E{Key: f.Key, Value: f.Value})
			}
			_, err := db.DeleteMany(ctx, fil)
			if err != nil {
				return err
			}
			return nil
		},
	}, nil
}
