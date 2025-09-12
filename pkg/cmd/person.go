// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/stainless-sdks/bruce-test-api-cli/pkg/jsonflag"
	"github.com/urfave/cli/v3"
)

var peopleCreate = cli.Command{
	Name:  "create",
	Usage: "Create a new person and add them to the system.",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "name.full",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.full",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name.nick",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.nick",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "pet-names.full",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "pet_names.#.full",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "pet-names.nick",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "pet_names.#.nick",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "+pet_name",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "pet_names.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
	},
	Action:          handlePeopleCreate,
	HideHelpCommand: true,
}

var peopleRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get a person's information by ID.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "person-id",
		},
	},
	Action:          handlePeopleRetrieve,
	HideHelpCommand: true,
}

var peopleUpdate = cli.Command{
	Name:  "update",
	Usage: "Update an existing person's information.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "person-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "full",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "full",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "nick",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "nick",
			},
		},
	},
	Action:          handlePeopleUpdate,
	HideHelpCommand: true,
}

var peopleList = cli.Command{
	Name:  "list",
	Usage: "Get a list of all people.",
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
	Action:          handlePeopleList,
	HideHelpCommand: true,
}

var peopleDelete = cli.Command{
	Name:  "delete",
	Usage: "Remove a person from the system.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "person-id",
		},
	},
	Action:          handlePeopleDelete,
	HideHelpCommand: true,
}

func handlePeopleCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.PersonNewParams{}
	res, err := cc.client.People.New(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("people create", res.RawJSON(), format)
}

func handlePeopleRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res, err := cc.client.People.Get(
		context.TODO(),
		cmd.Value("person-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("people retrieve", res.RawJSON(), format)
}

func handlePeopleUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.PersonUpdateParams{}
	res, err := cc.client.People.Update(
		context.TODO(),
		cmd.Value("person-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("people update", res.RawJSON(), format)
}

func handlePeopleList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.PersonListParams{}
	res, err := cc.client.People.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("people list", res.RawJSON(), format)
}

func handlePeopleDelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.People.Delete(
		context.TODO(),
		cmd.Value("person-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("people delete", string(res), format)
}
