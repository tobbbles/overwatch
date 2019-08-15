package server

import (
	"net/http"
	"time"

	"service/server/middleware/json"
	"service/server/middleware/path"

	abilitieslist "service/server/endpoints/abilities/list"
	abilitiessearch "service/server/endpoints/abilities/search"
	herosabilities "service/server/endpoints/heros/abilities"
	heroslist "service/server/endpoints/heros/list"
	herossearch "service/server/endpoints/heros/search"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Config to apply to the created Server.
type Config struct {
	Addr string

	Logger *zap.Logger
}

// Server harbours all dependencies and service used in serving requests.
type Server struct {
	addr string

	r      *mux.Router
	logger *zap.Logger
}

// New creates a configured Server with the passed endpoints.
func New(config *Config) (*Server, error) {
	s := &Server{
		addr: config.Addr,
		r:    mux.NewRouter(),

		logger: config.Logger,
	}

	// Assign the middlewares.
	s.r.Use(
		path.Middleware(s.logger),
		json.Middleware(s.logger),
		mux.CORSMethodMiddleware(s.r),
	)

	// Set strict trailing slash
	s.r.StrictSlash(true)

	// Unfurl our endpoint collections and attach them to the router
	for _, endpoint := range s.endpoints() {
		s.r.Handle(endpoint.Path(), endpoint).Methods(endpoint.Methods()...)
	}

	return s, nil
}

// endpoints instantiates our collection of endpoints
func (s *Server) endpoints() []Endpoint {
	return []Endpoint{
		&abilitieslist.Endpoint{Logger: s.logger},
		&abilitiessearch.Endpoint{Logger: s.logger},
		&herosabilities.Endpoint{Logger: s.logger},
		&herossearch.Endpoint{Logger: s.logger},
		&heroslist.Endpoint{Logger: s.logger},
	}
}

// Start serving the server with responsible defaults.
func (s *Server) Start() error {
	serve := &http.Server{
		Addr:    s.addr,
		Handler: s.r,

		// Sensible defaults
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       20 * time.Second,
		MaxHeaderBytes:    32 * 1024,
	}

	return serve.ListenAndServe()
}
