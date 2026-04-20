// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcomputeapi_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/boltz-compute-api-go"
	"github.com/stainless-sdks/boltz-compute-api-go/internal/testutil"
	"github.com/stainless-sdks/boltz-compute-api-go/option"
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Predictions.StructureAndBinding.Get(
		context.TODO(),
		"id",
		boltzcomputeapi.PredictionStructureAndBindingGetParams{
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

func TestPredictionStructureAndBindingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.List(context.TODO(), boltzcomputeapi.PredictionStructureAndBindingListParams{
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

func TestPredictionStructureAndBindingDeleteData(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Predictions.StructureAndBinding.EstimateCost(context.TODO(), boltzcomputeapi.PredictionStructureAndBindingEstimateCostParams{
		Input: boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInput{
			Entities: []boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputEntityUnion{{
				OfPredictionStructureAndBindingEstimateCostsInputEntityProteinEntity: &boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModificationUnion{{
						OfPredictionStructureAndBindingEstimateCostsInputEntityProteinEntityModificationCcdModification: &boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value:  "value",
					Cyclic: boltzcomputeapi.Bool(true),
				},
			}},
			Binding: boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputBindingUnion{
				OfPredictionStructureAndBindingEstimateCostsInputBindingLigandProteinBinding: &boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding{
					BinderChainID: "binder_chain_id",
				},
			},
			Bonds: []boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputBond{{
				Atom1: boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputBondAtom1Union{
					OfPredictionStructureAndBindingEstimateCostsInputBondAtom1LigandAtom: &boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputBondAtom2Union{
					OfPredictionStructureAndBindingEstimateCostsInputBondAtom2LigandAtom: &boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputConstraintUnion{{
				OfPredictionStructureAndBindingEstimateCostsInputConstraintPocketConstraint: &boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzcomputeapi.Bool(true),
				},
			}},
			ModelOptions: boltzcomputeapi.PredictionStructureAndBindingEstimateCostParamsInputModelOptions{
				RecyclingSteps: boltzcomputeapi.Int(1),
				SamplingSteps:  boltzcomputeapi.Int(1),
				StepScale:      boltzcomputeapi.Float(1.3),
			},
			NumSamples: boltzcomputeapi.Int(1),
		},
		Model:          "boltz-2.1",
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

func TestPredictionStructureAndBindingStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.Start(context.TODO(), boltzcomputeapi.PredictionStructureAndBindingStartParams{
		Input: boltzcomputeapi.PredictionStructureAndBindingStartParamsInput{
			Entities: []boltzcomputeapi.PredictionStructureAndBindingStartParamsInputEntityUnion{{
				OfPredictionStructureAndBindingStartsInputEntityProteinEntity: &boltzcomputeapi.PredictionStructureAndBindingStartParamsInputEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []boltzcomputeapi.PredictionStructureAndBindingStartParamsInputEntityProteinEntityModificationUnion{{
						OfPredictionStructureAndBindingStartsInputEntityProteinEntityModificationCcdModification: &boltzcomputeapi.PredictionStructureAndBindingStartParamsInputEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value:  "value",
					Cyclic: boltzcomputeapi.Bool(true),
				},
			}},
			Binding: boltzcomputeapi.PredictionStructureAndBindingStartParamsInputBindingUnion{
				OfPredictionStructureAndBindingStartsInputBindingLigandProteinBinding: &boltzcomputeapi.PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding{
					BinderChainID: "binder_chain_id",
				},
			},
			Bonds: []boltzcomputeapi.PredictionStructureAndBindingStartParamsInputBond{{
				Atom1: boltzcomputeapi.PredictionStructureAndBindingStartParamsInputBondAtom1Union{
					OfPredictionStructureAndBindingStartsInputBondAtom1LigandAtom: &boltzcomputeapi.PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: boltzcomputeapi.PredictionStructureAndBindingStartParamsInputBondAtom2Union{
					OfPredictionStructureAndBindingStartsInputBondAtom2LigandAtom: &boltzcomputeapi.PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []boltzcomputeapi.PredictionStructureAndBindingStartParamsInputConstraintUnion{{
				OfPredictionStructureAndBindingStartsInputConstraintPocketConstraint: &boltzcomputeapi.PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               boltzcomputeapi.Bool(true),
				},
			}},
			ModelOptions: boltzcomputeapi.PredictionStructureAndBindingStartParamsInputModelOptions{
				RecyclingSteps: boltzcomputeapi.Int(1),
				SamplingSteps:  boltzcomputeapi.Int(1),
				StepScale:      boltzcomputeapi.Float(1.3),
			},
			NumSamples: boltzcomputeapi.Int(1),
		},
		Model:          "boltz-2.1",
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
