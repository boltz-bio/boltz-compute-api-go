// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomboltzbioboltzcomputeapigo_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/boltz-bio/boltz-compute-api-go"
	"github.com/boltz-bio/boltz-compute-api-go/internal/testutil"
	"github.com/boltz-bio/boltz-compute-api-go/option"
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
	client := githubcomboltzbioboltzcomputeapigo.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Admin.Usage.List(context.TODO(), githubcomboltzbioboltzcomputeapigo.AdminUsageListParams{
		EndingAt:   time.Now(),
		StartingAt: time.Now(),
		WindowSize: githubcomboltzbioboltzcomputeapigo.AdminUsageListParamsWindowSizeHour,
		Applications: githubcomboltzbioboltzcomputeapigo.AdminUsageListParamsApplicationsUnion{
			OfAdminUsageListsApplicationsString: githubcomboltzbioboltzcomputeapigo.Opt(githubcomboltzbioboltzcomputeapigo.AdminUsageListParamsApplicationsStringStructureAndBinding),
		},
		GroupBy: githubcomboltzbioboltzcomputeapigo.AdminUsageListParamsGroupByUnion{
			OfAdminUsageListsGroupByString: githubcomboltzbioboltzcomputeapigo.Opt(githubcomboltzbioboltzcomputeapigo.AdminUsageListParamsGroupByStringWorkspaceID),
		},
		Limit: githubcomboltzbioboltzcomputeapigo.Int(1),
		Page:  githubcomboltzbioboltzcomputeapigo.String("page"),
		WorkspaceIDs: githubcomboltzbioboltzcomputeapigo.AdminUsageListParamsWorkspaceIDsUnion{
			OfString: githubcomboltzbioboltzcomputeapigo.String("string"),
		},
	})
	if err != nil {
		var apierr *githubcomboltzbioboltzcomputeapigo.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
