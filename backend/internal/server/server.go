// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import (
	"context"
	"fmt"
	"log"
	"trackposter/internal/repository"
	"trackposter/internal/soundcloud"

	"github.com/gin-gonic/gin"
)

type Server struct {
	verbose    bool
	router     *gin.Engine
	repository repository.Repository
	soundcloud soundcloud.SoundcloudConnector
}

// Is repository provider nil
func (s *Server) isRepositoryProviderValid() bool {
	return s.repository != nil
}

// Starts listening on port
func (s *Server) startServer(addr string) error {
	if s.verbose {
		log.Println("Starting listening on addr", addr)
	}

	if err := s.router.Run(addr); err != nil {
		return fmt.Errorf("run failed: %v", err)
	}

	return nil
}

// Run server with options
func (s *Server) Start(verbose bool, options ServerOptions) error {
	if s.verbose {
		log.Println("Using options:", options)
	}

	err := s.startServer(options.ValidAddress())
	if err != nil {
		return fmt.Errorf("failed to start handle")
	}

	return nil
}

// Turn on/off server verbose
func (s *Server) Verbose(v bool) {
	// ironically, isn't it?
	if s.verbose {
		log.Println("Setting verbose status to", v)
	}

	s.verbose = v
}

// Set repository provider for queue management
func (s *Server) SetRepository(repo repository.Repository) {
	if s.verbose {
		log.Printf("Setting new repository: %s\n", repo)
	}

	s.repository = repo
}

// Create new soundcloud client
func newSoundCloudClient() (*soundcloud.SoundcloudClient, error) {
	client, err := soundcloud.NewSCClient(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to create new client: %v", err)
	}

	return client, nil
}

// Creates new server instance
func NewServer(connector soundcloud.SoundcloudConnector) *Server {
	server := &Server{}
	server.router = server.setupRouter()
	server.soundcloud = connector

	return server
}
