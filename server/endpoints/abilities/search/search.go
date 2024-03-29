package search

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"service/server/context/id/ability"
	"service/store"

	"go.uber.org/zap"
)

var (
	Path    = "/api/abilities/{ability_id}/"
	Methods = []string{"GET"}
)

// Endpoint ought to contain any dependencies used by the endpoint, such as
// database store clients, cache clients, or loggers.
type Endpoint struct {
	Logger *zap.Logger

	Provider store.Provider
}

// ServeHTTP for the ability searcher endpoint.
func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := ability.FromContext(r.Context())
	if err != nil {
		e.Logger.Error("failed getting hero id from context", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Lookup the ID in the provider
	ability, err := e.Provider.Ability(id)
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNoContent)
		return
	} else if err != nil {
		e.Logger.Error("failed getting ability from provider", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal the result back to the user
	if err := json.NewEncoder(w).Encode(ability); err != nil {
		e.Logger.Error("failed encoding ability object", zap.Error(err))
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
