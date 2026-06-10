package msg

import (
	"strings"

	"trackposter/internal/model"
	"trackposter/internal/telegram/template"
)

const (
	averageInfoLen = 64
)

// TrackAdded template when track was added to the line and contains its ID.
func TrackAdded(trackID string) template.MessageTemplate {
	var builder strings.Builder
	builder.Grow(len(trackID) + averageInfoLen)
	builder.WriteString(
		`<b>Track was successfully added to the line with ID <code>`,
	)
	builder.WriteString(trackID)
	builder.WriteString(`</code>.</b>`)

	return template.MessageTemplate(builder.String())
}

// TracksQueue template for formatting tracks queue.
func TracksQueue(tracks []*model.Track) template.MessageTemplate {
	if len(tracks) == 0 {
		return `<b>Queue is empty.</b>`
	}

	var builder strings.Builder

	totalTracks := len(tracks)
	builderSize := averageInfoLen * totalTracks
	builder.Grow(builderSize)

	builder.WriteString(`<b>Current queue</b>:`)
	for _, track := range tracks {
		builder.WriteString(`- [`)
		builder.WriteString(track.Author)
		builder.WriteString(`] `)
		builder.WriteString(track.Title)
		builder.WriteString("\n")
	}

	return template.MessageTemplate(builder.String())
}

// MetadataSearchInfo template when looking for track metadata.
func MetadataSearchInfo() template.MessageTemplate {
	return `<b>Starting looking for track metadata...</b>`
}
