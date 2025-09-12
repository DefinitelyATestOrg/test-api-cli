// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/urfave/cli/v3"
)

var jsonTestRetrieve = cli.Command{
	Name:            "retrieve",
	Usage:           "Get a big JSON response for testing.",
	Flags:           []cli.Flag{},
	Action:          handleJsonTestRetrieve,
	HideHelpCommand: true,
}

func handleJsonTestRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.JsonTest.Get(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("json-test retrieve", string(res), format)
}
