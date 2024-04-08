package middleware

import (
	"fmt"
	"github.com/cstcen/gore/common"
	"github.com/gin-gonic/gin"
	"log/slog"
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
			slog.WarnContext(c, fmt.Sprintf("Error #%02d: %s", i+1, e.Err))
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
