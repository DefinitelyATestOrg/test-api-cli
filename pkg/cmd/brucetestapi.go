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

var listFoos = cli.Command{
	Name:  "list-foos",
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
	Action:          handleListFoos,
	HideHelpCommand: true,
}

var updateCount = cli.Command{
	Name:  "update-count",
	Usage: "Perform update-count operation",
	Flags: []cli.Flag{
		&requestflag.IntFlag{
			Name: "body",
			Config: requestflag.RequestConfig{
				BodyPath: "body",
			},
		},
	},
	Action:          handleUpdateCount,
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

func handleListFoos(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.ListFoosParams{}

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
		_, err = client.ListFoos(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "list-foos", obj, format, transform)
	} else {
		iter := client.ListFoosAutoPaging(ctx, params, options...)
		return streamOutput("list-foos", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "list-foos", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleUpdateCount(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.UpdateCountParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	return client.UpdateCount(ctx, params, options...)
}
