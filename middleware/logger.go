package middleware

import (
	"bytes"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/util"
	"github.com/gin-gonic/gin"
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
		requestID := util.GetRequestID(c)

		requestIDLog := log.WithField(util.RequestIDContextKey, requestID)

		// Process request
		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		requestIDLog.WithField(
			"resp",
			blw.body,
		).Infof(
			"%-6s %-25s %-6v %-6v %-12s",
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
		)
	}
}
