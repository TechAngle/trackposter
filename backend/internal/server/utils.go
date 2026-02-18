// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import (
	"fmt"
	"net/http"
	"regexp"
	"trackposter/internal/server/commons"
	"trackposter/internal/server/models"

	"github.com/gin-gonic/gin"
)

var (
	soundCloudRegex = regexp.MustCompile(`^https://(soundcloud\.com/[A-Za-z0-9\-_]+/[A-Za-z0-9\-_]+|on\.soundcloud\.com/[A-Za-z0-9]+)(\?.*)?$`)
)

// checks if port valid
func isPortValid(port int) bool {
	return port >= commons.MinPort && port <= commons.MaxPort
}

// check if string is SoundCloud link
func isSoundcloudLink(link string) bool {
	return soundCloudRegex.Match([]byte(link))
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
