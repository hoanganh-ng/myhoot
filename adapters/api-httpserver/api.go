package apihttpserver

import "github.com/hoanganh-ng/myhoot/app"

var appConfig *app.Config

type API struct {
	// gameController     *GameController
	questionController *QuestionController
}

func New(config *app.Config) *API {
	appConfig = config
	return &API{}
}
