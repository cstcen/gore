package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.tenvine.cn/backend/gore/log"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

type GinResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w GinResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w GinResponseWriter) WriteString(s string) (int, error) {
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
		respWriter := &GinResponseWriter{
			ResponseWriter: c.Writer,
			body:           bytes.NewBufferString(""),
		}
		var buf bytes.Buffer
		var body []byte
		if c.Request.Body != nil {
			tee := io.TeeReader(c.Request.Body, &buf)
			body, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&buf)
		}

		c.Writer = respWriter

		contextLog.Tracef("Request url   : %s", c.Request.URL.String())
		contextLog.Tracef("Request header: %+v", header)
		compactBody := new(bytes.Buffer)
		_ = json.Compact(compactBody, body)
		contextLog.Tracef("Request body  : %s", bytes.TrimSpace(body))

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

		logStr := fmt.Sprintf(
			"method=%s uri=%s status=%v latency=%v ip=%s",
			param.Method,
			param.Path,
			param.StatusCode,
			param.Latency,
			param.ClientIP,
		)
		if !skipLogResp(path) {
			contextLog.Tracef("Response body : %s", respWriter.body.String())
		}
		contextLog.Info(logStr)
	}
}

type ResponseWriter struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (rw ResponseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func SetupTrace(handler http.Handler, skipLogResp func(path string) bool) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		contextLog := log.WithContext(ctx)
		// Start timer
		start := time.Now()
		header := request.Header
		path := request.URL.Path
		raw := request.URL.RawQuery
		if len(raw) > 0 {
			path = path + "?" + raw
		}
		respWriter := &ResponseWriter{
			ResponseWriter: writer,
			body:           new(bytes.Buffer),
		}
		var buf bytes.Buffer
		var body []byte
		if request.Body != nil {
			tee := io.TeeReader(request.Body, &buf)
			body, _ = io.ReadAll(tee)
			request.Body = io.NopCloser(&buf)
		}

		writer = respWriter

		contextLog.Tracef("Request url   : [%v] %s", request.Method, path)
		contextLog.Tracef("Request header: %+v", header)
		contextLog.Tracef("Request body  : %s", bytes.TrimSpace(body))

		// Process request
		handler.ServeHTTP(writer, request)

		// Stop timer
		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		clientIP := request.RemoteAddr
		method := request.Method
		statusCode := respWriter.statusCode
		// bodySize := writer.Header().Get("Content-Length")

		logStr := fmt.Sprintf(
			"method=%v uri=%v status=%v latency=%v ip=%v",
			method,
			path,
			statusCode,
			latency,
			clientIP,
		)
		if !skipLogResp(path) {
			contextLog.Tracef("Response body : %s", respWriter.body.Bytes())
		}
		contextLog.Info(logStr)
	})
}
