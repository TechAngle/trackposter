// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package msg

import (
	"strings"

	"trackposter/internal/domain"
)

// TrackAdded template when track was added to the line and contains its ID.
func TrackAdded(trackID string) domain.MessageTemplate {
	var s strings.Builder
	s.Grow(128)
	s.WriteString(`<b>Track was successfully added to the line with ID <code>`)
	s.WriteString(trackID)
	s.WriteString(`</code>.</b>`)

	return domain.MessageTemplate(s.String())
}

func TracksQueue(tracks []*domain.Track) domain.MessageTemplate {
	var s strings.Builder
	s.Grow(348)

	s.WriteString(`<b>Current queue</b>:`)
	for _, track := range tracks {
		s.WriteString(`- [`)
		s.WriteString(track.Author)
		s.WriteString(`] `)
		s.WriteString(track.URL)
		s.WriteString("\n")
	}

	return domain.MessageTemplate(s.String())
}

// MetadataSearchInfo template when looking for track metadata.
func MetadataSearchInfo() domain.MessageTemplate {
	return `<b>Starting looking for track metadata...</b>`
}
