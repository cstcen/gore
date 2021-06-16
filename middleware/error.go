package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		last := c.Errors.ByType(gin.ErrorTypePrivate).Last()
		if last != nil {

			c.AbortWithStatusJSON(http.StatusOK, last.Error())

		}

	}
}
