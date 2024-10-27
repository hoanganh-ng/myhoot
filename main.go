package main

import (
	"log"

	"github.com/hoanganh-ng/myhoot/ports/web"
)

func main() {
	httpServer, err := web.NewHTTPServer()
	if err != nil {
		log.Fatal(err)
	}
	httpServer.Serve()
}
