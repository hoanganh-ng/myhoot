package app

import "log"

type Config struct {
	ErrLog  *log.Logger
	InfoLog *log.Logger
}
