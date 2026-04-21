// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcompute_test

import (
	"context"
	"os"
	"testing"

	"github.com/boltz-bio/boltz-compute-api-go"
	"github.com/boltz-bio/boltz-compute-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-compute-api-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := boltzcompute.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	response, err := client.Predictions.StructureAndBinding.Start(context.TODO(), boltzcompute.PredictionStructureAndBindingStartParams{
		Input: boltzcompute.PredictionStructureAndBindingStartParamsInput{
			Entities: []boltzcompute.PredictionStructureAndBindingStartParamsInputEntityUnion{{
				OfPredictionStructureAndBindingStartsInputEntityProteinEntity: &boltzcompute.PredictionStructureAndBindingStartParamsInputEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []boltzcompute.PredictionStructureAndBindingStartParamsInputEntityProteinEntityModificationUnion{{
						OfPredictionStructureAndBindingStartsInputEntityProteinEntityModificationCcdModification: &boltzcompute.PredictionStructureAndBindingStartParamsInputEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value: "value",
				},
			}},
		},
		Model: "boltz-2.1",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.ID)
}
