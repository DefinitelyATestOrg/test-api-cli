// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/bruce-test-api-cli/pkg/jsonflag"
	"github.com/stainless-sdks/bruce-test-api-go"
	"github.com/stainless-sdks/bruce-test-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var clientFnord = cli.Command{
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
		&jsonflag.JSONIntFlag{
			Name:  "first-query",
			Usage: "The first query param (required)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "first_query.#",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "+first-query",
			Usage: "The first query param (required)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "first_query.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "second-query",
			Usage: "The second query param (optional)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "second_query",
			},
		},
	},
	Action:          handleClientFnord,
	HideHelpCommand: true,
}

var clientPostFnord = cli.Command{
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
		&jsonflag.JSONIntFlag{
			Name:  "array-items",
			Usage: "The first query param (required)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "array_items.#",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "+array-item",
			Usage: "The first query param (required)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "array_items.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "name.full_name",
			Usage: "Full name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.full_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "name.nickname",
			Usage: "Nickname (if different from full name)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.nickname",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "second-query",
			Usage: "The second query param (optional)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "second_query",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "image-base64",
			Usage: "Image of the person (base64)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "image_base64",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "image-binary",
			Usage: "Image of the person (binary)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "image_binary",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "job",
			Usage: "The person's job",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "job",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "pets.name.full_name",
			Usage: "A list of pets for this person",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "pets.#.name.full_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "pets.name.nickname",
			Usage: "A list of pets for this person",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "pets.#.name.nickname",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "pets.species",
			Usage: "A list of pets for this person",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "pets.#.species",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "+pet",
			Usage: "A list of pets for this person",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "pets.-1",
				SetValue: map[string]interface{}{},
			},
		},
	},
	Action:          handleClientPostFnord,
	HideHelpCommand: true,
}

var clientTestForm = cli.Command{
	Name:  "test-form",
	Usage: "Test endpoint that accepts form-encoded data.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "email",
			Usage: "Email",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "email",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "username",
			Usage: "Username",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "username",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "age",
			Usage: "Age",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "age",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "subscribe",
			Usage: "Subscribe",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "subscribe",
				SetValue: true,
			},
			Value: false,
		},
	},
	Action:          handleClientTestForm,
	HideHelpCommand: true,
}

func handleClientFnord(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("second-pos") && len(unusedArgs) > 0 {
		cmd.Set("second-pos", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.FnordParams{}
	if cmd.IsSet("first-pos") {
		params.FirstPos = cmd.Value("first-pos").(string)
	}
	var res []byte
	_, err := cc.client.Fnord(
		ctx,
		cmd.Value("second-pos").(string),
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

func handleClientPostFnord(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("second-pos") && len(unusedArgs) > 0 {
		cmd.Set("second-pos", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PostFnordParams{}
	if cmd.IsSet("first-pos") {
		params.FirstPos = cmd.Value("first-pos").(string)
	}
	var res []byte
	_, err := cc.client.PostFnord(
		ctx,
		cmd.Value("second-pos").(string),
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
	return ShowJSON("client post-fnord", json, format, transform)
}

func handleClientTestForm(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.TestFormParams{}
	var res []byte
	_, err := cc.client.TestForm(
		ctx,
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
	return ShowJSON("client test-form", json, format, transform)
}
