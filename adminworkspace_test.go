// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcompute_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/boltz-bio/boltz-api-go"
	"github.com/boltz-bio/boltz-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-api-go/option"
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
	client := boltzcompute.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Admin.Workspaces.New(context.TODO(), boltzcompute.AdminWorkspaceNewParams{
		DataRetention: boltzcompute.AdminWorkspaceNewParamsDataRetention{
			Unit:  boltzcompute.AdminWorkspaceNewParamsDataRetentionUnitHours,
			Value: 1,
		},
		Name: boltzcompute.String("x"),
	})
	if err != nil {
		var apierr *boltzcompute.Error
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
	client := boltzcompute.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Admin.Workspaces.Get(context.TODO(), "workspace_id")
	if err != nil {
		var apierr *boltzcompute.Error
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
	client := boltzcompute.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Admin.Workspaces.Update(
		context.TODO(),
		"workspace_id",
		boltzcompute.AdminWorkspaceUpdateParams{
			DataRetention: boltzcompute.AdminWorkspaceUpdateParamsDataRetention{
				Unit:  boltzcompute.AdminWorkspaceUpdateParamsDataRetentionUnitHours,
				Value: 1,
			},
			Name: boltzcompute.String("x"),
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

func TestAdminWorkspaceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.Workspaces.List(context.TODO(), boltzcompute.AdminWorkspaceListParams{
		AfterID:  boltzcompute.String("after_id"),
		BeforeID: boltzcompute.String("before_id"),
		Limit:    boltzcompute.Int(1),
	})
	if err != nil {
		var apierr *boltzcompute.Error
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
	client := boltzcompute.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Admin.Workspaces.Archive(context.TODO(), "workspace_id")
	if err != nil {
		var apierr *boltzcompute.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
