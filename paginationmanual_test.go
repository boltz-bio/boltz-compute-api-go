// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomboltzbioboltzcomputeapigo_test

import (
	"context"
	"os"
	"testing"

	"github.com/boltz-bio/boltz-compute-api-go"
	"github.com/boltz-bio/boltz-compute-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-compute-api-go/option"
)

func TestManualPagination(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	page, err := client.Predictions.StructureAndBinding.List(context.TODO(), githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, structureAndBinding := range page.Data {
		t.Logf("%+v\n", structureAndBinding.ID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, structureAndBinding := range page.Data {
			t.Logf("%+v\n", structureAndBinding.ID)
		}
	}
}
