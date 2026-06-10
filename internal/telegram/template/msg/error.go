package msg

import (
	"html"

	"trackposter/internal/pkg/pool"
	"trackposter/internal/telegram/template"
)

const (
	msgLenMultiplier     = 2
	additionalContentLen = 32
)

// errorTemplate returns default template for error messages.
//
// Template look: ERROR: {message}.
func errorTemplate(message string) template.MessageTemplate {
	builder := pool.GetBuilder()
	defer pool.PutBuilder(builder)

	builder.Grow(len(message)*msgLenMultiplier + additionalContentLen)
	builder.WriteString("<b>ERROR</b>: <i>")
	builder.WriteString(html.EscapeString(message))
	builder.WriteString("</i>")

	return template.MessageTemplate(builder.String())
}

// ErrorInternal template when non-related to Telegram error occurred.
func ErrorInternal(err error) template.MessageTemplate {
	builder := pool.GetBuilder()
	defer pool.PutBuilder(builder)

	builder.WriteString("Internal error:\n")
	builder.WriteString(err.Error())

	return errorTemplate(builder.String())
}

// ErrorEmptyQueue template when tracks queue was found empty.
func ErrorEmptyQueue() template.MessageTemplate {
	return errorTemplate("Queue is empty.")
}

// ErrorTrackNotFound template when track was not found on platform.
func ErrorTrackNotFound() template.MessageTemplate {
	return errorTemplate("Track not found.")
}

// ErrorInvalidURL template when user has provided an invalid URL.
func ErrorInvalidURL() template.MessageTemplate {
	return errorTemplate("Invalid URL.")
}

// ErrorDownload template when connector returned an error.
func ErrorDownload(err error) template.MessageTemplate {
	builder := pool.GetBuilder()
	defer pool.PutBuilder(builder)

	builder.WriteString("Some track download failed due to error:\n")
	builder.WriteString(err.Error())

	return errorTemplate(builder.String())
}
