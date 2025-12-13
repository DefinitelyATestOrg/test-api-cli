// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/DefinitelyATestOrg/test-api-cli/internal/apiquery"
	"github.com/DefinitelyATestOrg/test-api-cli/internal/requestflag"
	"github.com/DefinitelyATestOrg/test-api-go"
	"github.com/urfave/cli/v3"
)

var formTest = cli.Command{
	Name:  "form-test",
	Usage: "Mixed parameter types",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name: "version",
		},
		&requestflag.Flag[string]{
			Name: "user-id",
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:      "date",
			QueryPath: "date",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "datetime",
			QueryPath: "datetime",
		},
		&requestflag.Flag[requestflag.TimeValue]{
			Name:      "time",
			QueryPath: "time",
		},
		&requestflag.Flag[any]{
			Name:      "filter",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "blorp",
			BodyPath: "blorp",
		},
		&requestflag.Flag[any]{
			Name:     "preferences",
			BodyPath: "preferences",
		},
		&requestflag.Flag[[]string]{
			Name:       "x-flag",
			HeaderPath: "X-Flags",
		},
		&requestflag.Flag[string]{
			Name:       "x-trace-id",
			HeaderPath: "X-Trace-ID",
		},
	},
	Action:          handleFormTest,
	HideHelpCommand: true,
}

var jsonTest = cli.Command{
	Name:  "json-test",
	Usage: "Mixed parameter types",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name: "version",
		},
		&requestflag.Flag[string]{
			Name: "user-id",
		},
		&requestflag.Flag[requestflag.DateValue]{
			Name:      "date",
			QueryPath: "date",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "datetime",
			QueryPath: "datetime",
		},
		&requestflag.Flag[requestflag.TimeValue]{
			Name:      "time",
			QueryPath: "time",
		},
		&requestflag.Flag[any]{
			Name:      "filter",
			QueryPath: "filter",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			QueryPath: "limit",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "blorp",
			BodyPath: "blorp",
		},
		&requestflag.Flag[any]{
			Name:     "preferences",
			BodyPath: "preferences",
		},
		&requestflag.Flag[[]string]{
			Name:       "x-flag",
			HeaderPath: "X-Flags",
		},
		&requestflag.Flag[string]{
			Name:       "x-trace-id",
			HeaderPath: "X-Trace-ID",
		},
	},
	Action:          handleJsonTest,
	HideHelpCommand: true,
}

var updateCount = cli.Command{
	Name:  "update-count",
	Usage: "Perform update-count operation",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "body",
			BodyRoot: true,
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
		Version: cmd.Value("version").(int64),
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
		cmd.Value("user-id").(string),
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
		Version: cmd.Value("version").(int64),
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
		cmd.Value("user-id").(string),
		params,
		options...,
	)
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
