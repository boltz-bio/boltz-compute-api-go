// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcomputeapi_test

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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.Get(
		context.TODO(),
		"screen_id",
		boltzcomputeapi.ProteinLibraryScreenGetParams{
			WorkspaceID: boltzcomputeapi.String("workspace_id"),
		},
	)
	if err != nil {
		var apierr *boltzcomputeapi.Error
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.List(context.TODO(), boltzcomputeapi.ProteinLibraryScreenListParams{
		AfterID:     boltzcomputeapi.String("after_id"),
		BeforeID:    boltzcomputeapi.String("before_id"),
		Limit:       boltzcomputeapi.Int(1),
		WorkspaceID: boltzcomputeapi.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzcomputeapi.Error
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.DeleteData(context.TODO(), "screen_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.EstimateCost(context.TODO(), boltzcomputeapi.ProteinLibraryScreenEstimateCostParams{
		Proteins: []boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsProtein{{
			Entities: []boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsProteinEntityUnion{{
				OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntity: &boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationUnion{{
						OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntityModificationCcdModification: &boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value:  "value",
					Cyclic: boltzcomputeapi.Bool(true),
				},
			}},
			ID: boltzcomputeapi.String("id"),
		}},
		Target: boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsTargetUnion{
			OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTarget: &boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetStructureURLSource: &boltzcomputeapi.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
		IdempotencyKey: boltzcomputeapi.String("idempotency_key"),
		WorkspaceID:    boltzcomputeapi.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzcomputeapi.Error
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.ListResults(
		context.TODO(),
		"screen_id",
		boltzcomputeapi.ProteinLibraryScreenListResultsParams{
			AfterID:     boltzcomputeapi.String("after_id"),
			BeforeID:    boltzcomputeapi.String("before_id"),
			Limit:       boltzcomputeapi.Int(1),
			WorkspaceID: boltzcomputeapi.String("workspace_id"),
		},
	)
	if err != nil {
		var apierr *boltzcomputeapi.Error
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.Start(context.TODO(), boltzcomputeapi.ProteinLibraryScreenStartParams{
		Proteins: []boltzcomputeapi.ProteinLibraryScreenStartParamsProtein{{
			Entities: []boltzcomputeapi.ProteinLibraryScreenStartParamsProteinEntityUnion{{
				OfProteinLibraryScreenStartsProteinEntityProteinEntity: &boltzcomputeapi.ProteinLibraryScreenStartParamsProteinEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []boltzcomputeapi.ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationUnion{{
						OfProteinLibraryScreenStartsProteinEntityProteinEntityModificationCcdModification: &boltzcomputeapi.ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value:  "value",
					Cyclic: boltzcomputeapi.Bool(true),
				},
			}},
			ID: boltzcomputeapi.String("id"),
		}},
		Target: boltzcomputeapi.ProteinLibraryScreenStartParamsTargetUnion{
			OfProteinLibraryScreenStartsTargetStructureTemplateTarget: &boltzcomputeapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzcomputeapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinLibraryScreenStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzcomputeapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzcomputeapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: boltzcomputeapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinLibraryScreenStartsTargetStructureTemplateTargetStructureURLSource: &boltzcomputeapi.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
		IdempotencyKey: boltzcomputeapi.String("idempotency_key"),
		WorkspaceID:    boltzcomputeapi.String("workspace_id"),
	})
	if err != nil {
		var apierr *boltzcomputeapi.Error
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.Stop(context.TODO(), "screen_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
