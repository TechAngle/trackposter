package model

// AudioFormat represents type for audio formats used for downloading.
type AudioFormat string

// Audio format.
const (
	MP3  AudioFormat = "mp3"
	M4A  AudioFormat = "m4a"
	WAV  AudioFormat = "wav"
	FLAC AudioFormat = "flac"
	OPUS AudioFormat = "opus"
	AAC  AudioFormat = "aac"
)

// Formats returns slice of available audio formats.
func Formats() []AudioFormat {
	return []AudioFormat{
		MP3,
		M4A,
		WAV,
		FLAC,
		OPUS,
		AAC,
	}
}
