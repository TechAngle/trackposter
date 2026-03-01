// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package server

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/gin-gonic/gin"
)

type HumaMiddlewareFunc func(ctx huma.Context, next func(huma.Context))

func corsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	ctx.Writer.Header().Add(
		"Access-Control-Allow-Headers",
		"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
	)

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(200)
		return
	}

	ctx.Next()
}

func (s *Server) repositoryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !s.isRepositoryProviderValid() {
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("Repository provider not set"))
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.Next()
	}
}
