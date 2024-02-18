package server

import (
	"log"
	"net/http"
)

func NewServer(
	logger *log.Logger,
	config Config,
	zajoncStore *ZajoncStore,
	zajoncService *ZajoncService,
) http.Handler {
	mux := http.NewServeMux()
	addRoutes(
		mux,
		logger,
		config,
		zajoncStore,
		zajoncService,
	)
	var handler http.Handler = mux
	// handler = authMiddleware(handler)
	// handler = someMiddleware(handler)
	// ...
	return handler
}

func addRoutes(
	mux *http.ServeMux,
	logger *log.Logger,
	config Config,
	clientStore *ZajoncStore,
	zajoncService *ZajoncService,
) {
}
