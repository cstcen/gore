package mongo

import (
	"context"
	"github.com/cstcen/gore/common"
	"github.com/cstcen/gore/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

type BaseRepo[ID any, T any] struct {
	Collector
}

func NewBaseRepo[ID any, T any](collector Collector) *BaseRepo[ID, T] {
	return &BaseRepo[ID, T]{Collector: collector}
}

func NewBaseRepoByCollectionName[ID any, T any](name string) *BaseRepo[ID, T] {
	return NewBaseRepo[ID, T](DefaultCollector(name))
}

func (r *BaseRepo[ID, T]) FindOneById(ctx context.Context, id ID) (*T, error) {
	var result T
	if err := r.Collection().FindOne(ctx, bson.D{WithFilterID(id)()}).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BaseRepo[ID, T]) FindOne(ctx context.Context, filter bson.D, opts ...*options.FindOneOptions) (*T, error) {
	var result T
	if err := r.Collection().FindOne(ctx, filter, opts...).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BaseRepo[ID, T]) FindOneConditions(ctx context.Context, id ID, filters ...bson.E) (*T, error) {
	var result T
	filter := bson.D{{"_id", id}}
	if len(filters) > 0 {
		filter = append(filter, filters...)
	}
	if err := r.Collection().FindOne(ctx, filter).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *BaseRepo[ID, T]) FindMany(ctx context.Context, ids []ID) ([]T, error) {
	result := make([]T, 0)
	cursor, err := r.Collection().Find(ctx, bson.D{WithFilterID(bson.D{{"$in", ids}})()})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepo[ID, T]) Find(ctx context.Context, filter bson.D, opts ...*options.FindOptions) ([]T, error) {
	result := make([]T, 0)
	cursor, err := r.Collection().Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepo[ID, T]) FindPaging(ctx context.Context, filter bson.D, pageData *common.PageData[T]) (*common.PageData[T], error) {
	if pageData == nil {
		pageData = common.NewPageData([]T{})
	}
	page := pageData.PageNo
	size := pageData.PageSize

	opts := options.Find().SetSkip(int64(size * (page - 1))).SetLimit(int64(size))
	if orderBys := strings.Split(pageData.OrderBy, ","); len(orderBys) > 0 {
		order := 1
		if pageData.Order == common.DESC {
			order = -1
		}
		sort := bson.D{}
		for _, orderBy := range strings.Split(pageData.OrderBy, ",") {
			sort = append(sort, bson.E{Key: orderBy, Value: order})
		}
		opts.SetSort(sort)
	}
	cursor, err := r.Collection().Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	result := make([]T, 0)
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	pageData.List = result
	if pageData.PageSize > len(result) {
		pageData.PageSize = len(result)
	}

	countDocuments, err := r.Collection().CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	return pageData.WithTotal(int(countDocuments)), nil
}

func (r *BaseRepo[ID, T]) FindAll(ctx context.Context) ([]T, error) {
	result := make([]T, 0)
	cursor, err := r.Collection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *BaseRepo[ID, T]) InsertOne(ctx context.Context, doc *T) error {
	result, err := r.Collection().InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	log.InfoCf(ctx, "inserted %T with ID %v", doc, result.InsertedID)
	return nil
}

func (r *BaseRepo[ID, T]) InsertMany(ctx context.Context, docs []*T) error {
	anies := make([]any, 0, len(docs))
	for _, doc := range docs {
		d := *doc
		anies = append(anies, &d)
	}

	result, err := r.Collection().InsertMany(ctx, anies)
	if err != nil {
		return err
	}
	log.InfoCf(ctx, "inserted %T with IDs %v", docs, result.InsertedIDs)
	return nil
}

func (r *BaseRepo[ID, T]) UpdateByIDs(ctx context.Context, ids []ID, update bson.D) error {
	result, err := r.Collection().UpdateMany(ctx, bson.D{WithFilterID(bson.D{{"$in", ids}})()}, update)
	if err != nil {
		log.WarningCf(ctx, "failed to modified documents, err: %s", err)
		return err
	}
	log.InfoCf(ctx, "modified %v documents", result.ModifiedCount)
	return nil
}

func (r *BaseRepo[ID, T]) UpsertOne(ctx context.Context, filter any, update any) error {
	opts := options.Update().SetUpsert(true)
	result, err := r.Collection().UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}
	log.InfoCf(ctx, "upserted document %v", result.UpsertedID)
	return nil
}

func (r *BaseRepo[ID, T]) UpsertByID(ctx context.Context, id ID, update any) error {
	opts := options.Update().SetUpsert(true)
	result, err := r.Collection().UpdateByID(ctx, id, update, opts)
	if err != nil {
		return err
	}
	log.InfoCf(ctx, "upserted document with ID %v", result.UpsertedID)
	return nil
}

func (r *BaseRepo[ID, T]) UpdateMany(ctx context.Context, filter bson.D, update bson.D) error {
	result, err := r.Collection().UpdateMany(ctx, filter, update)
	if err != nil {
		log.WarningCf(ctx, "failed to modified documents, err: %s", err)
		return err
	}
	log.InfoCf(ctx, "modified %v documents", result.ModifiedCount)
	return nil
}

func (r *BaseRepo[ID, T]) UpsertMany(ctx context.Context, filter bson.D, update bson.D) error {
	opts := options.Update().SetUpsert(true)
	result, err := r.Collection().UpdateMany(ctx, filter, update, opts)
	if err != nil {
		log.WarningCf(ctx, "failed to upserted documents, err: %s", err)
		return err
	}
	log.InfoCf(ctx, "upserted %v documents", result.UpsertedCount)
	return nil
}

func (r *BaseRepo[ID, T]) DeleteMany(ctx context.Context, ids []ID) error {
	deleteResult, err := r.Collection().DeleteMany(ctx, bson.D{WithFilterID(bson.D{{"$in", ids}})()})
	if err != nil {
		log.WarningCf(ctx, "failed to delete documents, err: %s", err)
		return err
	}
	log.InfoCf(ctx, "deleted %v documents", deleteResult.DeletedCount)
	return nil
}

func (r *BaseRepo[ID, T]) DeleteOne(ctx context.Context, id ID) error {
	deleteResult, err := r.Collection().DeleteOne(ctx, bson.D{WithFilterID(id)()})
	if err != nil {
		log.WarningCf(ctx, "failed to delete document, err: %s", err)
		return err
	}
	log.InfoCf(ctx, "deleted %v document", deleteResult.DeletedCount)
	return nil
}

func (r *BaseRepo[ID, T]) CountDocuments(ctx context.Context, filter any) (int64, error) {
	return r.Collection().CountDocuments(ctx, filter)
}
