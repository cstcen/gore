package middleware

import (
	"bytes"
	"git.tenvine.cn/backend/gore/log"
	"github.com/gin-gonic/gin"
	"io"
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

		contextLog := log.WithContext(c)

		contextLog.Infof("Request url: %s", c.FullPath())
		contextLog.Infof("Header: %+v", c.Request.Header)
		readCloser, err := c.Request.GetBody()
		if err == nil {
			body, err := io.ReadAll(readCloser)
			if err == nil {
				contextLog.Infof("Body: %+v", string(body))
			}
		}

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

		contextLog.Infof(
			"%-6s %-25s %-6v %-6v %-12s ---> %s",
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			blw.body.String(),
		)
	}
}
