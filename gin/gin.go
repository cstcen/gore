package gin

import (
	"context"
	"fmt"
	"github.com/cstcen/gore/common"
	"github.com/cstcen/gore/consul"
	"github.com/cstcen/gore/gonfig"
	goreMiddleware "github.com/cstcen/gore/middleware"
	"github.com/cstcen/gore/util"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
)

var engine *gin.Engine
var shutdownHooks = make([]func() error, 0)

// GetInstance Use Instance()
// Deprecated
func GetInstance() *gin.Engine {
	return Instance()
}

func Instance() *gin.Engine {
	return engine
}

func AddShutdownHook(fn func() error) {
	shutdownHooks = append(shutdownHooks, fn)
}

func Startup() error {

	port := gonfig.Instance().GetString("port")
	addr := fmt.Sprintf(":%s", port)

	srv := &http.Server{
		Addr:     addr,
		Handler:  engine,
		ErrorLog: log.Default(),
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("listen")
		}
	}()

	slog.Info("Server run", "port", port)
	slog.Info("        #####    #######   ######    #######        ")
	slog.Info("       #     #   #     #   #     #   #              ")
	slog.Info("       #         #     #   #     #   #              ")
	slog.Info("       #  ####   #     #   ######    #####          ")
	slog.Info("       #     #   #     #   #   #     #              ")
	slog.Info("       #     #   #     #   #    #    #              ")
	slog.Info("        #####    #######   #     #   #######        ")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	slog.Info("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	slog.Info("Server exiting")

	if consul.Enable() {
		if err := consul.Deregister(); err != nil {
			return err
		}
	}

	for _, fn := range shutdownHooks {
		_ = fn()
	}

	return nil
}

func Setup() error {

	dev := "xk5" != gonfig.Instance().GetString("env")
	if !dev {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin log to file
	writer := log.Default().Writer()
	gin.DefaultWriter = writer
	gin.DefaultErrorWriter = writer

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		slog.Info(fmt.Sprintf("%-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers))
	}

	engine = gin.New()

	engine.Use(goreMiddleware.RequestID())

	engine.Use(goreMiddleware.Logger(func(path string) bool {
		return strings.Contains(path, "/pprof") || strings.Contains(path, "/swagger")
	}))

	engine.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		switch e := err.(type) {
		default:
			c.AbortWithStatusJSON(http.StatusOK, common.ErrService.WithMsg(fmt.Sprintf("%s", err)))
		case common.Error:
			c.AbortWithStatusJSON(e.GetHttpStatus(), e)
		}
	}))

	// check health
	routeHealthcheck(engine)

	// swagger
	// routeSwagger(engine)

	// pprof router
	routePprof(engine)

	routeStatus(engine)

	return nil
}

const (
	RelativePathHealthCheck = "/healthcheck"
)

func routeHealthcheck(e *gin.Engine) {
	e.GET(RelativePathHealthCheck, func(c *gin.Context) {
		c.JSON(http.StatusOK, common.BaseResultSuccess)
	})
	e.HEAD(RelativePathHealthCheck, func(c *gin.Context) {
		c.JSON(http.StatusOK, common.BaseResultSuccess)
	})
}

func routePprof(e *gin.Engine) {
	if !gin.IsDebugging() {
		return
	}

	pprof.Register(e)
}

func routeStatus(e *gin.Engine) {

	// 404 Handler.
	e.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, common.BaseResultNotFound)
	})

	// 405 Handler.
	e.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, common.BaseResultNotFound)
	})

}

func routeSwagger(e *gin.Engine) {
	if !gin.IsDebugging() {
		return
	}

	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, func(c *ginSwagger.Config) {
		c.DocExpansion = "none"
		c.DefaultModelsExpandDepth = 0
		c.DeepLinking = true
	}))
}

func LogFormatter(param gin.LogFormatterParams) string {
	requestId, _ := param.Keys[util.RequestIDContextKey].(string)
	if len(requestId) == 0 {
		requestId, _ = param.Request.Context().Value(util.RequestIDContextKey).(string)
	}
	return fmt.Sprintf("%v [GORE] [%s] | %3d | %13v | %15s |%-7s %#v\n%s",
		param.TimeStamp.Format("2006/01/02 15:04:05"),
		requestId,
		param.StatusCode,
		param.Latency,
		param.ClientIP,
		param.Method,
		param.Path,
		param.ErrorMessage,
	)
}
