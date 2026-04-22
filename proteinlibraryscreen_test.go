// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcompute_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/boltz-bio/boltz-compute-api-go"
	"github.com/boltz-bio/boltz-compute-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-compute-api-go/option"
)

func TestProteinLibraryScreenGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.Get(
		context.TODO(),
		"id",
		boltzcompute.ProteinLibraryScreenGetParams{
			WorkspaceID: boltzcompute.String("workspace_id"),
		},
	)
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenListWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.List(context.TODO(), boltzcompute.ProteinLibraryScreenListParams{
		AfterID:     boltzcompute.String("after_id"),
		BeforeID:    boltzcompute.String("before_id"),
		Limit:       boltzcompute.Int(1),
		WorkspaceID: boltzcompute.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenDeleteData(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.EstimateCost(context.TODO(), boltzcompute.ProteinLibraryScreenEstimateCostParams{
		Proteins: []boltzcompute.ProteinLibraryScreenEstimateCostParamsProtein{{
			Entities: []boltzcompute.ProteinLibraryScreenEstimateCostParamsProteinEntityUnion{{
				OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntity: &boltzcompute.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzcompute.Bool(true),
					Modifications: []boltzcompute.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationUnion{{
						OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntityModificationCcdModification: &boltzcompute.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
				},
			}},
			ID: boltzcompute.String("id"),
		}},
		Target: boltzcompute.ProteinLibraryScreenEstimateCostParamsTargetUnion{
			OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTarget: &boltzcompute.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzcompute.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzcompute.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzcompute.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: boltzcompute.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetStructureURLSource: &boltzcompute.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
		IdempotencyKey: boltzcompute.String("idempotency_key"),
		WorkspaceID:    boltzcompute.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.ListResults(
		context.TODO(),
		"id",
		boltzcompute.ProteinLibraryScreenListResultsParams{
			AfterID:     boltzcompute.String("after_id"),
			BeforeID:    boltzcompute.String("before_id"),
			Limit:       boltzcompute.Int(1),
			WorkspaceID: boltzcompute.String("workspace_id"),
		},
	)
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.Start(context.TODO(), boltzcompute.ProteinLibraryScreenStartParams{
		Proteins: []boltzcompute.ProteinLibraryScreenStartParamsProtein{{
			Entities: []boltzcompute.ProteinLibraryScreenStartParamsProteinEntityUnion{{
				OfProteinLibraryScreenStartsProteinEntityProteinEntity: &boltzcompute.ProteinLibraryScreenStartParamsProteinEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzcompute.Bool(true),
					Modifications: []boltzcompute.ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationUnion{{
						OfProteinLibraryScreenStartsProteinEntityProteinEntityModificationCcdModification: &boltzcompute.ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
				},
			}},
			ID: boltzcompute.String("id"),
		}},
		Target: boltzcompute.ProteinLibraryScreenStartParamsTargetUnion{
			OfProteinLibraryScreenStartsTargetStructureTemplateTarget: &boltzcompute.ProteinLibraryScreenStartParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzcompute.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinLibraryScreenStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzcompute.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzcompute.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: boltzcompute.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinLibraryScreenStartsTargetStructureTemplateTargetStructureURLSource: &boltzcompute.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
		IdempotencyKey: boltzcompute.String("idempotency_key"),
		WorkspaceID:    boltzcompute.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinLibraryScreenStop(t *testing.T) {
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
	_, err := client.Protein.LibraryScreen.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
