package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/cstcen/gore/common"
	"github.com/ztrue/tracerr"
	"log/slog"
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

				slog.InfoContext(request.Context(), "recovery", "err", err.Error())
				slog.InfoContext(request.Context(), "recovery", "frames", frames)

				writer.WriteHeader(http.StatusOK)
				_ = json.NewEncoder(writer).Encode(common.BaseResultMetaServer.WithMsg(err.Error()))
			}
		}()

		handler.ServeHTTP(writer, request)
	})
}
