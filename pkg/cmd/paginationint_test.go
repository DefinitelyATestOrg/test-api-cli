// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/DefinitelyATestOrg/test-api-cli/internal/mocktest"
)

func TestPaginationIntsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"pagination:ints", "list",
		"--page", "1",
		"--size", "1",
	)
}
