# Track Poster
> **Go + TS extension** for automating sending tracks to *Telegram channel*.

***WIP***

## Build
### Prerequisites
- *Cloned repository using*
```bash
git clone https://github.com/TechAngle/trackposter  
cd ./trackposter
```
**[For backend](#Backend)**:
- Installed **Go 1.23+**
- Installed [**go-task**](https://taskfile.dev/)
- Installed **ffmpeg** and **yt-dlp**.

**[For extension](#Extension)**:
- Installed *[bun](https://bun.sh/)*  

### Backend
```bash
cd ./backend
task build
```

### Extension
```bash
cd ./extension
bun install

bun build-firefox # to build an extension and compress it to the archive
bun dev # for development
```

## Contributing
Check out [CONTRIBUTE](./CONTRIBUTE.md)

## LICENSE
Project is distributed under the [Mozilla Public License 2.0 (MPL)](./LICENSE) License
