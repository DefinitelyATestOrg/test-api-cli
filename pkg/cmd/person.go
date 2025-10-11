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

var peopleCreate = cli.Command{
	Name:  "create",
	Usage: "Create a new person and add them to the system.",
	Flags: []cli.Flag{
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
			Name:  "job",
			Usage: "The person's job",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "job",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "pets.name.full_name",
			Usage: "A list of pets for this person",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "pets.#.name.full_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "pets.name.nickname",
			Usage: "A list of pets for this person",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "pets.#.name.nickname",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "pets.species",
			Usage: "A list of pets for this person",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "pets.#.species",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "+pet",
			Usage: "A list of pets for this person",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "pets.-1",
				SetValue: map[string]interface{}{},
			},
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
			Name:  "person-id",
			Usage: "The unique identifier of the person to retrieve",
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
			Name:  "person-id",
			Usage: "The unique identifier of the person to update",
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
			Name:  "job",
			Usage: "The updated job of the person",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "job",
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
		&jsonflag.JSONStringFlag{
			Name:  "job",
			Usage: "Job name to search for",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "job",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "name",
			Usage: "Full name to search for",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "nickname",
			Usage: "Nickname to search for",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "nickname",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "page",
			Usage: "Page number",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page",
			},
			Value: 1,
		},
		&jsonflag.JSONIntFlag{
			Name:  "size",
			Usage: "Page size",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "size",
			},
			Value: 50,
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
			Name:  "person-id",
			Usage: "The unique identifier of the person to delete",
		},
	},
	Action:          handlePeopleDelete,
	HideHelpCommand: true,
}

func handlePeopleCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonNewParams{}
	var res []byte
	_, err := cc.client.People.New(
		ctx,
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
	return ShowJSON("people create", json, format, transform)
}

func handlePeopleRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err := cc.client.People.Get(
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
	return ShowJSON("people retrieve", json, format, transform)
}

func handlePeopleUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonUpdateParams{}
	var res []byte
	_, err := cc.client.People.Update(
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
	return ShowJSON("people update", json, format, transform)
}

func handlePeopleList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonListParams{}
	var res []byte
	_, err := cc.client.People.List(
		ctx,
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
	return ShowJSON("people list", json, format, transform)
}

func handlePeopleDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err := cc.client.People.Delete(
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
	return ShowJSON("people delete", json, format, transform)
}
