package mysql

import (
	"fmt"
	"git.tenvine.cn/backend/gore/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	db *gorm.DB
)

func SetupGorm() error {
	var logLvl = logger.Silent
	switch log.GetLevel().String() {
	case "error":
		logLvl = logger.Error
	case "warn":
		logLvl = logger.Warn
	case "info":
		logLvl = logger.Info
	}
	newLogger := logger.New(log.StandardLogger(), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logLvl,
	})

	var err error
	db, err = gorm.Open(
		mysql.New(mysql.Config{Conn: Instance()}),
		&gorm.Config{Logger: newLogger},
	)
	if err != nil {
		return fmt.Errorf("mysql: %w", err)
	}
	return nil
}

func GormDB() *gorm.DB {
	return db
}
