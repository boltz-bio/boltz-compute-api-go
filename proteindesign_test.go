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

func TestProteinDesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.Get(
		context.TODO(),
		"run_id",
		githubcomboltzbioboltzcomputeapigo.ProteinDesignGetParams{
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

func TestProteinDesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.List(context.TODO(), githubcomboltzbioboltzcomputeapigo.ProteinDesignListParams{
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

func TestProteinDesignDeleteData(t *testing.T) {
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
	_, err := client.Protein.Design.DeleteData(context.TODO(), "run_id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestProteinDesignEstimateCostWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.EstimateCost(context.TODO(), githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParams{
		BinderSpecification: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationUnion{
			OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpec: &githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec{
				ChainSelection: map[string]githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
					"B": {
						OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
							CropResidues: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
							},
							DesignMotifs: []githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
								OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
									DesignLengthRange: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
										Max: 8,
										Min: 4,
									},
									EndIndex:   5,
									StartIndex: 0,
								},
							}},
						},
					},
				},
				Modality: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
				Structure: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion{
					OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
						URL: "https://example.com",
					},
				},
				Rules: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules{
					ExcludedAminoAcids:     []string{"x"},
					ExcludedSequenceMotifs: []string{"string"},
					MaxHydrophobicFraction: githubcomboltzbioboltzcomputeapigo.Float(0),
				},
			},
		},
		NumProteins: 10,
		Target: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsTargetUnion{
			OfProteinDesignEstimateCostsTargetStructureTemplateTarget: &githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinDesignEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinDesignEstimateCostsTargetStructureTemplateTargetStructureURLSource: &githubcomboltzbioboltzcomputeapigo.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource{
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

func TestProteinDesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.ListResults(
		context.TODO(),
		"run_id",
		githubcomboltzbioboltzcomputeapigo.ProteinDesignListResultsParams{
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

func TestProteinDesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.Start(context.TODO(), githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParams{
		BinderSpecification: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationUnion{
			OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpec: &githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec{
				ChainSelection: map[string]githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
					"B": {
						OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
							CropResidues: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
							},
							DesignMotifs: []githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
								OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
									DesignLengthRange: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
										Max: 8,
										Min: 4,
									},
									EndIndex:   5,
									StartIndex: 0,
								},
							}},
						},
					},
				},
				Modality: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
				Structure: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion{
					OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
						URL: "https://example.com",
					},
				},
				Rules: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules{
					ExcludedAminoAcids:     []string{"x"},
					ExcludedSequenceMotifs: []string{"string"},
					MaxHydrophobicFraction: githubcomboltzbioboltzcomputeapigo.Float(0),
				},
			},
		},
		NumProteins: 10,
		Target: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsTargetUnion{
			OfProteinDesignStartsTargetStructureTemplateTarget: &githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinDesignStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinDesignStartsTargetStructureTemplateTargetStructureURLSource: &githubcomboltzbioboltzcomputeapigo.ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource{
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

func TestProteinDesignStop(t *testing.T) {
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
	_, err := client.Protein.Design.Stop(context.TODO(), "run_id")
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
