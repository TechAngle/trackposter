// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package msg

import (
	"fmt"
	"trackposter/internal/domain"
)

// TrackAdded template when track was added to the line and contains its ID.
func TrackAdded(lineID int) domain.MessageTemplate {
	return domain.MessageTemplate(
		fmt.Sprintf(
			`<b>Track was successfully added to the line with ID %d.</b>`,
			lineID,
		),
	)
}
