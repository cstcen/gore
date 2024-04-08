package middleware

import (
	"context"
	"github.com/cstcen/gore/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := util.GenerateRequestID()
		c.Set(util.RequestIDContextKey, id)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), util.RequestIDContextKey, id))

		c.Next()

		c.Writer.Header().Set(util.RequestIDContextKey, id)
	}
}

func SetupRequestID(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		id := util.GenerateRequestID()
		request = request.WithContext(context.WithValue(request.Context(), util.RequestIDContextKey, id))

		handler.ServeHTTP(writer, request)

		request.Response.Header.Add(util.RequestIDContextKey, id)

	})
}
