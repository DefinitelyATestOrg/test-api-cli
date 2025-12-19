// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/DefinitelyATestOrg/test-api-cli/internal/mocktest"
)

func TestLinkedListCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"linked-list", "create",
		"--value", "a",
		"--next", "{value: b}",
	)
}
