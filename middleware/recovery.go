package middleware

import (
	"encoding/json"
	"fmt"
	"git.tenvine.cn/backend/gore/log"
	"git.tenvine.cn/backend/gore/vo"
	"github.com/ztrue/tracerr"
	"net/http"
)

func SetupRecovery(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if !ok {
					err = fmt.Errorf("%v", err)
				}

				e := tracerr.Wrap(err)
				frames := e.StackTrace()[4:5]

				log.WithContext(request.Context()).Info(tracerr.SprintSourceColor(tracerr.CustomError(err, frames)))

				writer.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(writer).Encode(vo.BaseResult{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
			}
		}()

		handler.ServeHTTP(writer, request)
	})
}
