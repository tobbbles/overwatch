package search

import (
	"net/http"

	"go.uber.org/zap"
)

var (
	Path    = "/api/heros/{hero_id}/"
	Methods = []string{"GET"}
)

// Endpoint ought to contain any dependencies used by the endpoint, such as
// database store clients, cache clients, or loggers.
type Endpoint struct {
	Logger *zap.Logger
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotImplemented)
}

func (e *Endpoint) Path() string {
	return Path
}

func (e *Endpoint) Methods() []string {
	return Methods
}