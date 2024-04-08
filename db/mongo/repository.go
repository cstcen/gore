package mongo

import (
	"context"
	"github.com/cstcen/gore/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository[ID any, T any] interface {
	Collector
	Find(ctx context.Context, filter bson.D, opts ...*options.FindOptions) ([]T, error)
	FindOne(ctx context.Context, filter bson.D, opts ...*options.FindOneOptions) (*T, error)
	FindOneConditions(ctx context.Context, id ID, filters ...bson.E) (*T, error)
	FindMany(ctx context.Context, ids []ID) ([]T, error)
	FindAll(ctx context.Context) ([]T, error)
	InsertOne(ctx context.Context, doc *T) error
	InsertMany(ctx context.Context, docs []*T) error
	UpdateByIDs(ctx context.Context, ids []ID, update bson.D) error
	UpdateMany(ctx context.Context, filter bson.D, update bson.D) error
	UpsertOne(ctx context.Context, filter any, update any) error
	UpsertByID(ctx context.Context, id ID, update any) error
	UpsertMany(ctx context.Context, filter bson.D, update bson.D) error
	DeleteMany(ctx context.Context, ids []ID) error
	DeleteOne(ctx context.Context, id ID) error
	CountDocuments(ctx context.Context, filter any) (int64, error)
	FindPaging(ctx context.Context, filter bson.D, pageData *common.PageData[T]) (*common.PageData[T], error)
}

type RepoOptions func() bson.E

func WithFilter(key string, val any) RepoOptions {
	return func() bson.E {
		return bson.E{Key: key, Value: val}
	}
}

func WithFilterID(val any) RepoOptions {
	return func() bson.E {
		return bson.E{Key: "_id", Value: val}
	}
}
