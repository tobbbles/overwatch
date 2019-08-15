package search

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"service/server/context/id/hero"
	"service/store"

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

	Provider store.Provider
}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := hero.FromContext(r.Context())
	if err != nil {
		e.Logger.Error("failed getting hero id from context", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Lookup the ID in the provider
	hero, err := e.Provider.Hero(id)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNoContent)
		return
	} else if err != nil {
		e.Logger.Error("failed getting hero from provider", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal the result back to the user
	if err := json.NewEncoder(w).Encode(hero); err != nil {
		e.Logger.Error("failed encoding hero object", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (e *Endpoint) Path() string {
	return Path
}

func (e *Endpoint) Methods() []string {
	return Methods
}
