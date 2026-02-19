// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestYtDlpFound(t *testing.T) {
	ytDlpPath, err := ytDlpCommand()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", ytDlpPath)

	t.Logf("yt-dlp path: %s", ytDlpPath)
}

func TestFFMpegFound(t *testing.T) {
	ffmpegPath, err := ffmpegCommad()

	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", ffmpegPath)

	t.Logf("ffmpeg path: %s", ffmpegPath)
}
