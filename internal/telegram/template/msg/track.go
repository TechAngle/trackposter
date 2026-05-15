package msg

import (
	"strings"

	"trackposter/internal/domain"
)

const (
	averageInfoLen = 64
)

// TrackAdded template when track was added to the line and contains its ID.
func TrackAdded(trackID string) domain.MessageTemplate {
	var builder strings.Builder
	builder.Grow(len(trackID) + averageInfoLen)
	builder.WriteString(
		`<b>Track was successfully added to the line with ID <code>`,
	)
	builder.WriteString(trackID)
	builder.WriteString(`</code>.</b>`)

	return domain.MessageTemplate(builder.String())
}

// TracksQueue template for formatting tracks queue.
func TracksQueue(tracks []*domain.Track) domain.MessageTemplate {
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

	return domain.MessageTemplate(builder.String())
}

// MetadataSearchInfo template when looking for track metadata.
func MetadataSearchInfo() domain.MessageTemplate {
	return `<b>Starting looking for track metadata...</b>`
}
