package gin

import (
	"context"
	"fmt"
	"git.tenvine.cn/backend/gore/consul"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/middleware"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
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

	if "xk5" == gonfig.Instance().GetString("env") {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin log to file
	gin.DefaultWriter = log.Default().Writer()
	gin.DefaultErrorWriter = log.Default().Writer()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Printf("%-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}

	engine = gin.New()

	engine.Use(gin.Recovery())

	engine.Use(middleware.RequestID())

	engine.Use(middleware.Logger(func(path string) bool {
		return strings.Contains(path, "/swagger/") || strings.Contains(path, "/pprof/")
	}))

	engine.Use(middleware.Rest())

	engine.Use(middleware.Error())

	group := &Group{r: engine}

	// check health
	group.healthcheck()

	// swagger
	group.swagger()

	// pprof router
	group.pprof()

	group.status()

	return nil
}
