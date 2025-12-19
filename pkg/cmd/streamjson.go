// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/DefinitelyATestOrg/test-api-cli/internal/apiquery"
	"github.com/DefinitelyATestOrg/test-api-go"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var streamJsonStream = cli.Command{
	Name:            "stream",
	Usage:           "Streams JSON objects as a chunked response (NDJSON)",
	Flags:           []cli.Flag{},
	Action:          handleStreamJsonStream,
	HideHelpCommand: true,
}

func handleStreamJsonStream(ctx context.Context, cmd *cli.Command) error {
	client := brucetestapi.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatDots,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	stream := client.StreamJson.StreamStreaming(ctx, options...)
	for stream.Next() {
		response := stream.Current()
		jsonData, err := json.Marshal(response)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(jsonData)
		ShowJSON(os.Stdout, "stream-json stream", obj, format, transform)
	}
	return stream.Err()
}
