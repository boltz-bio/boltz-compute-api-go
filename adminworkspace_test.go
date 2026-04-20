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

func TestAdminWorkspaceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.Workspaces.New(context.TODO(), boltzcomputeapi.AdminWorkspaceNewParams{
		DataRetention: boltzcomputeapi.AdminWorkspaceNewParamsDataRetention{
			Unit:  boltzcomputeapi.AdminWorkspaceNewParamsDataRetentionUnitHours,
			Value: 1,
		},
		Name: boltzcomputeapi.String("x"),
	})
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdminWorkspaceGet(t *testing.T) {
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
	_, err := client.Admin.Workspaces.Get(context.TODO(), "workspace_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdminWorkspaceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.Workspaces.Update(
		context.TODO(),
		"workspace_id",
		boltzcomputeapi.AdminWorkspaceUpdateParams{
			DataRetention: boltzcomputeapi.AdminWorkspaceUpdateParamsDataRetention{
				Unit:  boltzcomputeapi.AdminWorkspaceUpdateParamsDataRetentionUnitHours,
				Value: 1,
			},
			Name: boltzcomputeapi.String("x"),
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

func TestAdminWorkspaceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.Workspaces.List(context.TODO(), boltzcomputeapi.AdminWorkspaceListParams{
		AfterID:  boltzcomputeapi.String("after_id"),
		BeforeID: boltzcomputeapi.String("before_id"),
		Limit:    boltzcomputeapi.Int(1),
	})
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAdminWorkspaceArchive(t *testing.T) {
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
	_, err := client.Admin.Workspaces.Archive(context.TODO(), "workspace_id")
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
