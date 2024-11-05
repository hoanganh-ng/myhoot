package web

import (
	"html/template"
	"log"
	"net/http"
)

type UtilitiesController struct {
}

func (gc *UtilitiesController) PageNotFound(
	w http.ResponseWriter,
	req *http.Request,
) {
	tmpl, err := template.New("not-found.html").ParseFiles(
		"./src/html/pages/not-found.html",
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
