// Copyright TechAngle 2026. All rights reserved.
// Use of this source code is controlled by MPL-2.0 that could be found in LICENSE file.
//
// Author: https://github.com/TechAngle

package config

const (
	// MaxQueueSize is Telegram maximum amount of files per message.
	MaxQueueSize = 10

	// MaxFileSize is the maximum size of file (in megabytes) which can Telegram bot send.
	//
	// NOTE:
	// If we have 10 files each 50 mb saved in RAM then bot will eat 500 of RAM only for
	// their storage. So that Memory repository is not recommended for production or
	// low-end server.
	MaxFileSize = 50
)
