package gore

import "github.com/gin-gonic/gin"

func GinRest() func(c *gin.Context) {
	return func(c *gin.Context) {

		c.Next()

		c.Request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
}
