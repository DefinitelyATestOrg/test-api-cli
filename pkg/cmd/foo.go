// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/DefinitelyATestOrg/test-api-cli/internal/apiquery"
	"github.com/DefinitelyATestOrg/test-api-cli/internal/requestflag"
	"github.com/DefinitelyATestOrg/test-api-go"
	"github.com/DefinitelyATestOrg/test-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var foosList = cli.Command{
	Name:  "list",
	Usage: "Get foos",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page",
			Usage:     "Page number",
			Default:   1,
			QueryPath: "page",
		},
		&requestflag.Flag[int64]{
			Name:      "size",
			Usage:     "Page size",
			Default:   50,
			QueryPath: "size",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			QueryPath: "tags",
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
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Foos.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "foos list", obj, format, transform)
	} else {
		iter := client.Foos.ListAutoPaging(ctx, params, options...)
		return streamOutput("foos list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "foos list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}
