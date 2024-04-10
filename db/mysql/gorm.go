package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var (
	db *gorm.DB
)

func SetupGorm() error {
	var logLvl = logger.Silent
	newLogger := logger.New(log.Default(), logger.Config{
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
