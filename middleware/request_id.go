package middleware

import (
	"git.tenvine.cn/backend/gore/util"
	"github.com/gin-gonic/gin"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := util.GenerateRequestID()
		c.Set(util.RequestIDContextKey, id)

		c.Next()

		c.Request.Header.Set(util.RequestIDContextKey, id)
	}
}
