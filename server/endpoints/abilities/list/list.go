package list

import (
	"database/sql"
	"encoding/json"
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
	abilities, err := e.Provider.Abilities()
	if err == sql.ErrNoRows {
		w.WriteHeader(http.StatusNoContent)
		return
	} else if err != nil {
		e.Logger.Error("failed getting abilities from provider", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Marshal the result back to the user
	if err := json.NewEncoder(w).Encode(abilities); err != nil {
		e.Logger.Error("failed encoding abilities collection", zap.Error(err))
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
