package gin

import (
	"git.tenvine.cn/backend/gore/db"
	goreCache "git.tenvine.cn/backend/gore/db/cache"
	goreEs "git.tenvine.cn/backend/gore/db/es"
	goreMongo "git.tenvine.cn/backend/gore/db/mongo"
	goreMysql "git.tenvine.cn/backend/gore/db/mysql"
	goreRedis "git.tenvine.cn/backend/gore/db/redis"
	"git.tenvine.cn/backend/gore/vo"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

const (
	RelativePathHealthCheck = "/healthcheck"
)

func CheckDB() *db.CheckResult {
	return db.Check(db.Config{
		Cache:         goreCache.NewConfig(),
		Elasticsearch: goreEs.NewConfig(),
		Mongo:         goreMongo.NewConfig(),
		Mysql:         goreMysql.NewConfig(),
		Redis:         goreRedis.NewConfig(),
	})
}

// Group struct
type Group struct {
	r *gin.Engine
}

func (g *Group) healthcheck() {
	g.r.GET(RelativePathHealthCheck, func(c *gin.Context) {
		c.JSON(http.StatusOK, CheckDB())
	})
}

func (g *Group) pprof() {
	if !gin.IsDebugging() {
		return
	}

	pprof.Register(g.r)
}

func (g *Group) status() {

	// 404 Handler.
	g.r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, vo.BaseResultNotFound)
	})

	// 405 Handler.
	g.r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, vo.BaseResultNotFound)
	})

}

func (g *Group) swagger() {
	if !gin.IsDebugging() {
		return
	}

	g.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
