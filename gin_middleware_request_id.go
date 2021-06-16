package gore

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strings"
)

const RequestIDContextKey = "X-Request-ID"

func GinRequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		id := uuid.New()

		ctx.Set(RequestIDContextKey, strings.ReplaceAll(id.String(), "-", ""))

		ctx.Next()
	}
}
