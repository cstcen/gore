package gin

import (
	"git.tenvine.cn/backend/gore/db"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/model"
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
		Cache:         gonfig.GetInstance().Gore.Cache,
		Elasticsearch: gonfig.GetInstance().Gore.Elasticsearch,
		Mongo:         gonfig.GetInstance().Gore.Mongo,
		Mysql:         gonfig.GetInstance().Gore.Mysql,
		Redis:         gonfig.GetInstance().Gore.Redis,
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
		c.JSON(http.StatusNotFound, model.BaseResultNotFound)
	})

	// 405 Handler.
	g.r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, model.BaseResultNotFound)
	})

}

func (g *Group) swagger() {
	if !gin.IsDebugging() {
		return
	}

	g.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
