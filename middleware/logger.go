package middleware

import (
	"bytes"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/model"
	"git.tenvine.cn/backend/gore/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		blw := &ResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		c.Writer = blw

		// Request ID
		requestID, exists := c.Get(RequestIDContextKey)
		if !exists {
			requestID = util.GetRequestID()
		}

		requestIDLog := log.WithField(RequestIDContextKey, requestID)

		// Process request
		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}
		errorMsgs := c.Errors.ByType(gin.ErrorTypePrivate)

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = errorMsgs.String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		if param.ErrorMessage != "" {
			requestIDLog.Infof(
				"%-6s %-25s %-6v %-6v %-12s ---> %+v \n%v",
				param.Method,
				param.Path,
				param.StatusCode,
				param.Latency,
				param.ClientIP,
				blw.body,
				param.ErrorMessage,
			)
			var obj interface{} = model.BaseResultService
			last := errorMsgs.Last()
			if last != nil {
				obj = last
			}
			c.JSON(http.StatusOK, obj)
			return
		}

		requestIDLog.Infof(
			"%-6s %-25s %-6v %-6v %-12s ---> %+v",
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			blw.body,
		)
	}
}
