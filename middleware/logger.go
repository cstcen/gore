package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"strings"
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

		slog.DebugContext(c, "Request", "URL", c.Request.URL.String())
		slog.DebugContext(c, "Request", "Header", header)

		if strings.Contains(header.Get("Content-Type"), "application/json") {
			slog.DebugContext(c, "Request", "Body", string(reqBody))
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

		param.Path = path

		if !skipLogResp(path) && strings.Contains(c.Writer.Header().Get("Content-Type"), "application/json") {
			slog.DebugContext(c, "Response", "body", respWriter.body.String())
		}
		slog.InfoContext(c, "GORE", "method", param.Method, "uri", param.Path, "status", param.StatusCode, "latency", param.Latency, "ip", param.ClientIP)
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

		slog.DebugContext(c, "Request", "URL", request.URL.String())
		slog.DebugContext(c, "Request", "Header", header)
		slog.DebugContext(c, "Request", "Body", string(reqBody))

		// Process request
		handler.ServeHTTP(writer, request)

		// Stop timer
		timeStamp := time.Now()
		latency := timeStamp.Sub(start)
		clientIP := request.RemoteAddr
		method := request.Method
		statusCode := respWriter.statusCode

		if !skipLogResp(path) {
			slog.DebugContext(c, "Response", "body", respWriter.body.Bytes())
		}
		slog.InfoContext(c, "GORE", "method", method, "uri", path, "status", statusCode, "latency", latency, "ip", clientIP)
	})
}
