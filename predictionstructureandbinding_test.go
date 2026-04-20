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

func TestPredictionStructureAndBindingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.Get(
		context.TODO(),
		"id",
		githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingGetParams{
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

func TestPredictionStructureAndBindingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.List(context.TODO(), githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingListParams{
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

func TestPredictionStructureAndBindingDeleteData(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.DeleteData(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Predictions.StructureAndBinding.EstimateCost(context.TODO(), githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParams{
		Input: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInput{
			Entities: []githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputEntityUnion{{
				OfPredictionStructureAndBindingEstimateCostsInputEntityProteinEntity: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModificationUnion{{
						OfPredictionStructureAndBindingEstimateCostsInputEntityProteinEntityModificationCcdModification: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value:  "value",
					Cyclic: githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
			Binding: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputBindingUnion{
				OfPredictionStructureAndBindingEstimateCostsInputBindingLigandProteinBinding: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputBindingLigandProteinBinding{
					BinderChainID: "binder_chain_id",
				},
			},
			Bonds: []githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputBond{{
				Atom1: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputBondAtom1Union{
					OfPredictionStructureAndBindingEstimateCostsInputBondAtom1LigandAtom: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputBondAtom2Union{
					OfPredictionStructureAndBindingEstimateCostsInputBondAtom2LigandAtom: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputConstraintUnion{{
				OfPredictionStructureAndBindingEstimateCostsInputConstraintPocketConstraint: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
			ModelOptions: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingEstimateCostParamsInputModelOptions{
				RecyclingSteps: githubcomboltzbioboltzcomputeapigo.Int(1),
				SamplingSteps:  githubcomboltzbioboltzcomputeapigo.Int(1),
				StepScale:      githubcomboltzbioboltzcomputeapigo.Float(1.3),
			},
			NumSamples: githubcomboltzbioboltzcomputeapigo.Int(1),
		},
		Model:          "boltz-2.1",
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

func TestPredictionStructureAndBindingStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Predictions.StructureAndBinding.Start(context.TODO(), githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParams{
		Input: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInput{
			Entities: []githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputEntityUnion{{
				OfPredictionStructureAndBindingStartsInputEntityProteinEntity: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputEntityProteinEntity{
					ChainIDs: []string{"string"},
					Modifications: []githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputEntityProteinEntityModificationUnion{{
						OfPredictionStructureAndBindingStartsInputEntityProteinEntityModificationCcdModification: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputEntityProteinEntityModificationCcdModification{
							ResidueIndex: 0,
							Value:        "value",
						},
					}},
					Value:  "value",
					Cyclic: githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
			Binding: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputBindingUnion{
				OfPredictionStructureAndBindingStartsInputBindingLigandProteinBinding: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputBindingLigandProteinBinding{
					BinderChainID: "binder_chain_id",
				},
			},
			Bonds: []githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputBond{{
				Atom1: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputBondAtom1Union{
					OfPredictionStructureAndBindingStartsInputBondAtom1LigandAtom: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputBondAtom1LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
				Atom2: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputBondAtom2Union{
					OfPredictionStructureAndBindingStartsInputBondAtom2LigandAtom: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputBondAtom2LigandAtom{
						AtomName: "atom_name",
						ChainID:  "chain_id",
					},
				},
			}},
			Constraints: []githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputConstraintUnion{{
				OfPredictionStructureAndBindingStartsInputConstraintPocketConstraint: &githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputConstraintPocketConstraint{
					BinderChainID: "binder_chain_id",
					ContactResidues: map[string][]int64{
						"A": {42, 43, 44, 67, 68, 69},
					},
					MaxDistanceAngstrom: 0,
					Force:               githubcomboltzbioboltzcomputeapigo.Bool(true),
				},
			}},
			ModelOptions: githubcomboltzbioboltzcomputeapigo.PredictionStructureAndBindingStartParamsInputModelOptions{
				RecyclingSteps: githubcomboltzbioboltzcomputeapigo.Int(1),
				SamplingSteps:  githubcomboltzbioboltzcomputeapigo.Int(1),
				StepScale:      githubcomboltzbioboltzcomputeapigo.Float(1.3),
			},
			NumSamples: githubcomboltzbioboltzcomputeapigo.Int(1),
		},
		Model:          "boltz-2.1",
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
