package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func DefaultCollector(name string) Collector {
	return CollectorFunc(func() *mongo.Collection {
		return Database().Collection(name)
	})
}

type Collector interface {
	Collection() *mongo.Collection
}

type CollectorFunc func() *mongo.Collection

func (fn CollectorFunc) Collection() *mongo.Collection {
	return fn()
}
