// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/DefinitelyATestOrg/test-api-cli/internal/mocktest"
)

func TestStreamJsonStream(t *testing.T) {
	t.Skip("Prism doesn't support application/x-ndjson responses")
	mocktest.TestRunMockTestWithFlags(
		t,
		"stream-json", "stream",
	)
}
