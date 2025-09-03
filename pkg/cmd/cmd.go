// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"github.com/urfave/cli/v3"
)

var Command = cli.Command{
	Name:    "bruce-test-api",
	Usage:   "CLI for the bruce-test-api API",
	Version: Version,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable debug logging",
		},
		&cli.StringFlag{
			Name:  "base-url",
			Usage: "Override the base URL for API requests",
		},
	},
	Commands: []*cli.Command{
		{
			Name:     "foo",
			Category: "API RESOURCE",
			Commands: []*cli.Command{
				&fooList,
			},
		},

		{
			Name:     "name",
			Category: "API RESOURCE",
			Commands: []*cli.Command{
				&nameSet,
			},
		},
	},
	EnableShellCompletion: true,
	HideHelpCommand:       true,
}
