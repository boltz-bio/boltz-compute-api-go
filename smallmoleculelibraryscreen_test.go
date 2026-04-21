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

func TestSmallMoleculeLibraryScreenGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Get(
		context.TODO(),
		"screen_id",
		githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenGetParams{
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

func TestSmallMoleculeLibraryScreenListWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.List(context.TODO(), githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenListParams{
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

func TestSmallMoleculeLibraryScreenDeleteData(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.DeleteData(context.TODO(), "screen_id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SmallMolecule.LibraryScreen.EstimateCost(context.TODO(), githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParams{
		Molecules: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsMolecule{{
			Smiles: "smiles",
			ID:     githubcomboltzbioboltzcomputeapigo.String("id"),
		}},
		Target: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTarget{
			Entities: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityModificationUnion{{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetEntityModificationCcdModification: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: githubcomboltzbioboltzcomputeapigo.Bool(true),
			}},
			Bonds: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetBond{{
				Atom1: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1Union{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom1LigandAtom: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2Union{
					OfSmallMoleculeLibraryScreenEstimateCostsTargetBondAtom2LigandAtom: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintUnion{{
				OfSmallMoleculeLibraryScreenEstimateCostsTargetConstraintPocketConstraint: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsTargetConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
			PocketResidues: map[string][]int64{
				"A": {42, 43, 44, 67, 68, 69},
			},
			ReferenceLigands: []string{"string"},
		},
		IdempotencyKey: githubcomboltzbioboltzcomputeapigo.String("idempotency_key"),
		MoleculeFilters: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeLibraryScreenEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter{
					MaxHba:               0,
					MaxHbd:               0,
					MaxLogp:              0,
					MaxMw:                0,
					AllowSingleViolation: githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
		},
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

func TestSmallMoleculeLibraryScreenListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.ListResults(
		context.TODO(),
		"screen_id",
		githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenListResultsParams{
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

func TestSmallMoleculeLibraryScreenStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Start(context.TODO(), githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParams{
		Molecules: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsMolecule{{
			Smiles: "smiles",
			ID:     githubcomboltzbioboltzcomputeapigo.String("id"),
		}},
		Target: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTarget{
			Entities: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetEntityModificationUnion{{
					OfSmallMoleculeLibraryScreenStartsTargetEntityModificationCcdModification: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: githubcomboltzbioboltzcomputeapigo.Bool(true),
			}},
			Bonds: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetBond{{
				Atom1: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetBondAtom1Union{
					OfSmallMoleculeLibraryScreenStartsTargetBondAtom1LigandAtom: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetBondAtom2Union{
					OfSmallMoleculeLibraryScreenStartsTargetBondAtom2LigandAtom: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetConstraintUnion{{
				OfSmallMoleculeLibraryScreenStartsTargetConstraintPocketConstraint: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsTargetConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
			PocketResidues: map[string][]int64{
				"A": {42, 43, 44, 67, 68, 69},
			},
			ReferenceLigands: []string{"string"},
		},
		IdempotencyKey: githubcomboltzbioboltzcomputeapigo.String("idempotency_key"),
		MoleculeFilters: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeLibraryScreenStartsMoleculeFiltersCustomFilterLipinskiFilter: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeLibraryScreenStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
					MaxHba:               0,
					MaxHbd:               0,
					MaxLogp:              0,
					MaxMw:                0,
					AllowSingleViolation: githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
		},
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

func TestSmallMoleculeLibraryScreenStop(t *testing.T) {
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
	_, err := client.SmallMolecule.LibraryScreen.Stop(context.TODO(), "screen_id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
