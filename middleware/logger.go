package middleware

import (
	"bytes"
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

func (w *GinResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *GinResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func Logger(skipLogResp func(path string) bool) gin.HandlerFunc {
	return func(c *gin.Context) {
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
		var reqBuf bytes.Buffer
		var reqBody []byte
		if c.Request.Body != nil {
			tee := io.TeeReader(c.Request.Body, &reqBuf)
			reqBody, _ = io.ReadAll(tee)
			c.Request.Body = io.NopCloser(&reqBuf)
		}

		c.Writer = respWriter

		log.DebugCf(c, "Req URL:    %s", c.Request.URL.String())
		log.DebugCf(c, "Req Header: %+v", header)
		log.DebugCf(c, "Req Body:   %q", string(reqBody))

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
			log.DebugCf(c, "Response body : %s", respWriter.body.String())
		}
		log.DebugCf(c, logStr)
	}
}

type ResponseWriter struct {
	http.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (rw *ResponseWriter) Write(b []byte) (int, error) {
	rw.body.Write(b)
	return rw.ResponseWriter.Write(b)
}

func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func SetupTrace(handler http.Handler, skipLogResp func(path string) bool) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		c := request.Context()
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
		var reqBuf bytes.Buffer
		var reqBody []byte
		if request.Body != nil {
			tee := io.TeeReader(request.Body, &reqBuf)
			reqBody, _ = io.ReadAll(tee)
			request.Body = io.NopCloser(&reqBuf)
		}

		writer = respWriter

		log.DebugCf(c, "Req URL:    %s", request.URL.String())
		log.DebugCf(c, "Req Header: %+v", header)
		log.DebugCf(c, "Req Body:   %q", string(reqBody))

		// Process request
		handler.ServeHTTP(writer, request)

		// Stop timer
		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		clientIP := request.RemoteAddr
		method := request.Method
		statusCode := respWriter.statusCode

		logStr := fmt.Sprintf(
			"method=%v uri=%v status=%v latency=%v ip=%v",
			method,
			path,
			statusCode,
			latency,
			clientIP,
		)
		if !skipLogResp(path) {
			log.DebugCf(c, "Response body : %s", respWriter.body.Bytes())
		}
		log.DebugCf(c, logStr)
	})
}
