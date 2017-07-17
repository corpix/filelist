package cli

import (
	"github.com/urfave/cli"
)

var (
	rootFlags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "Debug mode",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:  "log-level, l",
			Value: "info",
			Usage: "Log level(debug, info, warn, error, fatal, panic)",
		},
	}
	filterFlags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "filters.toml",
			Usage: "Filters that describes includes and excludes",
		},
	}
)
