package api

import "net/http"

type Controller interface {
	RegisterRoute() error
}

type AuthController interface {
	Controller
	Login(w http.ResponseWriter, r *http.Request) //login with default password for admin
}

type GameController interface {
	Controller
	JoinGame(w http.ResponseWriter, r *http.Request)     //join game for player
	StartGame(w http.ResponseWriter, r *http.Request)    //admin start the game
	NextQuestion(w http.ResponseWriter, r *http.Request) //admin click go to the next question
}
