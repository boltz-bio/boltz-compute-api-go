// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcomputeapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/boltz-bio/boltz-compute-api-go"
	"github.com/boltz-bio/boltz-compute-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-compute-api-go/option"
)

func TestAutoPagination(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	iter := client.Predictions.StructureAndBinding.ListAutoPaging(context.TODO(), boltzcomputeapi.PredictionStructureAndBindingListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		structureAndBinding := iter.Current()
		t.Logf("%+v\n", structureAndBinding.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
