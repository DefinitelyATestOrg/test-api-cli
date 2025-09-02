// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/stainless-sdks/bruce-test-api-cli/pkg/jsonflag"
	"github.com/urfave/cli/v3"
)

var webhooksRegister = cli.Command{
	Name:  "register",
	Usage: "Register Webhook",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "url",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "url",
			},
		},
	},
	Action:          handleWebhooksRegister,
	HideHelpCommand: true,
}

func handleWebhooksRegister(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.WebhookRegisterParams{}
	res, err := cc.client.Webhooks.Register(
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
