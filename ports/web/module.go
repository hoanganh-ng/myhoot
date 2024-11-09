package web

import "net/http"

type Handler interface {
	RegisterRoutes(*http.ServeMux)
}

type Module interface {
	Controllers() []*Handler
}
