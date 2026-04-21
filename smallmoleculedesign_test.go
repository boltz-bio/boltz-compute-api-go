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

func TestSmallMoleculeDesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Get(
		context.TODO(),
		"id",
		boltzcompute.SmallMoleculeDesignGetParams{
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

func TestSmallMoleculeDesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.List(context.TODO(), boltzcompute.SmallMoleculeDesignListParams{
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

func TestSmallMoleculeDesignDeleteData(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzcompute.Error
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
	client := boltzcompute.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SmallMolecule.Design.EstimateCost(context.TODO(), boltzcompute.SmallMoleculeDesignEstimateCostParams{
		NumMolecules: 10,
		Target: boltzcompute.SmallMoleculeDesignEstimateCostParamsTarget{
			Entities: []boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetEntityModificationUnion{{
					OfSmallMoleculeDesignEstimateCostsTargetEntityModificationCcdModification: &boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: boltzcompute.Bool(true),
			}},
			Bonds: []boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetBond{{
				Atom1: boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetBondAtom1Union{
					OfSmallMoleculeDesignEstimateCostsTargetBondAtom1LigandAtom: &boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetBondAtom2Union{
					OfSmallMoleculeDesignEstimateCostsTargetBondAtom2LigandAtom: &boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetConstraintUnion{{
				OfSmallMoleculeDesignEstimateCostsTargetConstraintPocketConstraint: &boltzcompute.SmallMoleculeDesignEstimateCostParamsTargetConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzcompute.Bool(true),
				},
			}},
			PocketResidues: map[string][]int64{
				"A": {42, 43, 44, 67, 68, 69},
			},
			ReferenceLigands: []string{"string"},
		},
		ChemicalSpace:  boltzcompute.SmallMoleculeDesignEstimateCostParamsChemicalSpaceEnamineReal,
		IdempotencyKey: boltzcompute.String("idempotency_key"),
		MoleculeFilters: boltzcompute.SmallMoleculeDesignEstimateCostParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzcompute.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzcompute.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter: &boltzcompute.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter{
					MaxHba:               0,
					MaxHbd:               0,
					MaxLogp:              0,
					MaxMw:                0,
					AllowSingleViolation: boltzcompute.Bool(true),
				},
			}},
		},
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

func TestSmallMoleculeDesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.ListResults(
		context.TODO(),
		"id",
		boltzcompute.SmallMoleculeDesignListResultsParams{
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

func TestSmallMoleculeDesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Start(context.TODO(), boltzcompute.SmallMoleculeDesignStartParams{
		NumMolecules: 10,
		Target: boltzcompute.SmallMoleculeDesignStartParamsTarget{
			Entities: []boltzcompute.SmallMoleculeDesignStartParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []boltzcompute.SmallMoleculeDesignStartParamsTargetEntityModificationUnion{{
					OfSmallMoleculeDesignStartsTargetEntityModificationCcdModification: &boltzcompute.SmallMoleculeDesignStartParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: boltzcompute.Bool(true),
			}},
			Bonds: []boltzcompute.SmallMoleculeDesignStartParamsTargetBond{{
				Atom1: boltzcompute.SmallMoleculeDesignStartParamsTargetBondAtom1Union{
					OfSmallMoleculeDesignStartsTargetBondAtom1LigandAtom: &boltzcompute.SmallMoleculeDesignStartParamsTargetBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzcompute.SmallMoleculeDesignStartParamsTargetBondAtom2Union{
					OfSmallMoleculeDesignStartsTargetBondAtom2LigandAtom: &boltzcompute.SmallMoleculeDesignStartParamsTargetBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzcompute.SmallMoleculeDesignStartParamsTargetConstraintUnion{{
				OfSmallMoleculeDesignStartsTargetConstraintPocketConstraint: &boltzcompute.SmallMoleculeDesignStartParamsTargetConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzcompute.Bool(true),
				},
			}},
			PocketResidues: map[string][]int64{
				"A": {42, 43, 44, 67, 68, 69},
			},
			ReferenceLigands: []string{"string"},
		},
		ChemicalSpace:  boltzcompute.SmallMoleculeDesignStartParamsChemicalSpaceEnamineReal,
		IdempotencyKey: boltzcompute.String("idempotency_key"),
		MoleculeFilters: boltzcompute.SmallMoleculeDesignStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzcompute.SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzcompute.SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterLipinskiFilter: &boltzcompute.SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
					MaxHba:               0,
					MaxHbd:               0,
					MaxLogp:              0,
					MaxMw:                0,
					AllowSingleViolation: boltzcompute.Bool(true),
				},
			}},
		},
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

func TestSmallMoleculeDesignStop(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Stop(context.TODO(), "id")
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
