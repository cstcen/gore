package middleware

import (
	"github.com/cstcen/gore/auth/v2"
	"github.com/cstcen/gore/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authentication() func(c *gin.Context) {
	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		claims, err := auth.ParseJwt(authorization)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrUnauthenticated)
			return
		}

		auth.SetClaimsToGinContext(c, claims)

		c.Next()

	}
}
