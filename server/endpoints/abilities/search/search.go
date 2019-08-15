package search

import (
	"go.uber.org/zap"
	"net/http"
	"service/server/context/id/ability"
)

var (
	Path    = "/api/abilities/{ability_id}/"
	Methods = []string{"GET"}
)

// Endpoint ought to contain any dependencies used by the endpoint, such as
// database store clients, cache clients, or loggers.
type Endpoint struct {
	Logger *zap.Logger
}

// ServeHTTP for the ability searcher endpoint.
func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := ability.FromContext(r.Context())
	if err != nil {
		panic(err)
	}

	// TODO: Make a database/cache request for the given Ability ID

	w.WriteHeader(http.StatusNotImplemented)
}

func (e *Endpoint) Path() string {
	return Path
}

func (e *Endpoint) Methods() []string {
	return Methods
}
