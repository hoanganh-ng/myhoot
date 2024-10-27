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
	tmpl, err := template.ParseFiles("./src/html/index.html", "./src/html/components.html")
	if err != nil {
		log.Println(err)
		http.Error(w, "something went wrong.", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "DarkModeSwitch", nil)
	if err != nil {
		log.Println(err)
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
