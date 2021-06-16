package middleware

import (
	"git.tenvine.cn/backend/gore/util"
	"github.com/gin-gonic/gin"
)

const RequestIDContextKey = "X-Request-ID"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := util.GetRequestID()
		c.Set(RequestIDContextKey, id)

		c.Next()

		c.Request.Response.Header.Set(RequestIDContextKey, id)
	}
}
