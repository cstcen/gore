package middleware

import (
	"bytes"
	"git.tenvine.cn/backend/gore/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

func Logger(skipLogResp func(path string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		contextLog := log.WithContext(c)
		// Start timer
		start := time.Now()
		header := c.Request.Header
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if len(raw) > 0 {
			path = path + "?" + raw
		}
		respWriter := &ResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		body, _ := io.ReadAll(tee)
		c.Request.Body = io.NopCloser(&buf)

		c.Writer = respWriter

		contextLog.Tracef("Request url:    %s", path)
		contextLog.Tracef("Request header: %+v", header)
		contextLog.Tracef("Request body:   %+v", string(body))

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

		param.Path = path

		fields := logrus.Fields{
			"method":  param.Method,
			"uri":     param.Path,
			"status":  param.StatusCode,
			"latency": param.Latency,
			"ip":      param.ClientIP,
			"body":    respWriter.body.String(),
		}
		if skipLogResp(path) {
			fields["body"] = ""
		}
		contextLog.WithFields(fields).Info()
	}
}
