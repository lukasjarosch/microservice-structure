package http

import (
	"context"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/lukasjarosch/microservice-structure/pkg/config"
)

// Endpoint wraps a HTTP endpoint
type Endpoint struct {
	Pattern string
	Handler http.Handler
}

// Server wraps a http.Server
type Server struct {
	Options    *Options
	HTTPServer *http.Server
	Endpoints  []Endpoint
}

// NewServer returns a pre-configured http.Server
func NewServer(opts ...Option) *Server {
	options := &Options{
		Network: config.Network{
			Host: "0.0.0.0",
			Port: 8080,
		},
	}

	for _, option := range opts {
		option(options)
	}

	return &Server{
		Options: options,
	}
}

// AddEndpoint adds an enpoint to the server
func (s *Server) AddEndpoint(endpoint Endpoint) {
	s.Endpoints = append(s.Endpoints, endpoint)
}

// ServeHTTP will serve the HTTP server in a goroutine
func (s *Server) ServeHTTP() {
	s.HTTPServer = &http.Server{Addr: s.Options.Network.Address()}

	// register all endpoint handlers
	for _, endpoint := range s.Endpoints {
		http.Handle(endpoint.Pattern, endpoint.Handler)
	}
	logrus.Infof("serving HTTP server at %s", s.Options.Network.Address())

	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Errorf("http server: ListenAndServe() error: %s", err)
			return
		}
	}()
}

// Shutdown handles the graceful shutdown of the HTTP server
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := s.HTTPServer.Shutdown(ctx); err != nil {
		logrus.Infof("timeout during shutdown of prometheus HTTP server: %v", err)
	}
}
