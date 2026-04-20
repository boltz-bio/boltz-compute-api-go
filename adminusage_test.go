// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcomputeapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/boltz-compute-api-go"
	"github.com/stainless-sdks/boltz-compute-api-go/internal/testutil"
	"github.com/stainless-sdks/boltz-compute-api-go/option"
)

func TestAdminUsageListWithOptionalParams(t *testing.T) {
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
	_, err := client.Admin.Usage.List(context.TODO(), boltzcomputeapi.AdminUsageListParams{
		EndingAt:   time.Now(),
		StartingAt: time.Now(),
		WindowSize: boltzcomputeapi.AdminUsageListParamsWindowSizeHour,
		Applications: boltzcomputeapi.AdminUsageListParamsApplicationsUnion{
			OfAdminUsageListsApplicationsString: boltzcomputeapi.Opt(boltzcomputeapi.AdminUsageListParamsApplicationsStringStructureAndBinding),
		},
		GroupBy: boltzcomputeapi.AdminUsageListParamsGroupByUnion{
			OfAdminUsageListsGroupByString: boltzcomputeapi.Opt(boltzcomputeapi.AdminUsageListParamsGroupByStringWorkspaceID),
		},
		Limit: boltzcomputeapi.Int(1),
		Page:  boltzcomputeapi.String("page"),
		WorkspaceIDs: boltzcomputeapi.AdminUsageListParamsWorkspaceIDsUnion{
			OfString: boltzcomputeapi.String("string"),
		},
	})
	if err != nil {
		var apierr *boltzcomputeapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
