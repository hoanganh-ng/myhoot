package api

import "github.com/hoanganh-ng/myhoot/app"

var appConfig *app.Config

type API struct {
}

func NewAPI(cfg *app.Config) {
	appConfig = cfg
}
