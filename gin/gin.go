package gin

import (
	"context"
	"fmt"
	"git.tenvine.cn/backend/gore/consul"
	"git.tenvine.cn/backend/gore/gonfig"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/middleware"
	"github.com/gin-gonic/gin"
	syslog "log"
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
		ErrorLog: syslog.New(log.GetLogWriter(), "", 0),
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.WithError(err).Fatal("listen")
		}
	}()

	log.Infof("Server run port: %s", port)
	log.Infof("        #####    #######   ######    #######        ")
	log.Infof("       #     #   #     #   #     #   #              ")
	log.Infof("       #         #     #   #     #   #              ")
	log.Infof("       #  ####   #     #   ######    #####          ")
	log.Infof("       #     #   #     #   #   #     #              ")
	log.Infof("       #     #   #     #   #    #    #              ")
	log.Infof("        #####    #######   #     #   #######        ")

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		return err
	}
	log.Info("Server exiting")

	if consul.Enable() {
		if err := consul.Deregister(); err != nil {
			return err
		}
	}

	return nil
}

func Setup() error {

	mode := gin.DebugMode
	if "xk5" == gonfig.Instance().GetString("env") {
		mode = gin.ReleaseMode
	}

	gin.SetMode(mode)

	// gin log to file
	gin.DefaultWriter = log.GetLogWriter()
	gin.DefaultErrorWriter = log.GetLogWriter()

	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debugf("%-6s %-25s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
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
