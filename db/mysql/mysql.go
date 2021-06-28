package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

var sdb *sql.DB

type Config struct {
	Enable bool

	ConnMaxLifeTime time.Duration `yaml:"conn-max-life-time"`
	MaxOpenConns    int           `yaml:"max-open-conns"`
	MaxIdleConns    int           `yaml:"max-idle-conns"`

	// DSN(Data Source Name)
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value&...&paramN=valueN]
	// Example:
	//     username:password@protocol(address)/dbname?param=value
	Dsn struct {
		Username string
		Password string
		// `tcp` or `unix`
		Protocol string
		// Example:
		//     localhost:1111
		Address string
		Dbname  string
		// Example:
		//     ?charset=UTF8&loc=UTC
		Params string
	}
}

func GetInstance() *sql.DB {
	return sdb
}

func Setup(cfg Config) error {
	if !cfg.Enable {
		return nil
	}

	var err error

	driverName := "mysql"
	sdb, err = sql.Open(driverName, getDSN(cfg))
	if err != nil {
		return err
	}

	sdb.SetConnMaxLifetime(cfg.ConnMaxLifeTime)
	sdb.SetMaxOpenConns(cfg.MaxOpenConns)
	sdb.SetMaxIdleConns(cfg.MaxIdleConns)

	return nil
}

func getDSN(config Config) string {
	sb := new(strings.Builder)
	if config.Dsn.Username != "" {
		sb.WriteString(config.Dsn.Username)
		if config.Dsn.Password != "" {
			sb.WriteString(":")
			sb.WriteString(config.Dsn.Password)
		}
		sb.WriteString("@")
	}
	if config.Dsn.Protocol != "" {
		sb.WriteString(config.Dsn.Protocol)
		if config.Dsn.Address != "" {
			sb.WriteString("(")
			sb.WriteString(config.Dsn.Address)
			sb.WriteString(")")
		}
	}
	sb.WriteString("/")
	sb.WriteString(config.Dsn.Dbname)
	if config.Dsn.Params != "" {
		sb.WriteString(config.Dsn.Params)
	}

	return sb.String()
}
