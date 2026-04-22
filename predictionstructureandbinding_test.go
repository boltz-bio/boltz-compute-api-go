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

func TestPredictionStructureAndBindingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.Get(
		context.TODO(),
		"id",
		boltzcompute.PredictionStructureAndBindingGetParams{
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

func TestPredictionStructureAndBindingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.List(context.TODO(), boltzcompute.PredictionStructureAndBindingListParams{
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

func TestPredictionStructureAndBindingDeleteData(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPredictionStructureAndBindingEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.EstimateCost(context.TODO(), boltzcompute.PredictionStructureAndBindingEstimateCostParams{
		Input: boltzcompute.PredictionStructureAndBindingEstimateCostParamsInput{
			Entities: []boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputEntityUnion{{
				OfPredictionStructureAndBindingEstimateCostsInputEntityProteinEntity: &boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzcompute.Bool(true),
					Modifications: []boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModificationUnion{{
						OfPredictionStructureAndBindingEstimateCostsInputEntityProteinEntityModificationCcdModification: &boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
				},
			}},
			Binding: boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputBindingUnion{
				OfPredictionStructureAndBindingEstimateCostsInputBindingLigandProteinBinding: &boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding{
					BinderChainID: "binder_chain_id",
				},
			},
			Bonds: []boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputBond{{
				Atom1: boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputBondAtom1Union{
					OfPredictionStructureAndBindingEstimateCostsInputBondAtom1LigandAtom: &boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputBondAtom2Union{
					OfPredictionStructureAndBindingEstimateCostsInputBondAtom2LigandAtom: &boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputConstraintUnion{{
				OfPredictionStructureAndBindingEstimateCostsInputConstraintPocketConstraint: &boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzcompute.Bool(true),
				},
			}},
			ModelOptions: boltzcompute.PredictionStructureAndBindingEstimateCostParamsInputModelOptions{
				RecyclingSteps: boltzcompute.Int(1),
				SamplingSteps:  boltzcompute.Int(1),
				StepScale:      boltzcompute.Float(1.3),
			},
			NumSamples: boltzcompute.Int(1),
		},
		Model:          "boltz-2.1",
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

func TestPredictionStructureAndBindingStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.Start(context.TODO(), boltzcompute.PredictionStructureAndBindingStartParams{
		Input: boltzcompute.PredictionStructureAndBindingStartParamsInput{
			Entities: []boltzcompute.PredictionStructureAndBindingStartParamsInputEntityUnion{{
				OfPredictionStructureAndBindingStartsInputEntityProteinEntity: &boltzcompute.PredictionStructureAndBindingStartParamsInputEntityProteinEntity{
					ChainIDs: []string{"string"},
					Value:    "value",
					Cyclic:   boltzcompute.Bool(true),
					Modifications: []boltzcompute.PredictionStructureAndBindingStartParamsInputEntityProteinEntityModificationUnion{{
						OfPredictionStructureAndBindingStartsInputEntityProteinEntityModificationCcdModification: &boltzcompute.PredictionStructureAndBindingStartParamsInputEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
				},
			}},
			Binding: boltzcompute.PredictionStructureAndBindingStartParamsInputBindingUnion{
				OfPredictionStructureAndBindingStartsInputBindingLigandProteinBinding: &boltzcompute.PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding{
					BinderChainID: "binder_chain_id",
				},
			},
			Bonds: []boltzcompute.PredictionStructureAndBindingStartParamsInputBond{{
				Atom1: boltzcompute.PredictionStructureAndBindingStartParamsInputBondAtom1Union{
					OfPredictionStructureAndBindingStartsInputBondAtom1LigandAtom: &boltzcompute.PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzcompute.PredictionStructureAndBindingStartParamsInputBondAtom2Union{
					OfPredictionStructureAndBindingStartsInputBondAtom2LigandAtom: &boltzcompute.PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzcompute.PredictionStructureAndBindingStartParamsInputConstraintUnion{{
				OfPredictionStructureAndBindingStartsInputConstraintPocketConstraint: &boltzcompute.PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzcompute.Bool(true),
				},
			}},
			ModelOptions: boltzcompute.PredictionStructureAndBindingStartParamsInputModelOptions{
				RecyclingSteps: boltzcompute.Int(1),
				SamplingSteps:  boltzcompute.Int(1),
				StepScale:      boltzcompute.Float(1.3),
			},
			NumSamples: boltzcompute.Int(1),
		},
		Model:          "boltz-2.1",
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
