// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package ytdlp

import (
	"testing"
	"trackposter/internal/utils"

	"github.com/stretchr/testify/assert"
)

// Tests that fail GH Actions builds because
// 	they check if yt-dlp and ffmpeg are installed, and gh actions don't have ones.
// Therefore we just skip them in tests.

func TestYtDlpFound(t *testing.T) {
	// skipping if its CI enviroment
	utils.SkipIfCI(t)

	ytDlpPath, err := ytDlpCommand()
	assert.NoError(t, err)
	assert.NotEqual(t, "", ytDlpPath)

	t.Logf("yt-dlp path: %s", ytDlpPath)
}

func TestFFMpegFound(t *testing.T) {
	// skipping if its CI enviroment
	utils.SkipIfCI(t)

	ffmpegPath, err := ffmpegCommad()
	assert.NoError(t, err)
	assert.NotEqual(t, "", ffmpegPath)

	t.Logf("ffmpeg path: %s", ffmpegPath)
}
