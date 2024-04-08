package mysql

import (
	"database/sql"
	"github.com/cstcen/gore/gonfig"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

var sdb *sql.DB

type Config struct {
	Enable bool

	ConnMaxLifeTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int

	// DSN(Data Source Name)
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value&...&paramN=valueN]
	// Example:
	//     username:password@protocol(address)/dbname?param=value
	Dsn DataSourceName
}

type DataSourceName struct {
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

func Instance() *sql.DB {
	return sdb
}

func SetupDefault() error {
	cfg := DefaultConfig()
	return Setup(cfg)
}

func Setup(cfg *Config) error {

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

func DefaultConfig() *Config {
	viper := gonfig.Instance()
	cfg := &Config{
		Enable:          viper.GetBool("gore.mysql.enable"),
		ConnMaxLifeTime: viper.GetDuration("gore.mysql.connMaxLifeTime"),
		MaxOpenConns:    viper.GetInt("gore.mysql.maxOpenConns"),
		MaxIdleConns:    viper.GetInt("gore.mysql.maxIdleConns"),
		Dsn: DataSourceName{
			Username: viper.GetString("gore.mysql.dsn.username"),
			Password: viper.GetString("gore.mysql.dsn.password"),
			Protocol: viper.GetString("gore.mysql.dsn.protocol"),
			Address:  viper.GetString("gore.mysql.dsn.address"),
			Dbname:   viper.GetString("gore.mysql.dsn.dbname"),
			Params:   viper.GetString("gore.mysql.dsn.params"),
		},
	}
	return cfg
}

func getDSN(config *Config) string {
	sb := new(strings.Builder)
	if len(config.Dsn.Username) > 0 {
		sb.WriteString(config.Dsn.Username)
		if len(config.Dsn.Password) > 0 {
			sb.WriteString(":")
			sb.WriteString(config.Dsn.Password)
		}
		sb.WriteString("@")
	}
	if len(config.Dsn.Protocol) > 0 {
		sb.WriteString(config.Dsn.Protocol)
		if len(config.Dsn.Address) > 0 {
			sb.WriteString("(")
			sb.WriteString(config.Dsn.Address)
			sb.WriteString(")")
		}
	}
	sb.WriteString("/")
	sb.WriteString(config.Dsn.Dbname)
	if len(config.Dsn.Params) > 0 {
		sb.WriteString(config.Dsn.Params)
	}

	return sb.String()
}
