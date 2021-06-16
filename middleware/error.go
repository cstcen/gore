package middleware

import (
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/model"
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

		switch last.Err.(type) {
		case model.BaseResult:
			c.JSON(http.StatusOK, last.Err)
		default:
			c.JSON(http.StatusOK, model.BaseResult{Code: http.StatusBadRequest, Message: last.Err.Error()})
		}
		return

	}
}
