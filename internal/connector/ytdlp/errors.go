package ytdlp

import "errors"

var (
	ErrYtDlpCommand  = errors.New("yt-dlp command err")
	ErrMissingYtDlp  = errors.New("missing yt-dlp path")
	ErrMissingFFMpeg = errors.New("missing ffmpeg path")
	ErrLookPath      = errors.New("look path err")
	ErrUnmarshal     = errors.New("unmarshal err")
)
