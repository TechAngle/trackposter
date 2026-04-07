// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

import "github.com/gin-gonic/gin"

// Handler is the main interface that should be implemented by any handler in this
// package.
type Handler interface {
	// RegisterRoutes registers all routes associated with this handler.
	RegisterRoutes(apiGroup *gin.RouterGroup)
}
