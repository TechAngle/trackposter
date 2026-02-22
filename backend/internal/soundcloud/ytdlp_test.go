// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package soundcloud_test

import (
	"context"
	"log"
	"testing"
	"trackposter/internal/soundcloud/ytdlp"
	"trackposter/internal/utils"

	"github.com/stretchr/testify/assert"
)

var (
	connector     *ytdlp.YtDlpConnector
	options       *ytdlp.YtDlpConnectorOptions
	ctx           context.Context
	cancelContext context.CancelFunc
)

const (
	// Using MGR soundtrack for test :3
	testURL = "https://soundcloud.com/shadows-9/the-hot-wind-blowing"
)

func createContext() {
	ctx, cancelContext = context.WithCancel(context.Background())
}

func TestDefaultOptions(t *testing.T) {
	// get default options
	defaultOptions, err := ytdlp.DefaultOptions()
	assert.NoError(t, err, "cannot get default options: %v", err)

	options = defaultOptions
}

func TestCreateConnector(t *testing.T) {
	// skipping if its CI enviroment
	utils.SkipIfCI(t)

	createContext()

	client, err := ytdlp.NewSoundcloudConnector(ctx, options)
	if err != nil {
		cancelContext()
		log.Panicf("failed to create yt-dlp soundcloud connector: %v\n", err)
	}

	connector = client
}

func TestTrackMetadataFromURL(t *testing.T) {
	// skipping if its CI enviroment
	utils.SkipIfCI(t)

	assert.NotNil(t, connector)

	metadata, err := connector.TrackMetadataFromURL(ctx, testURL)
	assert.NoError(t, err)

	t.Logf("%v", metadata)
}

func TestTrackDataFromURL(t *testing.T) {
	// skipping if its CI enviroment
	utils.SkipIfCI(t)

	assert.NotNil(t, connector)

	_, err := connector.TrackFromURL(ctx, testURL)
	assert.NoError(t, err)

	// write to the file
	// err = os.WriteFile("test.m4a", content, os.ModePerm)
	// assert.NoError(t, err)
}
