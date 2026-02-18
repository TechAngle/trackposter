// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import (
	"fmt"
	"net/http"
	"strings"
	"trackposter/internal/server/models"

	"github.com/gin-gonic/gin"
)

// Put new track to line.
func (s *Server) addTrack(ctx *gin.Context) {
	// TODO: Parse length of track by URL if duration was not set

	// processing track request
	var trackRequest models.Track
	if err := ctx.ShouldBindJSON(&trackRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, models.StatusResponse{
			StatusMessage: "Unprocessable Entity",
		})
		return
	}

	// check if link valid
	if !isSoundcloudLink(trackRequest.URL) {
		ctx.JSON(http.StatusBadRequest, models.StatusResponse{
			StatusMessage: "Invalid SoundCloud link",
		})
		return
	}

	// adding track
	trackId, err := s.repository.AddTrack(&trackRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.StatusResponse{
			StatusMessage: fmt.Sprintf("failed to add track to queue: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusAccepted, models.AddTrackResponse{
		TrackID:       trackId,
		StatusMessage: "Track successfully added",
	})
}

// Get track by its ID
func (s *Server) trackById(ctx *gin.Context) {
	request, err := parseRequest[models.GetTrackRequest](ctx)
	if err != nil {
		// response was set in method
		return
	}

	// checking that trackId is valid
	if strings.TrimSpace(request.TrackId) == "" {
		ctx.JSON(http.StatusBadRequest, models.StatusResponse{
			StatusMessage: "Track id required",
		})
	}

	// finding it in repository
	track := s.repository.TrackByID(request.TrackId)
	if track == nil {
		ctx.JSON(http.StatusNotFound, models.StatusResponse{
			StatusMessage: "Track not found",
		})

		return
	}

	ctx.JSON(http.StatusOK, models.GetTrackResponse{
		Track: *track,
	})
}

// TODO: rewrite to gin
// Remove track by its ID
func (s *Server) removeTrack(ctx *gin.Context) {
	request, err := parseRequest[models.RemoveTrackRequest](ctx)
	if err != nil {
		// response was set in method
		return
	}

	err = s.repository.RemoveTrack(request.TrackID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.StatusResponse{
			StatusMessage: fmt.Sprintf("Failed to remove track: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, &models.StatusResponse{
		StatusMessage: "Track removed",
	})
}

// Get tracks list
func (s *Server) tracksList(ctx *gin.Context) {
	tracksList := s.repository.Queue()
	ctx.JSON(http.StatusOK, tracksList)
}
