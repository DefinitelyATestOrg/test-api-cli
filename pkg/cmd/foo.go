// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/bruce-hill/bruce-test-api-cli/pkg/jsonflag"
	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/urfave/cli/v3"
)

var foosCreate = cli.Command{
	Name:  "create",
	Usage: "Add a Foo to the list of all Foos.",
	Flags: []cli.Flag{
		&jsonflag.JSONIntFlag{
			Name: "list-of-nums",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "list_of_nums.#",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "+list_of_num",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "list_of_nums.-1",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "random-number",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "random_number",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "text",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "text",
			},
		},
	},
	Action:          handleFoosCreate,
	HideHelpCommand: true,
}

var foosRetrieve = cli.Command{
	Name:            "retrieve",
	Usage:           "Get a Foo that has text, a random number, and a list of random numbers.",
	Flags:           []cli.Flag{},
	Action:          handleFoosRetrieve,
	HideHelpCommand: true,
}

var foosList = cli.Command{
	Name:  "list",
	Usage: "Get a list of all of the Foos.",
	Flags: []cli.Flag{
		&jsonflag.JSONIntFlag{
			Name: "page",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "size",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "size",
			},
		},
	},
	Action:          handleFoosList,
	HideHelpCommand: true,
}

func handleFoosCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.FooNewParams{}
	res := []byte{}
	_, err := cc.client.Foos.New(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("foos create", string(res), format)
}

func handleFoosRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res, err := cc.client.Foos.Get(context.TODO(), option.WithMiddleware(cc.AsMiddleware()))
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("foos retrieve", res.RawJSON(), format)
}

func handleFoosList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.FooListParams{}
	res, err := cc.client.Foos.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("foos list", res.RawJSON(), format)
}
