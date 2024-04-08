package middleware

import "github.com/gin-gonic/gin"

func Rest() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.Next()

		c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}
}
