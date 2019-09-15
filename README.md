# colorsizes

Reads `du` output and colorises every three digits of the sizes differently so
they stand out better.

## Install

```bash
go get github.com/anzdaddy/colorsizes
```

## Use

```
du -b | sort -n | tail -n $(( $(tput lines) - 2 )) | colorsizes
```
