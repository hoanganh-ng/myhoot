package app

import (
	"errors"
	"log"
	"os"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	config         *Config
	adminHashedPwd []byte
	sessionStore   *sessions.CookieStore
}

func New() (*App, error) {
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		return nil, errors.New("unset admin password")
	}
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
	hashed, err := bcrypt.GenerateFromPassword([]byte(adminPassword), 14)
	if err != nil {
		return nil, err
	}
	_app := &App{
		config:         appConfig,
		adminHashedPwd: hashed,
		sessionStore:   sessions.NewCookieStore([]byte(sessionKey)),
	}
	return _app, nil
}
