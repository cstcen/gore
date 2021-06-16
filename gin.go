package gore

import (
	"git.tenvine.cn/backend/gore/log"
	"github.com/gin-gonic/gin"
)

// ginMode
// - gin.DebugMode
// - gin.ReleaseMode
// - gin.TestMode
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

	r.Use(GinLogger())

	r.Use(GinRequestID())

	r.Use(GinRest())

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
