package middleware

import (
	"git.tenvine.cn/backend/gore/httputil"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/util"
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

		requestIDLog := log.WithField(util.RequestIDContextKey, util.GetRequestID(c))

		for i, e := range errorMsgs {
			requestIDLog.Warnf("Error #%02d: %s", i+1, e.Err)
		}

		last := errorMsgs.Last()

		switch x := last.Err.(type) {
		case httputil.BaseResult:
			c.JSON(http.StatusOK, x)
		default:
			c.JSON(http.StatusOK, httputil.BaseResult{Code: http.StatusInternalServerError, Message: x.Error()})
		}
		return

	}
}
