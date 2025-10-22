// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/stainless-sdks/bruce-test-api-cli/pkg/jsonflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var clientFnord = cli.Command{
	Name:  "fnord",
	Usage: "Get a pet from a person.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "pos1",
			Usage: "The first positional arg",
		},
		&cli.StringFlag{
			Name:  "pos2",
			Usage: "The second positional arg",
		},
		&jsonflag.JSONStringFlag{
			Name:  "query1",
			Usage: "The first query param (required)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "query1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "query2",
			Usage: "The second query param (optional)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "query2",
			},
		},
	},
	Action:          handleClientFnord,
	HideHelpCommand: true,
}

func handleClientFnord(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pos2") && len(unusedArgs) > 0 {
		cmd.Set("pos2", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.FnordParams{}
	if cmd.IsSet("pos1") {
		params.Pos1 = cmd.Value("pos1").(string)
	}
	var res []byte
	_, err := cc.client.Fnord(
		ctx,
		cmd.Value("pos2").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("client fnord", json, format, transform)
}
