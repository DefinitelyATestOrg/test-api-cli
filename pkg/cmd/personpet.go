// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

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
			Name:  "person-id",
			Usage: "The unique identifier of the person to add a pet to",
		},
		&jsonflag.JSONStringFlag{
			Name:  "name.full_name",
			Usage: "Full name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.full_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "name.nickname",
			Usage: "Nickname (if different from full name)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.nickname",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "species",
			Usage: "The species of the pet",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "species",
			},
			Value: "Unknown",
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
			Name:  "person-id",
			Usage: "The unique identifier of the person who owns the pet",
		},
		&cli.StringFlag{
			Name:  "pet-id",
			Usage: "The unique identifier of the pet to update",
		},
		&jsonflag.JSONStringFlag{
			Name:  "name.full_name",
			Usage: "Full name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.full_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "name.nickname",
			Usage: "Nickname (if different from full name)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name.nickname",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "species",
			Usage: "The updated species of the pet",
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
			Name:  "person-id",
			Usage: "The unique identifier of the person whose pets to retrieve",
		},
	},
	Action:          handlePeoplePetsList,
	HideHelpCommand: true,
}

var peoplePetsDelete = cli.Command{
	Name:  "delete",
	Usage: "Remove a pet from a person.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "person-id",
			Usage: "The unique identifier of the person who owns the pet",
		},
		&cli.StringFlag{
			Name:  "pet-id",
			Usage: "The unique identifier of the pet to delete",
		},
	},
	Action:          handlePeoplePetsDelete,
	HideHelpCommand: true,
}

var peoplePetsFnord = cli.Command{
	Name:  "fnord",
	Usage: "Get a pet from a person.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "person-id",
			Usage: "The unique identifier of the person to update",
		},
		&jsonflag.JSONStringFlag{
			Name:  "pet-name",
			Usage: "The pet's name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "pet_name",
			},
		},
	},
	Action:          handlePeoplePetsFnord,
	HideHelpCommand: true,
}

var peoplePetsFrob = cli.Command{
	Name:  "frob",
	Usage: "Get a pet from a person.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "person-id",
			Usage: "The unique identifier of the person to update",
		},
		&jsonflag.JSONStringFlag{
			Name:  "pet-name",
			Usage: "The pet's name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "pet_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "frob",
			Usage: "The pet's frob",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "frob",
			},
		},
	},
	Action:          handlePeoplePetsFrob,
	HideHelpCommand: true,
}

func handlePeoplePetsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetNewParams{}
	var res []byte
	_, err := cc.client.People.Pets.New(
		ctx,
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
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetUpdateParams{}
	if cmd.IsSet("person-id") {
		params.PersonID = cmd.Value("person-id").(string)
	}
	var res []byte
	_, err := cc.client.People.Pets.Update(
		ctx,
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
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.People.Pets.List(
		ctx,
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

func handlePeoplePetsDelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetDeleteParams{}
	if cmd.IsSet("person-id") {
		params.PersonID = cmd.Value("person-id").(string)
	}
	var res []byte
	_, err := cc.client.People.Pets.Delete(
		ctx,
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
	return ShowJSON("people:pets delete", json, format, transform)
}

func handlePeoplePetsFnord(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetFnordParams{}
	var res []byte
	_, err := cc.client.People.Pets.Fnord(
		ctx,
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
	return ShowJSON("people:pets fnord", json, format, transform)
}

func handlePeoplePetsFrob(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetFrobParams{}
	var res []byte
	_, err := cc.client.People.Pets.Frob(
		ctx,
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
	return ShowJSON("people:pets frob", json, format, transform)
}
