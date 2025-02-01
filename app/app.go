package app

import (
	"errors"
	"log"
	"os"

	"github.com/gorilla/sessions"
)

type App struct {
	config       *Config
	sessionStore *sessions.CookieStore
}

func New() (*App, error) {
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		return nil, errors.New("unset session key")
	}
	infoLog := log.New(os.Stdout, "[Info]\t", log.Ldate|log.Ltime)
	errLog := log.New(os.Stdout, "[Err]\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig := &Config{
		ErrLog:  errLog,
		InfoLog: infoLog,
	}

	_app := &App{
		config:       appConfig,
		sessionStore: sessions.NewCookieStore([]byte(sessionKey)),
	}
	return _app, nil
}
