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

func TestSmallMoleculeDesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Get(
		context.TODO(),
		"run_id",
		githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignGetParams{
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

func TestSmallMoleculeDesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.List(context.TODO(), githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignListParams{
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

func TestSmallMoleculeDesignDeleteData(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.DeleteData(context.TODO(), "run_id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.SmallMolecule.Design.EstimateCost(context.TODO(), githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParams{
		NumMolecules: 10,
		Target: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTarget{
			Entities: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetEntityModificationUnion{{
					OfSmallMoleculeDesignEstimateCostsTargetEntityModificationCcdModification: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: githubcomboltzbioboltzcomputeapigo.Bool(true),
			}},
			Bonds: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetBond{{
				Atom1: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetBondAtom1Union{
					OfSmallMoleculeDesignEstimateCostsTargetBondAtom1LigandAtom: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetBondAtom2Union{
					OfSmallMoleculeDesignEstimateCostsTargetBondAtom2LigandAtom: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetConstraintUnion{{
				OfSmallMoleculeDesignEstimateCostsTargetConstraintPocketConstraint: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsTargetConstraintPocketConstraint{
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
		ChemicalSpace:  githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsChemicalSpaceEnamineReal,
		IdempotencyKey: githubcomboltzbioboltzcomputeapigo.String("idempotency_key"),
		MoleculeFilters: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeDesignEstimateCostsMoleculeFiltersCustomFilterLipinskiFilter: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignEstimateCostParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeDesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.ListResults(
		context.TODO(),
		"run_id",
		githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignListResultsParams{
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

func TestSmallMoleculeDesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Start(context.TODO(), githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParams{
		NumMolecules: 10,
		Target: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTarget{
			Entities: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetEntity{{
				ChainIDs: []string{"string"},
				Modifications: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetEntityModificationUnion{{
					OfSmallMoleculeDesignStartsTargetEntityModificationCcdModification: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetEntityModificationCcdModification{
						ResidueIndex: 0,
						Value:        "value",
					},
				}},
				Value:  "value",
				Cyclic: githubcomboltzbioboltzcomputeapigo.Bool(true),
			}},
			Bonds: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetBond{{
				Atom1: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetBondAtom1Union{
					OfSmallMoleculeDesignStartsTargetBondAtom1LigandAtom: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetBondAtom2Union{
					OfSmallMoleculeDesignStartsTargetBondAtom2LigandAtom: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetConstraintUnion{{
				OfSmallMoleculeDesignStartsTargetConstraintPocketConstraint: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsTargetConstraintPocketConstraint{
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
		ChemicalSpace:  githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsChemicalSpaceEnamineReal,
		IdempotencyKey: githubcomboltzbioboltzcomputeapigo.String("idempotency_key"),
		MoleculeFilters: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsMoleculeFilters{
			BoltzSmartsCatalogFilterLevel: githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsMoleculeFiltersBoltzSmartsCatalogFilterLevelRecommended,
			CustomFilters: []githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterUnion{{
				OfSmallMoleculeDesignStartsMoleculeFiltersCustomFilterLipinskiFilter: &githubcomboltzbioboltzcomputeapigo.SmallMoleculeDesignStartParamsMoleculeFiltersCustomFilterLipinskiFilter{
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

func TestSmallMoleculeDesignStop(t *testing.T) {
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
	_, err := client.SmallMolecule.Design.Stop(context.TODO(), "run_id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
