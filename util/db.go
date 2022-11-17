package util

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrDBNotFound = errors.New("db not found")
)

func MustDB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value("DB").(*gorm.DB)
	if !ok {
		panic(ErrDBNotFound)
	}
	return db
}
