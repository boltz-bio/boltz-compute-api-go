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

func TestProteinDesignGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.Get(
		context.TODO(),
		"run_id",
		boltzcompute.ProteinDesignGetParams{
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

func TestProteinDesignListWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.List(context.TODO(), boltzcompute.ProteinDesignListParams{
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

func TestProteinDesignDeleteData(t *testing.T) {
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
	_, err := client.Protein.Design.DeleteData(context.TODO(), "run_id")
	if err != nil {
		var apierr *boltzcompute.Error
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
	client := boltzcompute.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Protein.Design.EstimateCost(context.TODO(), boltzcompute.ProteinDesignEstimateCostParams{
		BinderSpecification: boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationUnion{
			OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpec: &boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpec{
				ChainSelection: map[string]boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
					"B": {
						OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
							CropResidues: boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
							},
							DesignMotifs: []boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
								OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
									DesignLengthRange: boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
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
				Modality: boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
				Structure: boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion{
					OfProteinDesignEstimateCostsBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
						URL: "https://example.com",
					},
				},
				Rules: boltzcompute.ProteinDesignEstimateCostParamsBinderSpecificationStructureTemplateBinderSpecRules{
					ExcludedAminoAcids:     []string{"x"},
					ExcludedSequenceMotifs: []string{"string"},
					MaxHydrophobicFraction: boltzcompute.Float(0),
				},
			},
		},
		NumProteins: 10,
		Target: boltzcompute.ProteinDesignEstimateCostParamsTargetUnion{
			OfProteinDesignEstimateCostsTargetStructureTemplateTarget: &boltzcompute.ProteinDesignEstimateCostParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzcompute.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinDesignEstimateCostsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzcompute.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzcompute.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: boltzcompute.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinDesignEstimateCostsTargetStructureTemplateTargetStructureURLSource: &boltzcompute.ProteinDesignEstimateCostParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
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

func TestProteinDesignListResultsWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.ListResults(
		context.TODO(),
		"run_id",
		boltzcompute.ProteinDesignListResultsParams{
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

func TestProteinDesignStartWithOptionalParams(t *testing.T) {
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
	_, err := client.Protein.Design.Start(context.TODO(), boltzcompute.ProteinDesignStartParams{
		BinderSpecification: boltzcompute.ProteinDesignStartParamsBinderSpecificationUnion{
			OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpec: &boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpec{
				ChainSelection: map[string]boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionUnion{
					"B": {
						OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec: &boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpec{
							CropResidues: boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
							},
							DesignMotifs: []boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifUnion{{
								OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif: &boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotif{
									DesignLengthRange: boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecChainSelectionStructureTemplatePolymerChainSpecDesignMotifReplacementMotifDesignLengthRange{
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
				Modality: boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecModalityPeptide,
				Structure: boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureUnion{
					OfProteinDesignStartsBinderSpecificationStructureTemplateBinderSpecStructureURLSource: &boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecStructureURLSource{
						URL: "https://example.com",
					},
				},
				Rules: boltzcompute.ProteinDesignStartParamsBinderSpecificationStructureTemplateBinderSpecRules{
					ExcludedAminoAcids:     []string{"x"},
					ExcludedSequenceMotifs: []string{"string"},
					MaxHydrophobicFraction: boltzcompute.Float(0),
				},
			},
		},
		NumProteins: 10,
		Target: boltzcompute.ProteinDesignStartParamsTargetUnion{
			OfProteinDesignStartsTargetStructureTemplateTarget: &boltzcompute.ProteinDesignStartParamsTargetStructureTemplateTarget{
				ChainSelection: map[string]boltzcompute.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionUnion{
					"A": {
						OfProteinDesignStartsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec: &boltzcompute.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpec{
							CropResidues: boltzcompute.ProteinDesignStartParamsTargetStructureTemplateTargetChainSelectionStructureTemplateTargetPolymerChainSpecCropResiduesUnion{
								OfIntArray: []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
							},
							EpitopeResidues:  []int64{10, 11, 12},
							FlexibleResidues: []int64{5, 6, 7},
						},
					},
				},
				Structure: boltzcompute.ProteinDesignStartParamsTargetStructureTemplateTargetStructureUnion{
					OfProteinDesignStartsTargetStructureTemplateTargetStructureURLSource: &boltzcompute.ProteinDesignStartParamsTargetStructureTemplateTargetStructureURLSource{
						URL: "https://example.com",
					},
				},
			},
		},
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

func TestProteinDesignStop(t *testing.T) {
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
	_, err := client.Protein.Design.Stop(context.TODO(), "run_id")
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
