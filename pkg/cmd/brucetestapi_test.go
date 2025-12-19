// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/DefinitelyATestOrg/test-api-cli/internal/mocktest"
)

func TestFormTest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"form-test",
		"--version", "2",
		"--user-id", "usr_abc123",
		"--date", "2019-12-27",
		"--datetime", "2019-12-27T18:11:19.117Z",
		"--time", "18:11:19.117Z",
		"--filter", "{meta: {level: 0}, status: status}",
		"--id-or-index", "0",
		"--limit", "1",
		"--tag", "string",
		"--blorp", "example value",
		"--pls-null", "null",
		"--preferences", "{alerts: true, theme: dark}",
		"--x-flag", "feature-a",
		"--x-flag", "feature-b",
		"--x-trace-id", "trace-xyz-789",
	)
}

func TestJsonTest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"json-test",
		"--version", "3",
		"--user-id", "usr_def456",
		"--date", "2019-12-27",
		"--datetime", "2019-12-27T18:11:19.117Z",
		"--time", "18:11:19.117Z",
		"--filter", "{meta: {level: 0}, status: status}",
		"--id-or-index", "0",
		"--limit", "1",
		"--tag", "string",
		"--blorp", "test data",
		"--pls-null", "null",
		"--preferences", "{alerts: false, theme: light}",
		"--x-flag", "experimental",
		"--x-flag", "verbose",
		"--x-trace-id", "trace-abc-123",
	)
}

func TestUpdateCount(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"update-count",
		"--body", "42",
	)
}
