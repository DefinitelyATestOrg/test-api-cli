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

var peopleCreate = cli.Command{
	Name:  "create",
	Usage: "Create a new person and add them to the system.",
	Flags: []cli.Flag{
		&cli.GenericFlag{
			Name:      "image-base64",
			Usage:     "Image of the person (base64)",
			Value:     &fileReader{Base64Encoded: true},
			TakesFile: true,
		},
		&cli.GenericFlag{
			Name:      "image-binary",
			Usage:     "Image of the person (binary)",
			Value:     &fileReader{},
			TakesFile: true,
		},
		&cli.StringFlag{
			Name:  "job",
			Usage: "The person's job",
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
		&cli.StringFlag{
			Name:  "job",
			Usage: "The updated job of the person",
		},
	},
	Action:          handlePeopleUpdate,
	HideHelpCommand: true,
}

var peopleList = cli.Command{
	Name:  "list",
	Usage: "Get a list of all people.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "job",
			Usage: "Job name to search for",
		},
		&cli.StringFlag{
			Name:  "name",
			Usage: "Full name to search for",
		},
		&cli.GenericFlag{
			Name:      "nickname",
			Usage:     "Nickname to search for",
			Value:     &fileReader{Base64Encoded: true},
			TakesFile: true,
		},
		&cli.Int64Flag{
			Name:  "page",
			Usage: "Page number",
			Value: 1,
		},
		&cli.Int64Flag{
			Name:  "size",
			Usage: "Page size",
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
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonNewParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"image-base64": "image_base64",
		"image-binary": "image_binary",
		"job":          "job",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.People.New(
		ctx,
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
	return ShowJSON("people create", json, format, transform)
}

func handlePeopleRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err := client.People.Get(
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
	return ShowJSON("people retrieve", json, format, transform)
}

func handlePeopleUpdate(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("person-id") && len(unusedArgs) > 0 {
		cmd.Set("person-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonUpdateParams{}
	if err := unmarshalStdinWithFlags(cmd, map[string]string{
		"job": "job",
	}, &params); err != nil {
		return err
	}
	var res []byte
	_, err := client.People.Update(
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
	return ShowJSON("people update", json, format, transform)
}

func handlePeopleList(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := brucetestapi.PersonListParams{}
	if cmd.IsSet("job") {
		params.Job = brucetestapi.Opt(cmd.Value("job").(string))
	}
	if cmd.IsSet("name") {
		params.Name = brucetestapi.Opt(cmd.Value("name").(string))
	}
	if cmd.IsSet("nickname") {
		params.Nickname = brucetestapi.Opt(cmd.Value("nickname").(string))
	}
	if cmd.IsSet("page") {
		params.Page = brucetestapi.Opt(cmd.Value("page").(int64))
	}
	if cmd.IsSet("size") {
		params.Size = brucetestapi.Opt(cmd.Value("size").(int64))
	}
	var res []byte
	_, err := client.People.List(
		ctx,
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
	return ShowJSON("people list", json, format, transform)
}

func handlePeopleDelete(ctx context.Context, cmd *cli.Command) error {
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
	_, err := client.People.Delete(
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
	return ShowJSON("people delete", json, format, transform)
}
