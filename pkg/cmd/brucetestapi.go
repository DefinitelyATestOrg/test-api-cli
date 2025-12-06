// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/stainless-sdks/bruce-test-api-cli/internal/apiquery"
	"github.com/stainless-sdks/bruce-test-api-cli/internal/requestflag"
	"github.com/stainless-sdks/bruce-test-api-go"
	"github.com/stainless-sdks/bruce-test-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var formTest = cli.Command{
	Name:  "form-test",
	Usage: "Mixed parameter types",
	Flags: []cli.Flag{
		&requestflag.IntFlag{
			Name: "version",
		},
		&requestflag.StringFlag{
			Name: "user-id",
		},
		&requestflag.DateFlag{
			Name: "date",
			Config: requestflag.RequestConfig{
				QueryPath: "date",
			},
		},
		&requestflag.DateTimeFlag{
			Name: "datetime",
			Config: requestflag.RequestConfig{
				QueryPath: "datetime",
			},
		},
		&requestflag.TimeFlag{
			Name: "time",
			Config: requestflag.RequestConfig{
				QueryPath: "time",
			},
		},
		&requestflag.YAMLFlag{
			Name: "filter",
			Config: requestflag.RequestConfig{
				QueryPath: "filter",
			},
		},
		&requestflag.IntFlag{
			Name: "limit",
			Config: requestflag.RequestConfig{
				QueryPath: "limit",
			},
		},
		&requestflag.StringSliceFlag{
			Name: "tag",
			Config: requestflag.RequestConfig{
				QueryPath: "tags",
			},
		},
		&requestflag.StringFlag{
			Name: "blorp",
			Config: requestflag.RequestConfig{
				BodyPath: "blorp",
			},
		},
		&requestflag.YAMLFlag{
			Name: "preferences",
			Config: requestflag.RequestConfig{
				BodyPath: "preferences",
			},
		},
		&requestflag.StringSliceFlag{
			Name: "x-flag",
			Config: requestflag.RequestConfig{
				HeaderPath: "X-Flags",
			},
		},
		&requestflag.StringFlag{
			Name: "x-trace-id",
			Config: requestflag.RequestConfig{
				HeaderPath: "X-Trace-ID",
			},
		},
	},
	Action:          handleFormTest,
	HideHelpCommand: true,
}

var jsonTest = cli.Command{
	Name:  "json-test",
	Usage: "Mixed parameter types",
	Flags: []cli.Flag{
		&requestflag.IntFlag{
			Name: "version",
		},
		&requestflag.StringFlag{
			Name: "user-id",
		},
		&requestflag.DateFlag{
			Name: "date",
			Config: requestflag.RequestConfig{
				QueryPath: "date",
			},
		},
		&requestflag.DateTimeFlag{
			Name: "datetime",
			Config: requestflag.RequestConfig{
				QueryPath: "datetime",
			},
		},
		&requestflag.TimeFlag{
			Name: "time",
			Config: requestflag.RequestConfig{
				QueryPath: "time",
			},
		},
		&requestflag.YAMLFlag{
			Name: "filter",
			Config: requestflag.RequestConfig{
				QueryPath: "filter",
			},
		},
		&requestflag.IntFlag{
			Name: "limit",
			Config: requestflag.RequestConfig{
				QueryPath: "limit",
			},
		},
		&requestflag.StringSliceFlag{
			Name: "tag",
			Config: requestflag.RequestConfig{
				QueryPath: "tags",
			},
		},
		&requestflag.StringFlag{
			Name: "blorp",
			Config: requestflag.RequestConfig{
				BodyPath: "blorp",
			},
		},
		&requestflag.YAMLFlag{
			Name: "preferences",
			Config: requestflag.RequestConfig{
				BodyPath: "preferences",
			},
		},
		&requestflag.StringSliceFlag{
			Name: "x-flag",
			Config: requestflag.RequestConfig{
				HeaderPath: "X-Flags",
			},
		},
		&requestflag.StringFlag{
			Name: "x-trace-id",
			Config: requestflag.RequestConfig{
				HeaderPath: "X-Trace-ID",
			},
		},
	},
	Action:          handleJsonTest,
	HideHelpCommand: true,
}

var paginatedTest = cli.Command{
	Name:  "paginated-test",
	Usage: "Get foos",
	Flags: []cli.Flag{
		&requestflag.IntFlag{
			Name:  "page",
			Usage: "Page number",
			Value: requestflag.Value[int64](1),
			Config: requestflag.RequestConfig{
				QueryPath: "page",
			},
		},
		&requestflag.IntFlag{
			Name:  "size",
			Usage: "Page size",
			Value: requestflag.Value[int64](50),
			Config: requestflag.RequestConfig{
				QueryPath: "size",
			},
		},
		&requestflag.StringSliceFlag{
			Name: "tag",
			Config: requestflag.RequestConfig{
				QueryPath: "tags",
			},
		},
	},
	Action:          handlePaginatedTest,
	HideHelpCommand: true,
}

func handleFormTest(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.FormTestParams{
		Version: requestflag.CommandRequestValue[int64](cmd, "version"),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
	)
	if err != nil {
		return err
	}

	return client.FormTest(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "user-id"),
		params,
		options...,
	)
}

func handleJsonTest(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("user-id") && len(unusedArgs) > 0 {
		cmd.Set("user-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.JsonTestParams{
		Version: requestflag.CommandRequestValue[int64](cmd, "version"),
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	return client.JsonTest(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "user-id"),
		params,
		options...,
	)
}

func handlePaginatedTest(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PaginatedTestParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.PaginatedTest(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "paginated-test", obj, format, transform)
}
