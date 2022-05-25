package util

import (
	"context"
	"errors"
	"git.tenvine.cn/backend/gore/constant"
	"gorm.io/gorm"
)

var (
	ErrDBNotFound = errors.New("db not found")
)

func MustDB(ctx context.Context) *gorm.DB {
	db, ok := ctx.Value(constant.ContextKeyDB).(*gorm.DB)
	if !ok {
		panic(ErrDBNotFound)
	}
	return db
}
