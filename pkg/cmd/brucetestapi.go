// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/urfave/cli/v3"
)

var clientJsonTest = cli.Command{
	Name:            "json-test",
	Usage:           "Get a big JSON response for testing.",
	Flags:           []cli.Flag{},
	Action:          handleClientJsonTest,
	HideHelpCommand: true,
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

	format := cmd.Root().String("format")
	return ShowJSON("client json-test", string(res), format)
}
