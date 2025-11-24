// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/bruce-test-api-go"
	"github.com/stainless-sdks/bruce-test-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var foosList = cli.Command{
	Name:  "list",
	Usage: "Get foos",
	Flags: []cli.Flag{
		&cli.Int64Flag{
			Name:  "page",
			Usage: "Page number",
			Value: 1,
		},
		&cli.Int64Flag{
			Name:  "size",
			Usage: "Page size",
			Value: 50,
		},
	},
	Action:          handleFoosList,
	HideHelpCommand: true,
}

func handleFoosList(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.FooListParams{}
	if cmd.IsSet("page") {
		params.Page = brucetestapi.Opt(cmd.Value("page").(int64))
	}
	if cmd.IsSet("size") {
		params.Size = brucetestapi.Opt(cmd.Value("size").(int64))
	}
	var res []byte
	_, err := client.Foos.List(
		ctx,
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("foos list", json, format, transform)
}
