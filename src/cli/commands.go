package cli

import (
	"github.com/codegangsta/cli"
)

var (
	commands = []cli.Command{
		{
			Name:      "filter",
			ShortName: "f",
			Usage:     "filter files in a directory with specified globbing configuration",
			Flags:     filterFlags,
			Action:    filterAction,
		},
	}
)
