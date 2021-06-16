package gore

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// Group struct
type Group struct {
	r *gin.Engine
}

func (g *Group) ping() {
	g.r.GET("/healthcheck", func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
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
		c.String(http.StatusNotFound, "The incorrect API router.")
	})

	// 405 Handler.
	g.r.NoMethod(func(c *gin.Context) {
		c.String(http.StatusMethodNotAllowed, "The incorrect API router.")
	})

}

func (g *Group) swagger() {
	if !gin.IsDebugging() {
		return
	}

	g.r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
