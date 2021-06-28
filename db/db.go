package db

import (
	"context"
	"database/sql"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	"git.tenvine.cn/backend/gore/db/es"
	goreMongo "git.tenvine.cn/backend/gore/db/mongo"
	goreMysql "git.tenvine.cn/backend/gore/db/mysql"
	"github.com/go-redis/cache/v8"
	"github.com/olivere/elastic"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	Cc   goreCache.Config
	Es   es.Config
	Mgo  goreMongo.Config
	Msql goreMysql.Config
}

type CheckResult struct {
	Elasticsearch ElasticsearchResult
	Cache         CacheResult
	Mongo         MongoResult
	Mysql         MysqlResult
}

type ElasticsearchResult struct {
	Info *elastic.PingResult
	Code int
	Err  error `json:"err,omitempty"`
}

type CacheResult struct {
	Stats *cache.Stats
}

type MongoResult struct {
	Err error `json:"err,omitempty"`
}

type MysqlResult struct {
	Stats sql.DBStats
	Err   error `json:"err,omitempty"`
}

func Check(cfg Config) *CheckResult {
	result := new(CheckResult)

	ccCli := goreCache.GetInstance()
	if ccCli != nil {
		result.Cache.Stats = ccCli.Stats()
	}

	esCli := es.GetInstance()
	if esCli != nil {
		info, code, err := esCli.Ping(cfg.Es.URL).Do(context.Background())
		result.Elasticsearch.Info = info
		result.Elasticsearch.Code = code
		result.Elasticsearch.Err = err
	}

	mgCli := goreMongo.GetInstance()
	if mgCli != nil {
		result.Mongo.Err = mgCli.Ping(context.Background(), readpref.Primary())
	}

	msCli := goreMysql.GetInstance()
	if msCli != nil {
		result.Mysql.Stats = msCli.Stats()
		err := msCli.Ping()
		if err != nil {
			result.Mysql.Err = err
		}
	}

	return result
}
