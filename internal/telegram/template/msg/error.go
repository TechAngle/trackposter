// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package msg

import (
	"html"

	"trackposter/internal/domain"
	"trackposter/internal/pkg/pool"
)

// errorTemplate returns default template for error messages.
//
// Template look: ERROR: {message}.
func errorTemplate(message string) domain.MessageTemplate {
	s := pool.GetBuilder()
	defer pool.PutBuilder(s)

	s.Grow(len(message)*2 + 32)
	s.WriteString("<b>ERROR</b>: <i>")
	s.WriteString(html.EscapeString(message))
	s.WriteString("</i>")

	return domain.MessageTemplate(s.String())
}

func ErrorInternal(err error) domain.MessageTemplate {
	s := pool.GetBuilder()
	defer pool.PutBuilder(s)

	s.WriteString("Internal error:\n")
	s.WriteString(err.Error())

	return errorTemplate(s.String())
}

// ErrorEmptyQueue template when tracks queue was found empty.
func ErrorEmptyQueue() domain.MessageTemplate {
	return errorTemplate("Queue is empty.")
}

// ErrorTrackNotFound template when track was not found on platform.
func ErrorTrackNotFound() domain.MessageTemplate {
	return errorTemplate("Track not found.")
}

// ErrorInvalidURL template when user has provided an invalid URL.
func ErrorInvalidURL() domain.MessageTemplate {
	return errorTemplate("Invalid URL.")
}

// ErrorDownload template when connector returned an error.
func ErrorDownload(err error) domain.MessageTemplate {
	s := pool.GetBuilder()
	defer pool.PutBuilder(s)

	s.WriteString("Some track download failed due to error:\n")
	s.WriteString(err.Error())

	return errorTemplate(s.String())
}
