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

func TestSmallMoleculeDesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Get(
		context.TODO(),
		"run_id",
		boltzcomputeapi.SmallMoleculeDesignGetParams{
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

func TestSmallMoleculeDesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.List(context.TODO(), boltzcomputeapi.SmallMoleculeDesignListParams{
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

func TestSmallMoleculeDesignDeleteData(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.DeleteData(context.TODO(), "run_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSmallMoleculeDesignEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.EstimateCost(context.TODO(), boltzcomputeapi.SmallMoleculeDesignEstimateCostParams{
		NumMolecules: 10,
		Target: boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsTarget{
			Entities: []boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsTargetEntityModificationUnion{{
					OfSmallMoleculeDesignEstimateCostsTargetEntityModificationCcdModification: &boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsTargetEntityModificationCcdModification{
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
		ChemicalSpace:  boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsChemicalSpaceEnamineReal,
		IdempotencyKey: boltzcomputeapi.String("idempotency_key"),
		MoleculeFilters: boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter: &boltzcomputeapi.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeDesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.ListResults(
		context.TODO(),
		"run_id",
		boltzcomputeapi.SmallMoleculeDesignListResultsParams{
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

func TestSmallMoleculeDesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Start(context.TODO(), boltzcomputeapi.SmallMoleculeDesignStartParams{
		NumMolecules: 10,
		Target: boltzcomputeapi.SmallMoleculeDesignStartParamsTarget{
			Entities: []boltzcomputeapi.SmallMoleculeDesignStartParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []boltzcomputeapi.SmallMoleculeDesignStartParamsTargetEntityModificationUnion{{
					OfSmallMoleculeDesignStartsTargetEntityModificationCcdModification: &boltzcomputeapi.SmallMoleculeDesignStartParamsTargetEntityModificationCcdModification{
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
		ChemicalSpace:  boltzcomputeapi.SmallMoleculeDesignStartParamsChemicalSpaceEnamineReal,
		IdempotencyKey: boltzcomputeapi.String("idempotency_key"),
		MoleculeFilters: boltzcomputeapi.SmallMoleculeDesignStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzcomputeapi.SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzcomputeapi.SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterLipinskiFilter: &boltzcomputeapi.SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeDesignStop(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Stop(context.TODO(), "run_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
