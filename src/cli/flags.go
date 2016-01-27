package cli

import (
	"github.com/codegangsta/cli"
)

var (
	rootFlags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug mode",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: "log level(debug, info, warn, error, fatal, panic)",
		},
	}
	filterFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "filters.toml",
			Usage: "filters that describes includes and excludes",
		},
	}
)
