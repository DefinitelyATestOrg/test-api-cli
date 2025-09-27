// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/bruce-hill/bruce-test-api-go"
	"github.com/bruce-hill/bruce-test-api-go/option"
	"github.com/stainless-sdks/bruce-test-api-cli/pkg/jsonflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var peoplePetsCreate = cli.Command{
	Name:  "create",
	Usage: "Add a new pet to an existing person.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "person-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "name.full_name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.full_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name.nickname",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.nickname",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "species",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "species",
			},
		},
	},
	Action:          handlePeoplePetsCreate,
	HideHelpCommand: true,
}

var peoplePetsUpdate = cli.Command{
	Name:  "update",
	Usage: "Update an existing pet's information.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "person-id",
		},
		&cli.StringFlag{
			Name: "pet-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "name.full_name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.full_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name.nickname",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.nickname",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "species",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "species",
			},
		},
	},
	Action:          handlePeoplePetsUpdate,
	HideHelpCommand: true,
}

var peoplePetsList = cli.Command{
	Name:  "list",
	Usage: "Get all pets belonging to a specific person by their ID.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "person-id",
		},
	},
	Action:          handlePeoplePetsList,
	HideHelpCommand: true,
}

var peoplePetsQdelete = cli.Command{
	Name:  "qdelete",
	Usage: "Remove a pet from a person.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "person-id",
		},
		&cli.StringFlag{
			Name: "pet-id",
		},
	},
	Action:          handlePeoplePetsQdelete,
	HideHelpCommand: true,
}

func handlePeoplePetsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.PersonPetNewParams{}
	var res []byte
	_, err := cc.client.People.Pets.New(
		context.TODO(),
		cmd.Value("person-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("people:pets create", json, format, transform)
}

func handlePeoplePetsUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.PersonPetUpdateParams{}
	if cmd.IsSet("person-id") {
		params.PersonID = cmd.Value("person-id").(string)
	}
	var res []byte
	_, err := cc.client.People.Pets.Update(
		context.TODO(),
		cmd.Value("pet-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("people:pets update", json, format, transform)
}

func handlePeoplePetsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.People.Pets.List(
		context.TODO(),
		cmd.Value("person-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("people:pets list", json, format, transform)
}

func handlePeoplePetsQdelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := brucetestapi.PersonPetQdeleteParams{}
	if cmd.IsSet("person-id") {
		params.PersonID = cmd.Value("person-id").(string)
	}
	var res []byte
	_, err := cc.client.People.Pets.Qdelete(
		context.TODO(),
		cmd.Value("pet-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("people:pets qdelete", json, format, transform)
}
