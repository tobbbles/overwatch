package list

import (
	"net/http"

	"service/store"

	"go.uber.org/zap"
)

var (
	Path    = "/api/abilities/"
	Methods = []string{"GET"}
)

// Endpoint ought to contain any dependencies used by the endpoint, such as
// database store clients, cache clients, or loggers.
type Endpoint struct {
	Logger *zap.Logger

	Provider store.Provider
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
