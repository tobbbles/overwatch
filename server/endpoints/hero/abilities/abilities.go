package abilities

import "net/http"

var (
	Path    = "/api/hero/{hero_id}/abilities/"
	Methods = []string{"GET"}
)

// Endpoint ought to contain any dependencies used by the endpoint, such as
// database store clients, cache clients, or loggers.
type Endpoint struct{}

func (e *Endpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := []byte(`{"status": "ok"}`)

	w.Write(response)
}

func (e *Endpoint) Path() string {
	return Path
}

func (e *Endpoint) Methods() []string {
	return Methods
}
