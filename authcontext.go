// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcompute

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/boltz-bio/boltz-compute-api-go/internal/apijson"
	"github.com/boltz-bio/boltz-compute-api-go/internal/requestconfig"
	"github.com/boltz-bio/boltz-compute-api-go/option"
	"github.com/boltz-bio/boltz-compute-api-go/packages/respjson"
	"github.com/boltz-bio/boltz-compute-api-go/shared/constant"
)

// Inspect the authentication context for the current credential, including the
// organization or workspace scope for API keys and the available organization
// memberships for OAuth bearer tokens. OAuth callers can use this information to
// choose which organization to send with future requests.
//
// AuthContextService contains methods and other services that help with
// interacting with the boltz-compute API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthContextService] method instead.
type AuthContextService struct {
	Options []option.RequestOption
}

// NewAuthContextService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuthContextService(opts ...option.RequestOption) (r AuthContextService) {
	r = AuthContextService{}
	r.Options = opts
	return
}

// Returns the organization context available to the current API key or OAuth
// bearer token. OAuth callers can use X-Boltz-Organization-Id to select one of
// their organizations.
func (r *AuthContextService) Me(ctx context.Context, opts ...option.RequestOption) (res *AuthContextMeResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/auth/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// AuthContextMeResponseUnion contains all possible properties and values from
// [AuthContextMeResponseAuthMeAPIKeyResponse],
// [AuthContextMeResponseAuthMeUserResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AuthContextMeResponseUnion struct {
	// This field is from variant [AuthContextMeResponseAuthMeAPIKeyResponse].
	APIKeyID string `json:"api_key_id"`
	// This field is from variant [AuthContextMeResponseAuthMeAPIKeyResponse].
	KeyType AuthContextMeResponseAuthMeAPIKeyResponseKeyType `json:"key_type"`
	// This field is from variant [AuthContextMeResponseAuthMeAPIKeyResponse].
	Mode AuthContextMeResponseAuthMeAPIKeyResponseMode `json:"mode"`
	// This field is from variant [AuthContextMeResponseAuthMeAPIKeyResponse].
	OrganizationID         string `json:"organization_id"`
	PrincipalType          string `json:"principal_type"`
	SelectedOrganizationID string `json:"selected_organization_id"`
	// This field is from variant [AuthContextMeResponseAuthMeAPIKeyResponse].
	WorkspaceID string `json:"workspace_id"`
	// This field is from variant [AuthContextMeResponseAuthMeUserResponse].
	ActiveOrganizationID string `json:"active_organization_id"`
	// This field is from variant [AuthContextMeResponseAuthMeUserResponse].
	OrganizationMemberships []AuthContextMeResponseAuthMeUserResponseOrganizationMembership `json:"organization_memberships"`
	// This field is from variant [AuthContextMeResponseAuthMeUserResponse].
	UserID string `json:"user_id"`
	JSON   struct {
		APIKeyID                respjson.Field
		KeyType                 respjson.Field
		Mode                    respjson.Field
		OrganizationID          respjson.Field
		PrincipalType           respjson.Field
		SelectedOrganizationID  respjson.Field
		WorkspaceID             respjson.Field
		ActiveOrganizationID    respjson.Field
		OrganizationMemberships respjson.Field
		UserID                  respjson.Field
		raw                     string
	} `json:"-"`
}

func (u AuthContextMeResponseUnion) AsAuthContextMeResponseAuthMeAPIKeyResponse() (v AuthContextMeResponseAuthMeAPIKeyResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuthContextMeResponseUnion) AsAuthContextMeResponseAuthMeUserResponse() (v AuthContextMeResponseAuthMeUserResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AuthContextMeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AuthContextMeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthContextMeResponseAuthMeAPIKeyResponse struct {
	APIKeyID string `json:"api_key_id" api:"required"`
	// Any of "admin", "workspace".
	KeyType AuthContextMeResponseAuthMeAPIKeyResponseKeyType `json:"key_type" api:"required"`
	// Any of "live", "test".
	Mode                   AuthContextMeResponseAuthMeAPIKeyResponseMode `json:"mode" api:"required"`
	OrganizationID         string                                        `json:"organization_id" api:"required"`
	PrincipalType          constant.APIKey                               `json:"principal_type" default:"api_key"`
	SelectedOrganizationID string                                        `json:"selected_organization_id" api:"required"`
	WorkspaceID            string                                        `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeyID               respjson.Field
		KeyType                respjson.Field
		Mode                   respjson.Field
		OrganizationID         respjson.Field
		PrincipalType          respjson.Field
		SelectedOrganizationID respjson.Field
		WorkspaceID            respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthContextMeResponseAuthMeAPIKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthContextMeResponseAuthMeAPIKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthContextMeResponseAuthMeAPIKeyResponseKeyType string

const (
	AuthContextMeResponseAuthMeAPIKeyResponseKeyTypeAdmin     AuthContextMeResponseAuthMeAPIKeyResponseKeyType = "admin"
	AuthContextMeResponseAuthMeAPIKeyResponseKeyTypeWorkspace AuthContextMeResponseAuthMeAPIKeyResponseKeyType = "workspace"
)

type AuthContextMeResponseAuthMeAPIKeyResponseMode string

const (
	AuthContextMeResponseAuthMeAPIKeyResponseModeLive AuthContextMeResponseAuthMeAPIKeyResponseMode = "live"
	AuthContextMeResponseAuthMeAPIKeyResponseModeTest AuthContextMeResponseAuthMeAPIKeyResponseMode = "test"
)

type AuthContextMeResponseAuthMeUserResponse struct {
	ActiveOrganizationID    string                                                          `json:"active_organization_id" api:"required"`
	OrganizationMemberships []AuthContextMeResponseAuthMeUserResponseOrganizationMembership `json:"organization_memberships" api:"required"`
	PrincipalType           constant.User                                                   `json:"principal_type" default:"user"`
	SelectedOrganizationID  string                                                          `json:"selected_organization_id" api:"required"`
	UserID                  string                                                          `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ActiveOrganizationID    respjson.Field
		OrganizationMemberships respjson.Field
		PrincipalType           respjson.Field
		SelectedOrganizationID  respjson.Field
		UserID                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthContextMeResponseAuthMeUserResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthContextMeResponseAuthMeUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthContextMeResponseAuthMeUserResponseOrganizationMembership struct {
	OrganizationID string `json:"organization_id" api:"required"`
	// Any of "Admin", "Scientist", "Viewer".
	Role AuthContextMeResponseAuthMeUserResponseOrganizationMembershipRole `json:"role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrganizationID respjson.Field
		Role           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthContextMeResponseAuthMeUserResponseOrganizationMembership) RawJSON() string {
	return r.JSON.raw
}
func (r *AuthContextMeResponseAuthMeUserResponseOrganizationMembership) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthContextMeResponseAuthMeUserResponseOrganizationMembershipRole string

const (
	AuthContextMeResponseAuthMeUserResponseOrganizationMembershipRoleAdmin     AuthContextMeResponseAuthMeUserResponseOrganizationMembershipRole = "Admin"
	AuthContextMeResponseAuthMeUserResponseOrganizationMembershipRoleScientist AuthContextMeResponseAuthMeUserResponseOrganizationMembershipRole = "Scientist"
	AuthContextMeResponseAuthMeUserResponseOrganizationMembershipRoleViewer    AuthContextMeResponseAuthMeUserResponseOrganizationMembershipRole = "Viewer"
)

type AuthContextMeResponseKeyType string

const (
	AuthContextMeResponseKeyTypeAdmin     AuthContextMeResponseKeyType = "admin"
	AuthContextMeResponseKeyTypeWorkspace AuthContextMeResponseKeyType = "workspace"
)

type AuthContextMeResponseMode string

const (
	AuthContextMeResponseModeLive AuthContextMeResponseMode = "live"
	AuthContextMeResponseModeTest AuthContextMeResponseMode = "test"
)
