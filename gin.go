package gore

import (
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/middleware"
	"github.com/gin-gonic/gin"
)

// ginMode:
// - gin.DebugMode: 表示开发环境
// - gin.ReleaseMode: 表示正式环境
// - gin.TestMode: 暂时不用
func SetupGin(ginMode string) *gin.Engine {

	gin.SetMode(ginMode)

	// gin log to file
	gin.DefaultWriter = log.GetLogWriter()
	gin.DefaultErrorWriter = log.GetLogWriter()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debugf("%-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(middleware.Logger())

	r.Use(middleware.RequestID())

	r.Use(middleware.Rest())

	group := &Group{r: r}

	// check health
	group.ping()

	// swagger
	group.swagger()

	// pprof router
	group.pprof()

	group.status()

	return r
}
