package web

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	router              *mux.Router
	gameController      *GameController
	utilitiesController *UtilitiesController
}

func NewHTTPServer() (*HTTPServer, error) {
	gameController := &GameController{}
	utilitiesController := &UtilitiesController{}
	router := mux.NewRouter()
	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static/"))),
	)
	router.HandleFunc("/", gameController.HomeController)
	router.HandleFunc("/page-not-found", utilitiesController.PageNotFound)
	return &HTTPServer{
		router:              router,
		gameController:      gameController,
		utilitiesController: utilitiesController,
	}, nil
}

func (server *HTTPServer) Serve() {
	addr := ":8092"
	log.Printf("http server running at %s\n", addr)
	err := http.ListenAndServe(
		addr,
		server.router,
	)
	log.Fatal(err)
}
