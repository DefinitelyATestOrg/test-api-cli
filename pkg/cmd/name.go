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

var nameSet = cli.Command{
	Name:  "set",
	Usage: "Set Name",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "name",
			},
		},
	},
	Action:          handleNameSet,
	HideHelpCommand: true,
}

func handleNameSet(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.NameSetParams{}
	res := []byte{}
	_, err := cc.client.Name.Set(
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
