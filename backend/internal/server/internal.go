// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import (
	"net/http"
	"time"
	"trackposter/internal/server/models"

	"github.com/gin-gonic/gin"
)

// just returns 200 and status ok if we are alive
func (s *Server) status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, models.StatusResponse{
		StatusMessage: "ok",
	})
}

// calculate difference between server and client
func (s *Server) ping(ctx *gin.Context) {
	var pingRequest models.PingRequest
	if err := ctx.ShouldBindJSON(&pingRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, models.StatusResponse{
			StatusMessage: "Bad request",
		})
		return
	}

	// getting current timestamp
	t := time.Now().Unix()
	delta := t - int64(pingRequest.Timestamp)

	ctx.JSON(http.StatusOK, models.PingResponse{
		Delta: int(delta),
	})
}
