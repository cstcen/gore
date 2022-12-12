package gin

import (
	"context"
	"fmt"
	"git.tenvine.cn/backend/gore/common"
	"git.tenvine.cn/backend/gore/consul"
	"git.tenvine.cn/backend/gore/gonfig"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var engine *gin.Engine

func GetInstance() *gin.Engine {
	return engine
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

	log.Printf("Server run port: %s", port)
	log.Printf("        #####    #######   ######    #######        ")
	log.Printf("       #     #   #     #   #     #   #              ")
	log.Printf("       #         #     #   #     #   #              ")
	log.Printf("       #  ####   #     #   ######    #####          ")
	log.Printf("       #     #   #     #   #   #     #              ")
	log.Printf("       #     #   #     #   #    #    #              ")
	log.Printf("        #####    #######   #     #   #######        ")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Print("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	log.Print("Server exiting")

	if consul.Enable() {
		if err := consul.Deregister(); err != nil {
			return err
		}
	}

	return nil
}

func Setup() error {

	dev := "xk5" == gonfig.Instance().GetString("env")
	if !dev {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin log to file
	gin.DefaultWriter = log.Default().Writer()
	gin.DefaultErrorWriter = log.Default().Writer()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	engine = gin.New()

	engine.Use(gin.CustomRecovery(func(c *gin.Context, err any) {
		c.AbortWithStatusJSON(http.StatusOK, common.BaseResultService.WithMsg(fmt.Sprintf("%s", err)))
	}))

	if dev {
		engine.Use(cors.Default())
	}

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
