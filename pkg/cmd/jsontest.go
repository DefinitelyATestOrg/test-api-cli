// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var jsonTestRetrieve = cli.Command{
	Name:            "retrieve",
	Usage:           "Get a big JSON response for testing.",
	Flags:           []cli.Flag{},
	Action:          handleJsonTestRetrieve,
	HideHelpCommand: true,
}

func handleJsonTestRetrieve(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.JsonTest.Get(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("json-test retrieve", json, format, transform)
}
