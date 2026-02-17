// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://codeberg.com/TechAngle

package server

import (
	"fmt"
	"strings"
	"trackposter/internal/server/commons"
)

// Server options structure
type ServerOptions struct {
	Host string
	Port int
}

// Get full valid address.
//
// If host is not set, it uses default value (DefaultHost).
// If port is invalid, it replaces with default value (DefaultPort).
func (o *ServerOptions) ValidAddress() string {
	if strings.TrimSpace(o.Host) == "" {
		o.Host = commons.DefaultHost
	}

	if !isPortValid(o.Port) {
		o.Port = commons.DefaultPort
	}

	return fmt.Sprintf("%s:%d", o.Host, o.Port)
}
