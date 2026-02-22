// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import (
	"fmt"
	"net/http"
	"trackposter/internal/server/commons"
	"trackposter/internal/server/models"

	"github.com/gin-gonic/gin"
)

// checks if port valid
func isPortValid(port int) bool {
	return port >= commons.MinPort && port <= commons.MaxPort
}

// Binds json to strutcure. If something gone wrong - sets 500 to context and returns an error
func parseRequest[T any](ctx *gin.Context) (*T, error) {
	var obj T
	if err := ctx.ShouldBindJSON(&obj); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.StatusResponse{
			StatusMessage: fmt.Sprintf("Cannot bind JSON: %v", err),
		})
		return nil, err
	}

	return &obj, nil
}
