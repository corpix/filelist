package cli

import (
	"github.com/urfave/cli"
)

var (
	commands = []cli.Command{
		{
			Name:      "filter",
			ShortName: "f",
			Usage:     "Filter files in a directory with specified globbing configuration",
			Flags:     filterFlags,
			Action:    filterAction,
		},
	}
)
