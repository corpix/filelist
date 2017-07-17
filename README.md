filelist
--------

[![Build Status](https://travis-ci.org/corpix/filelist.svg?branch=master)](https://travis-ci.org/corpix/filelist)

Simple tool that creates filtered lists of files.

# Initial run
* Make sure you have typical golang directory structure
* Make sure your GOPATH is ok
* Make `go get github.com/corpix/filelist`
* Run `filelist -h` or `go run main.go -h` in `GOROOT/src/github.com/corpix/filelist`

It should print a help for you:

```
NAME:
   main - Create file list and filter it with globs

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   development

AUTHOR(S):
   Dmitry Moskowski <me@corpix.ru>

COMMANDS:
   filter, f	Filter files in a directory with specified globbing configuration
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --debug			Debug mode [$DEBUG]
   --log-level, -l "info"	Log level(debug, info, warn, error, fatal, panic)
   --help, -h			show help
   --version, -v		print the version
```

## Filtering
You need a config with the filters. You could find one in the repository root:
```
$ cat filters.toml
Excludes = [
    "**/.git/**",
    "**/Godeps/**"
]

Includes = [
    "**.go"
]
```

Lets run `filelist` with this filter on `src` directory:
```bash
$ go run filelist/filelist.go filter --config=filters.toml src
/home/corpix/go/src/github.com/corpix/filelist
/home/corpix/go/src/github.com/corpix/filelist/cli
/home/corpix/go/src/github.com/corpix/filelist/cli/actions.go
/home/corpix/go/src/github.com/corpix/filelist/cli/cli.go
/home/corpix/go/src/github.com/corpix/filelist/cli/commands.go
/home/corpix/go/src/github.com/corpix/filelist/cli/flags.go
/home/corpix/go/src/github.com/corpix/filelist/config
/home/corpix/go/src/github.com/corpix/filelist/config/config.go
/home/corpix/go/src/github.com/corpix/filelist/matcher
/home/corpix/go/src/github.com/corpix/filelist/matcher/matcher.go
```
