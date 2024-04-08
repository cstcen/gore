package mysql

import (
	"fmt"
	"github.com/cstcen/gore/log"
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
	switch log.GetLevel() {
	case log.LevelError:
		logLvl = logger.Error
	case log.LevelWarning:
		logLvl = logger.Warn
	case log.LevelInfo, log.LevelDebug:
		logLvl = logger.Info
	}
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
