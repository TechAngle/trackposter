// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package domain

// This package defines constant used by Connector and its options.

type AudioFormat string

// Audio format
const (
	MP3  AudioFormat = "mp3"
	M4A  AudioFormat = "m4a"
	WAV  AudioFormat = "wav"
	FLAC AudioFormat = "flac"
	AAC  AudioFormat = "aac"
)
