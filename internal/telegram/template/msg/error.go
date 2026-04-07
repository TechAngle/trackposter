// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package msg

import (
	"fmt"
	"html"
	"trackposter/internal/domain"
)

// errorTemplate returns default template for error messages.
//
// Template look: ERROR: {message}
func errorTemplate(message string) domain.MessageTemplate {
	return domain.MessageTemplate(
		fmt.Sprintf(
			"<b>ERROR</b>: <i>%s</i>",
			html.EscapeString(message),
		),
	)
}

// ErrorTrackNotFound template when track was not found on platform.
func ErrorTrackNotFound() domain.MessageTemplate {
	return errorTemplate("Track not found.")
}

// ErrorInvalidURL template when user has provided an invalid URL.
func ErrorInvalidURL() domain.MessageTemplate {
	return errorTemplate("Invalid URL.")
}
