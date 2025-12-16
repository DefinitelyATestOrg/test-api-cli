// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/DefinitelyATestOrg/test-api-cli/internal/apiquery"
	"github.com/DefinitelyATestOrg/test-api-go"
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
	)
	if err != nil {
		return err
	}

	stream := client.StreamJson.StreamStreaming(ctx, options...)
	for stream.Next() {
		fmt.Printf("%s\n", stream.Current().RawJSON())
	}
	return stream.Err()
}
