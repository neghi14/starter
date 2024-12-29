package mongodb

import "context"

type MongoDB struct{}

type MongoModel[T any] struct {
}

func New() *MongoDB {
	return &MongoDB{}
}

func Model[T any](db *MongoDB, model T) *MongoModel[T] {
	return &MongoModel[T]{}
}

func (mm *MongoModel[T]) Find(filter ...string) *mongoFind[T] {
	return &mongoFind[T]{
		filter: filter,
	}
}

func (mm *MongoModel[T]) Save(body T) *mongoSave[T] {
	return &mongoSave[T]{
		body: body,
	}
}

func (mm *MongoModel[T]) Update(filter ...string) *mongoUpdate[T] {
	return &mongoUpdate[T]{
		filter: filter,
	}
}

func main() {

	var db = New()
	var m = Model(db, map[string]string{})

	m.Find().One().Column().Exec(context.Background())
}
