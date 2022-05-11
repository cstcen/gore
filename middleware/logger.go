package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.tenvine.cn/backend/gore/log"
	"io"
	"net/http"
	"time"
)

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
		compactBody := new(bytes.Buffer)
		_ = json.Compact(compactBody, body)
		contextLog.Tracef("Request body  : %s", compactBody)

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
			logStr = logStr + fmt.Sprintf(" body=%s", respWriter.body.String())
		}
		contextLog.Info(logStr)
	})
}
