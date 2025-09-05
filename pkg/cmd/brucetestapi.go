// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/urfave/cli/v3"
)

var clientGetFoo = cli.Command{
	Name:            "get-foo",
	Usage:           "Get a Foo that has text, a random number, and a list of random numbers.",
	Flags:           []cli.Flag{},
	Action:          handleClientGetFoo,
	HideHelpCommand: true,
}

var clientJsonTest = cli.Command{
	Name:            "json-test",
	Usage:           "Get a big JSON response for testing.",
	Flags:           []cli.Flag{},
	Action:          handleClientJsonTest,
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

func handleClientJsonTest(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.JsonTest(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", ColorizeJSON(string(res), os.Stdout))
	return nil
}
