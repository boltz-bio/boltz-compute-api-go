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

func TestSmallMoleculeLibraryScreenGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Get(
		context.TODO(),
		"screen_id",
		boltzcomputeapi.SmallMoleculeLibraryScreenGetParams{
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

func TestSmallMoleculeLibraryScreenListWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.List(context.TODO(), boltzcomputeapi.SmallMoleculeLibraryScreenListParams{
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

func TestSmallMoleculeLibraryScreenDeleteData(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.DeleteData(context.TODO(), "screen_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmallMoleculeLibraryScreenEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.EstimateCost(context.TODO(), boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParams{
		Molecules: []boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsMolecule{{
			Smiles: "smiles",
			ID:     boltzcomputeapi.String("id"),
		}},
		Target: boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsTarget{
			Entities: []boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityModificationUnion{{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetEntityModificationCcdModification: &boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: boltzcomputeapi.Bool(true),
			}},
			PocketResidues: map[string][]int64{
				"A": {42, 43, 44, 67, 68, 69},
			},
			ReferenceLigands: []string{"string"},
		},
		IdempotencyKey: boltzcomputeapi.String("idempotency_key"),
		MoleculeFilters: boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter: &boltzcomputeapi.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter{
					MaxHba:               0,
					MaxHbd:               0,
					MaxLogp:              0,
					MaxMw:                0,
					AllowSingleViolation: boltzcomputeapi.Bool(true),
				},
			}},
		},
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

func TestSmallMoleculeLibraryScreenListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.ListResults(
		context.TODO(),
		"screen_id",
		boltzcomputeapi.SmallMoleculeLibraryScreenListResultsParams{
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

func TestSmallMoleculeLibraryScreenStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Start(context.TODO(), boltzcomputeapi.SmallMoleculeLibraryScreenStartParams{
		Molecules: []boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsMolecule{{
			Smiles: "smiles",
			ID:     boltzcomputeapi.String("id"),
		}},
		Target: boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsTarget{
			Entities: []boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsTargetEntityModificationUnion{{
					OfSmallMoleculeLibraryScreenStartsTargetEntityModificationCcdModification: &boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: boltzcomputeapi.Bool(true),
			}},
			PocketResidues: map[string][]int64{
				"A": {42, 43, 44, 67, 68, 69},
			},
			ReferenceLigands: []string{"string"},
		},
		IdempotencyKey: boltzcomputeapi.String("idempotency_key"),
		MoleculeFilters: boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterLipinskiFilter: &boltzcomputeapi.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
					MaxHba:               0,
					MaxHbd:               0,
					MaxLogp:              0,
					MaxMw:                0,
					AllowSingleViolation: boltzcomputeapi.Bool(true),
				},
			}},
		},
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

func TestSmallMoleculeLibraryScreenStop(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Stop(context.TODO(), "screen_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
