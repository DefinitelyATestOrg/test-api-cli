// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/bruce-test-api-go"
	"github.com/stainless-sdks/bruce-test-api-go/option"
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
		&cli.StringFlag{
			Name:  "species",
			Usage: "The species of the pet",
			Value: "Unknown",
		},
	},
	Action:          handlePeoplePetsCreate,
	HideHelpCommand: true,
}

var peoplePetsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get a pet from a person.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "person-id",
			Usage: "The unique identifier of the person to update",
		},
		&cli.StringFlag{
			Name:  "pet-name",
			Usage: "The pet's name",
		},
	},
	Action:          handlePeoplePetsRetrieve,
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
		&cli.StringFlag{
			Name:  "species",
			Usage: "The updated species of the pet",
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

func handlePeoplePetsCreate(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetNewParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"species": "species",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.People.Pets.New(
		ctx,
		cmd.Value("person-id").(string),
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
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

func handlePeoplePetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetGetParams{
		PetName: cmd.Value("pet-name").(string),
	}
	var res []byte
	_, err := client.People.Pets.Get(
		ctx,
		cmd.Value("person-id").(string),
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("people:pets retrieve", json, format, transform)
}

func handlePeoplePetsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetUpdateParams{
		PersonID: cmd.Value("person-id").(string),
	}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"species": "species",
	}, &params); err != nil {
		return err
	}
	if cmd.IsSet("person-id") {
		params.PersonID = cmd.Value("person-id").(string)
	}
	var res []byte
	_, err := client.People.Pets.Update(
		ctx,
		cmd.Value("pet-id").(string),
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
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
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := client.People.Pets.List(
		ctx,
		cmd.Value("person-id").(string),
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
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
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("pet-id") && len(unusedArgs) > 0 {
		cmd.Set("pet-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonPetDeleteParams{
		PersonID: cmd.Value("person-id").(string),
	}
	if cmd.IsSet("person-id") {
		params.PersonID = cmd.Value("person-id").(string)
	}
	var res []byte
	_, err := client.People.Pets.Delete(
		ctx,
		cmd.Value("pet-id").(string),
		params,
		option.WithMiddleware(debugMiddleware(cmd.Bool("debug"))),
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
