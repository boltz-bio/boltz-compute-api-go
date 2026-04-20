// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomboltzbioboltzcomputeapigo_test

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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.Get(
		context.TODO(),
		"screen_id",
		githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenGetParams{
			WorkspaceID: githubcomboltzbioboltzcomputeapigo.String("workspace_id"),
		},
	)
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.List(context.TODO(), githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenListParams{
		AfterID:     githubcomboltzbioboltzcomputeapigo.String("after_id"),
		BeforeID:    githubcomboltzbioboltzcomputeapigo.String("before_id"),
		Limit:       githubcomboltzbioboltzcomputeapigo.Int(1),
		WorkspaceID: githubcomboltzbioboltzcomputeapigo.String("workspace_id"),
	})
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.DeleteData(context.TODO(), "screen_id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.EstimateCost(context.TODO(), githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParams{
		Proteins: []githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsProtein{{
			Entities: []githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsProteinEntityUnion{{
				OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntity: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationUnion{{
						OfProteinLibraryScreenEstimateCostsProteinEntityProteinEntityModificationCcdModification: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsProteinEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value:  "value",
					Cyclic: githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
			ID: githubcomboltzbioboltzcomputeapigo.String("id"),
		}},
		Target: githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsTargetUnion{
			OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTarget: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinLibraryScreenEstimateCostsTargetStructureTemplateTargetStructureURLSource: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenEstimateCostParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
		IdempotencyKey: githubcomboltzbioboltzcomputeapigo.String("idempotency_key"),
		WorkspaceID:    githubcomboltzbioboltzcomputeapigo.String("workspace_id"),
	})
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.ListResults(
		context.TODO(),
		"screen_id",
		githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenListResultsParams{
			AfterID:     githubcomboltzbioboltzcomputeapigo.String("after_id"),
			BeforeID:    githubcomboltzbioboltzcomputeapigo.String("before_id"),
			Limit:       githubcomboltzbioboltzcomputeapigo.Int(1),
			WorkspaceID: githubcomboltzbioboltzcomputeapigo.String("workspace_id"),
		},
	)
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.Start(context.TODO(), githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParams{
		Proteins: []githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsProtein{{
			Entities: []githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsProteinEntityUnion{{
				OfProteinLibraryScreenStartsProteinEntityProteinEntity: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsProteinEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationUnion{{
						OfProteinLibraryScreenStartsProteinEntityProteinEntityModificationCcdModification: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsProteinEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value:  "value",
					Cyclic: githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
			ID: githubcomboltzbioboltzcomputeapigo.String("id"),
		}},
		Target: githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsTargetUnion{
			OfProteinLibraryScreenStartsTargetStructureTemplateTarget: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinLibraryScreenStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinLibraryScreenStartsTargetStructureTemplateTargetStructureURLSource: &githubcomboltzbioboltzcomputeapigo.ProteinLibraryScreenStartParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
		IdempotencyKey: githubcomboltzbioboltzcomputeapigo.String("idempotency_key"),
		WorkspaceID:    githubcomboltzbioboltzcomputeapigo.String("workspace_id"),
	})
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.LibraryScreen.Stop(context.TODO(), "screen_id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
