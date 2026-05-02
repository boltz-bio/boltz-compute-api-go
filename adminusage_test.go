// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzapi_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/boltz-bio/boltz-api-go"
	"github.com/boltz-bio/boltz-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-api-go/option"
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
	client := boltzapi.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Admin.Usage.List(context.TODO(), boltzapi.AdminUsageListParams{
		EndingAt:   time.Now(),
		StartingAt: time.Now(),
		WindowSize: boltzapi.AdminUsageListParamsWindowSizeHour,
		Applications: boltzapi.AdminUsageListParamsApplicationsUnion{
			OfAdminUsageListsApplicationsString: boltzapi.Opt(boltzapi.AdminUsageListParamsApplicationsStringStructureAndBinding),
		},
		GroupBy: boltzapi.AdminUsageListParamsGroupByUnion{
			OfAdminUsageListsGroupByString: boltzapi.Opt(boltzapi.AdminUsageListParamsGroupByStringWorkspaceID),
		},
		Limit: boltzapi.Int(1),
		Page:  boltzapi.String("page"),
		WorkspaceIDs: boltzapi.AdminUsageListParamsWorkspaceIDsUnion{
			OfString: boltzapi.String("string"),
		},
	})
	if err != nil {
		var apierr *boltzapi.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
