package middleware

import (
	"context"
	"git.tenvine.cn/backend/gore/util"
	"net/http"
)

func SetupRequestID(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		id := util.GenerateRequestID()
		request = request.WithContext(context.WithValue(request.Context(), util.RequestIDContextKey, id))

		handler.ServeHTTP(writer, request)

		request.Header.Add(util.RequestIDContextKey, id)

	})
}
