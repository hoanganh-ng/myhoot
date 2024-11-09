package httpserver

import (
	"net/http"

	"github.com/hoanganh-ng/myhoot/ports/api"
)

type QuestionController struct {
	router          *http.ServeMux
	questionService *api.QuestionService
}

func (c QuestionController) RegisterRoutes() {
	c.router.HandleFunc("POST", c.CreateQuestion)
}

func (c QuestionController) CreateQuestion(
	w http.ResponseWriter,
	r *http.Request,
) {

}
