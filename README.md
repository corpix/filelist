go-filelist
-----------

Simple tool that creates filtered lists of files.

**IT IS IN ALPHA STATE NOW, API CHANGES ARE COMMING**

# Initial run
* Make sure you have typical golang directory structure
* Make sure your GOPATH is ok
* Make `go get github.com/corpix/go-filelist`
* Run `go-filelist -h` or `go run main.go -h` in `GOROOT/src/github.com/corpix/go-filelist`

It should print a help for you:

```
NAME:
   main - Create file list and filter it with globs

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.0-dev

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
$ go-filelist:master  cat filters.toml
Excludes = [
    "**/.git/**",
    "**/Godeps/**"
]

Includes = [
    "**.go"
]
```

Lets run go-filelist with this filter on `src` directory:
```bash
$ go-filelist:master  go run main.go filter --config=filters.toml src
/home/corpix/go/src/github.com/corpix/go-filelist/src
/home/corpix/go/src/github.com/corpix/go-filelist/src/cli
/home/corpix/go/src/github.com/corpix/go-filelist/src/cli/actions.go
/home/corpix/go/src/github.com/corpix/go-filelist/src/cli/cli.go
/home/corpix/go/src/github.com/corpix/go-filelist/src/cli/commands.go
/home/corpix/go/src/github.com/corpix/go-filelist/src/cli/flags.go
/home/corpix/go/src/github.com/corpix/go-filelist/src/config
/home/corpix/go/src/github.com/corpix/go-filelist/src/config/config.go
/home/corpix/go/src/github.com/corpix/go-filelist/src/matcher
/home/corpix/go/src/github.com/corpix/go-filelist/src/matcher/matcher.go
```
