package middleware

import (
	"context"
	"github.com/cstcen/gore/db/mysql"
	"net/http"
	"time"
)

func SetupDB(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelFunc()
		ctx := context.WithValue(request.Context(), "DB", mysql.GormDB().WithContext(timeoutCtx))
		handler.ServeHTTP(writer, request.WithContext(ctx))
	})
}
