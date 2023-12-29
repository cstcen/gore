package middleware

import (
	"git.tenvine.cn/backend/gore/common"
	"git.tenvine.cn/backend/gore/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Error() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		errorMsgs := c.Errors.ByType(gin.ErrorTypePrivate)
		if len(errorMsgs) == 0 {
			return
		}

		for i, e := range errorMsgs {
			log.WarningCf(c, "Error #%02d: %s", i+1, e.Err)
		}

		last := errorMsgs.Last()

		switch x := last.Err.(type) {
		case common.Error:
			c.JSON(http.StatusOK, x)
		default:
			c.JSON(http.StatusOK, common.NewBaseResult(http.StatusInternalServerError, x.Error()))
		}
		return

	}
}
