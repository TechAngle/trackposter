// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package main

import (
	"log"
	"trackposter/internal/repository"
	api "trackposter/internal/server"
	"trackposter/internal/server/commons"
)

func main() {
	server := api.NewServer()
	server.SetRepository(repository.NewMemoryQueue())

	if err := server.Start(true, api.ServerOptions{
		Port: commons.DefaultPort,
	}); err != nil {
		log.Panicln(err)
	}
}
