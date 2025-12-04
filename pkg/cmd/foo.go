// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/bruce-test-api-cli/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-cli/internal/requestflag"
	"github.com/stainless-sdks/bruce-test-api-go"
	"github.com/stainless-sdks/bruce-test-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var foosList = cli.Command{
	Name:  "list",
	Usage: "Get foos",
	Flags: []cli.Flag{
		&requestflag.IntFlag{
			Name:  "page",
			Usage: "Page number",
			Value: 1,
			Config: requestflag.RequestConfig{
				QueryPath: "page",
			},
		},
		&requestflag.IntFlag{
			Name:  "size",
			Usage: "Page size",
			Value: 50,
			Config: requestflag.RequestConfig{
				QueryPath: "size",
			},
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

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatRepeat,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Foos.List(
		ctx,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("foos list", json, format, transform)
}
