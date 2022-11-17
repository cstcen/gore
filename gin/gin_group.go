package gin

import (
	"git.tenvine.cn/backend/gore/common"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

const (
	RelativePathHealthCheck = "/healthcheck"
)

// Group struct
type Group struct {
	r *gin.Engine
}

func (g *Group) healthcheck() {
	g.r.GET(RelativePathHealthCheck, func(c *gin.Context) {
		c.JSON(http.StatusOK, common.BaseResultSuccess)
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
		c.JSON(http.StatusNotFound, common.BaseResultNotFound)
	})

	// 405 Handler.
	g.r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, common.BaseResultNotFound)
	})

}

func (g *Group) swagger() {
	if !gin.IsDebugging() {
		return
	}

	g.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		c.DocExpansion = "none"
		c.DefaultModelsExpandDepth = 0
		c.DeepLinking = true
	}))
}
