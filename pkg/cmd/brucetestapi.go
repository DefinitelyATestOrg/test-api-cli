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

var fnord = cli.Command{
	Name:  "fnord",
	Usage: "Test GET endpoint for positional and query params.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "first-pos",
			Usage: "The first positional arg",
		},
		&cli.StringFlag{
			Name:  "second-pos",
			Usage: "The second positional arg",
		},
		&cli.Int64SliceFlag{
			Name:  "first-query",
			Usage: "The first query param (required)",
		},
		&cli.StringFlag{
			Name:  "second-query",
			Usage: "The second query param (optional)",
		},
	},
	Action:          handleFnord,
	HideHelpCommand: true,
}

var postFnord = cli.Command{
	Name:  "post-fnord",
	Usage: "Test POST endpoint for positional and query params.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "first-pos",
			Usage: "The first positional arg",
		},
		&cli.StringFlag{
			Name:  "second-pos",
			Usage: "The second positional arg",
		},
		&cli.Int64SliceFlag{
			Name:  "array-item",
			Usage: "The first query param (required)",
		},
		&cli.StringFlag{
			Name:  "second-query",
			Usage: "The second query param (optional)",
		},
		&cli.GenericFlag{
			Name:      "image-base64",
			Usage:     "Image of the person (base64)",
			Value:     &fileReader{Base64Encoded: true},
			TakesFile: true,
		},
		&cli.GenericFlag{
			Name:      "image-binary",
			Usage:     "Image of the person (binary)",
			Value:     &fileReader{},
			TakesFile: true,
		},
		&cli.StringFlag{
			Name:  "job",
			Usage: "The person's job",
		},
	},
	Action:          handlePostFnord,
	HideHelpCommand: true,
}

var testForm = cli.Command{
	Name:  "test-form",
	Usage: "Test endpoint that accepts form-encoded data.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "email",
			Usage: "Email",
		},
		&cli.GenericFlag{
			Name:      "username",
			Usage:     "Username",
			Value:     &fileReader{Base64Encoded: true},
			TakesFile: true,
		},
		&cli.Int64Flag{
			Name:  "age",
			Usage: "Age",
		},
		&cli.BoolFlag{
			Name:  "subscribe",
			Usage: "Subscribe",
		},
	},
	Action:          handleTestForm,
	HideHelpCommand: true,
}

func handleFnord(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("second-pos") && len(unusedArgs) > 0 {
		cmd.Set("second-pos", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.FnordParams{
		FirstPos:   cmd.Value("first-pos").(string),
		FirstQuery: cmd.Value("first-query").([]int64),
	}
	if cmd.IsSet("second-query") {
		params.SecondQuery = brucetestapi.Opt(cmd.Value("second-query").(string))
	}
	if cmd.IsSet("first-pos") {
		params.FirstPos = cmd.Value("first-pos").(string)
	}
	var res []byte
	_, err := client.Fnord(
		ctx,
		cmd.Value("second-pos").(string),
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
	return ShowJSON("fnord", json, format, transform)
}

func handlePostFnord(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("second-pos") && len(unusedArgs) > 0 {
		cmd.Set("second-pos", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PostFnordParams{
		FirstPos:   cmd.Value("first-pos").(string),
		ArrayItems: cmd.Value("array-item").([]int64),
	}
	if cmd.IsSet("second-query") {
		params.SecondQuery = brucetestapi.Opt(cmd.Value("second-query").(string))
	}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"image-base64": "image_base64",
		"image-binary": "image_binary",
		"job":          "job",
	}, &params); err != nil {
		return err
	}
	if cmd.IsSet("first-pos") {
		params.FirstPos = cmd.Value("first-pos").(string)
	}
	var res []byte
	_, err := client.PostFnord(
		ctx,
		cmd.Value("second-pos").(string),
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
	return ShowJSON("post-fnord", json, format, transform)
}

func handleTestForm(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.TestFormParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"email":     "email",
		"username":  "username",
		"age":       "age",
		"subscribe": "subscribe",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.TestForm(
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
	return ShowJSON("test-form", json, format, transform)
}
