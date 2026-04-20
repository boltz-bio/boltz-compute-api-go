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

func TestProteinDesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.Get(
		context.TODO(),
		"run_id",
		boltzcomputeapi.ProteinDesignGetParams{
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

func TestProteinDesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.List(context.TODO(), boltzcomputeapi.ProteinDesignListParams{
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

func TestProteinDesignDeleteData(t *testing.T) {
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
	_, err := client.Protein.Design.DeleteData(context.TODO(), "run_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
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
	client := boltzcomputeapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.Design.EstimateCost(context.TODO(), boltzcomputeapi.ProteinDesignEstimateCostParams{
		BinderSpecification: boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationUnion{
			OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpec: &boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec{
				ChainSelection: map[string]boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
					"B": {
						OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
							CropResidues: boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
							},
							DesignMotifs: []boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
								OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
									DesignLengthRange: boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
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
				Modality: boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
				Structure: boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion{
					OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
						URL: "https://example.com",
					},
				},
				Rules: boltzcomputeapi.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules{
					ExcludedAminoAcids:     []string{"x"},
					ExcludedSequenceMotifs: []string{"string"},
					MaxHydrophobicFraction: boltzcomputeapi.Float(0),
				},
			},
		},
		NumProteins: 10,
		Target: boltzcomputeapi.ProteinDesignEstimateCostParamsTargetUnion{
			OfProteinDesignEstimateCostsTargetStructureTemplateTarget: &boltzcomputeapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzcomputeapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinDesignEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzcomputeapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzcomputeapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: boltzcomputeapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinDesignEstimateCostsTargetStructureTemplateTargetStructureURLSource: &boltzcomputeapi.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
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

func TestProteinDesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.ListResults(
		context.TODO(),
		"run_id",
		boltzcomputeapi.ProteinDesignListResultsParams{
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

func TestProteinDesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.Start(context.TODO(), boltzcomputeapi.ProteinDesignStartParams{
		BinderSpecification: boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationUnion{
			OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpec: &boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec{
				ChainSelection: map[string]boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
					"B": {
						OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
							CropResidues: boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
							},
							DesignMotifs: []boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
								OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
									DesignLengthRange: boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
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
				Modality: boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
				Structure: boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion{
					OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
						URL: "https://example.com",
					},
				},
				Rules: boltzcomputeapi.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules{
					ExcludedAminoAcids:     []string{"x"},
					ExcludedSequenceMotifs: []string{"string"},
					MaxHydrophobicFraction: boltzcomputeapi.Float(0),
				},
			},
		},
		NumProteins: 10,
		Target: boltzcomputeapi.ProteinDesignStartParamsTargetUnion{
			OfProteinDesignStartsTargetStructureTemplateTarget: &boltzcomputeapi.ProteinDesignStartParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzcomputeapi.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinDesignStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzcomputeapi.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzcomputeapi.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: boltzcomputeapi.ProteinDesignStartParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinDesignStartsTargetStructureTemplateTargetStructureURLSource: &boltzcomputeapi.ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
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

func TestProteinDesignStop(t *testing.T) {
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
	_, err := client.Protein.Design.Stop(context.TODO(), "run_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
