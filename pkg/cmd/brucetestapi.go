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

var clientGetFoo = cli.Command{
	Name:            "get_foo",
	Usage:           "Get a Foo that has text, a random number, and a list of random numbers.",
	Flags:           []cli.Flag{},
	Action:          handleClientGetFoo,
	HideHelpCommand: true,
}

var clientSetText = cli.Command{
	Name:  "set_text",
	Usage: "Set the text that is returned when getting a Foo.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "name",
			},
		},
	},
	Action:          handleClientSetText,
	HideHelpCommand: true,
}

func handleClientGetFoo(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res, err := cc.client.GetFoo(context.TODO(), option.WithMiddleware(cc.AsMiddleware()))
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", ColorizeJSON(res.RawJSON(), os.Stdout))
	return nil
}

func handleClientSetText(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.SetTextParams{}
	res := []byte{}
	_, err := cc.client.SetText(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", ColorizeJSON(string(res), os.Stdout))
	return nil
}
