package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/hoanganh-ng/myhoot/app"
)

var appConfig *app.Config

func InitHelper(cfg *app.Config) {
	appConfig = cfg
}

func ClientError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	appConfig.ErrLog.Println(trace)
	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}
