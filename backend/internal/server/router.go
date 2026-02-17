// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://codeberg.com/TechAngle

package server

import (
	"github.com/gin-gonic/gin"
)

// registers /api/tracks group and its endpoints
func (s *Server) registerTrackApiGroup(apiGroup *gin.RouterGroup) {
	trackApiRoute := apiGroup.Group("/tracks")
	trackApiRoute.Use(s.repositoryMiddleware())

	// /tracks requests
	{
		trackApiRoute.POST("/addTrack", s.addTrack)
		trackApiRoute.GET("/track/{id}", s.trackById)
		trackApiRoute.GET("/queue", s.tracksList)
		trackApiRoute.DELETE("/removeTrack", s.removeTrack)
	}
}

// sets routes for /api
func (s *Server) setApiRoutes(router *gin.Engine) {
	// creating /api group
	apiGroup := router.Group("/api")
	// apiGroup.UseMiddleware(corsHeadersMiddleware)

	apiGroup.GET("/status", s.status)
	apiGroup.POST("/ping", s.ping)
	s.registerTrackApiGroup(apiGroup)
}

// initializes new router
func (s *Server) setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware)

	s.setApiRoutes(router)

	return router
}
