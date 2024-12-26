package database

import "context"

type DatabaseAdapter struct {
	Name      string
	FindOne   func(ctx context.Context, filter, result interface{}) error
	Find      func(ctx context.Context, filter, result interface{}) error
	Save      func(ctx context.Context, data interface{}) error
	UpdateOne func(ctx context.Context, filter, update interface{}) error
	Update    func(ctx context.Context, filter, update interface{}) error
	DeleteOne func(ctx context.Context, filter interface{}) error
	Delete    func(ctx context.Context, filter interface{}) error
	Disconnet func(ctx context.Context) error
}

type DatabaseAdapterOptions struct {
}

type DatabaseConfig func(*DatabaseAdapterOptions)
