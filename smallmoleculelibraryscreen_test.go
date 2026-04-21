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

func TestSmallMoleculeLibraryScreenGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Get(
		context.TODO(),
		"screen_id",
		boltzcompute.SmallMoleculeLibraryScreenGetParams{
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

func TestSmallMoleculeLibraryScreenListWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.List(context.TODO(), boltzcompute.SmallMoleculeLibraryScreenListParams{
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

func TestSmallMoleculeLibraryScreenDeleteData(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.DeleteData(context.TODO(), "screen_id")
	if err != nil {
		var apierr *boltzcompute.Error
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
	client := boltzcompute.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SmallMolecule.LibraryScreen.EstimateCost(context.TODO(), boltzcompute.SmallMoleculeLibraryScreenEstimateCostParams{
		Molecules: []boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsMolecule{{
			Smiles: "smiles",
			ID:     boltzcompute.String("id"),
		}},
		Target: boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTarget{
			Entities: []boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityModificationUnion{{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetEntityModificationCcdModification: &boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: boltzcompute.Bool(true),
			}},
			Bonds: []boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetBond{{
				Atom1: boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1Union{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1LigandAtom: &boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2Union{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2LigandAtom: &boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintUnion{{
				OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintPocketConstraint: &boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintPocketConstraint{
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
		IdempotencyKey: boltzcompute.String("idempotency_key"),
		MoleculeFilters: boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter: &boltzcompute.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeLibraryScreenListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.ListResults(
		context.TODO(),
		"screen_id",
		boltzcompute.SmallMoleculeLibraryScreenListResultsParams{
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

func TestSmallMoleculeLibraryScreenStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Start(context.TODO(), boltzcompute.SmallMoleculeLibraryScreenStartParams{
		Molecules: []boltzcompute.SmallMoleculeLibraryScreenStartParamsMolecule{{
			Smiles: "smiles",
			ID:     boltzcompute.String("id"),
		}},
		Target: boltzcompute.SmallMoleculeLibraryScreenStartParamsTarget{
			Entities: []boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetEntityModificationUnion{{
					OfSmallMoleculeLibraryScreenStartsTargetEntityModificationCcdModification: &boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: boltzcompute.Bool(true),
			}},
			Bonds: []boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetBond{{
				Atom1: boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetBondAtom1Union{
					OfSmallMoleculeLibraryScreenStartsTargetBondAtom1LigandAtom: &boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetBondAtom2Union{
					OfSmallMoleculeLibraryScreenStartsTargetBondAtom2LigandAtom: &boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetConstraintUnion{{
				OfSmallMoleculeLibraryScreenStartsTargetConstraintPocketConstraint: &boltzcompute.SmallMoleculeLibraryScreenStartParamsTargetConstraintPocketConstraint{
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
		IdempotencyKey: boltzcompute.String("idempotency_key"),
		MoleculeFilters: boltzcompute.SmallMoleculeLibraryScreenStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: boltzcompute.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []boltzcompute.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterLipinskiFilter: &boltzcompute.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeLibraryScreenStop(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Stop(context.TODO(), "screen_id")
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
