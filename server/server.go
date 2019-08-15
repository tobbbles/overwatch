package server

import (
	"fmt"
	"net/http"
	"time"

	"service/server/middleware/json"
	"service/server/middleware/path"
	"service/store"

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

	Logger   *zap.Logger
	Provider store.Provider
}

// Server harbours all dependencies and service used in serving requests.
type Server struct {
	addr string

	r        *mux.Router
	logger   *zap.Logger
	provider store.Provider
}

// New creates a configured Server with the passed endpoints.
func New(config *Config) (*Server, error) {
	if config.Provider == nil {
		return nil, fmt.Errorf("%T.Provider must not be nil", config)
	}
	if config.Logger == nil {
		return nil, fmt.Errorf("%T.Logger must not be nil", config)
	}
	if len(config.Addr) == 0 {
		return nil, fmt.Errorf("%T.Addr must not be empty", config)
	}

	s := &Server{
		addr: config.Addr,
		r:    mux.NewRouter(),

		logger:   config.Logger,
		provider: config.Provider,
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
		&abilitieslist.Endpoint{Logger: s.logger, Provider: s.provider},
		&abilitiessearch.Endpoint{Logger: s.logger, Provider: s.provider},
		&herosabilities.Endpoint{Logger: s.logger, Provider: s.provider},
		&herossearch.Endpoint{Logger: s.logger, Provider: s.provider},
		&heroslist.Endpoint{Logger: s.logger, Provider: s.provider},
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
