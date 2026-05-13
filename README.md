# Track Poster
[//]: # "Badges"

![Telegram](https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

![Zed](https://img.shields.io/badge/zedindustries-084CCF.svg?style=for-the-badge&logo=zedindustries&logoColor=white)
![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white)
![FFmpeg](https://shields.io/badge/FFmpeg-%23171717.svg?logo=ffmpeg&style=for-the-badge&labelColor=171717&logoColor=5cb85c)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)
![Task](https://img.shields.io/badge/Task-2B2D34?style=for-the-badge&logo=task&logoColor=white)
![Go Report Card](https://goreportcard.com/badge/github.com/TechAngle/trackposter)
![License](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg?style=for-the-badge)


> **Go Telegram Bot** for automating sending tracks to *Telegram channel*. Uses `yt-dlp` for downloading music and `ffmpeg` for  audio conversion.

**Project status: *Work In Progress.***

## How to use
- Download prebuilt binary from [Releases](https://github.com/TechAngle/trackposter/releases) or build it manually (read in [Build](#Build) section).
- Create `.env` file near to `trackposter` binary with such fields:
```bash
# Telegram token (get it from @BotFather)
TOKEN=

# Telegram IDs that can use bot (separate with ', ')
ALLOWED_ID=
```
- Run binary and wait until you'll similar lines:

![Start lines](docs/start.jpg)

- Send any URL (that supported by `yt-dlp`) to bot you've created.
- Wait a bit while bot will download your track and upload to Telegram.

## Build
### Prerequisites
- *Cloned repository using:*
```bash
$ git clone https://github.com/TechAngle/trackposter  
$ cd ./trackposter
```
- Installed **Go 1.23+**
- Installed [**go-task**](https://taskfile.dev/)
- Installed **ffmpeg** and **yt-dlp**.

### How to build
```bash
$ task build
```

## Contributing
Check out [CONTRIBUTE](./CONTRIBUTE.md)

## LICENSE
Project is distributed under the [Mozilla Public License 2.0 (MPL)](./LICENSE) License
