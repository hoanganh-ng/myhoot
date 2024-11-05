package web

import (
	"html/template"
	"log"
	"net/http"
)

type GameController struct {
}

func (gc *GameController) HomeController(
	w http.ResponseWriter,
	req *http.Request,
) {
	reqParams := req.URL.Query()
	game := reqParams.Get("game")
	if game != "" {
		http.Redirect(w, req, "/page-not-found", http.StatusTemporaryRedirect)
		return
	}

	tmpl, err := template.New("index.html").ParseFiles(
		"./src/html/pages/index.html",
		"./src/html/layouts/root.html",
		// "./src/html/blocks/head.html",
		// "./src/html/blocks/css.html",
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "something went wrong.", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
