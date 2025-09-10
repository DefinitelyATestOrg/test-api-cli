// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bruce-hill/bruce-test-api-cli/pkg/jsonflag"
	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/urfave/cli/v3"
)

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

func handleFoosRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res, err := cc.client.Foos.Get(context.TODO(), option.WithMiddleware(cc.AsMiddleware()))
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", ColorizeJSON(res.RawJSON(), os.Stdout))
	return nil
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

	fmt.Printf("%s\n", ColorizeJSON(res.RawJSON(), os.Stdout))
	return nil
}
