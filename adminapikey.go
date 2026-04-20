// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcomputeapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/boltz-compute-api-go/internal/apijson"
	"github.com/stainless-sdks/boltz-compute-api-go/internal/apiquery"
	"github.com/stainless-sdks/boltz-compute-api-go/internal/requestconfig"
	"github.com/stainless-sdks/boltz-compute-api-go/option"
	"github.com/stainless-sdks/boltz-compute-api-go/packages/pagination"
	"github.com/stainless-sdks/boltz-compute-api-go/packages/param"
	"github.com/stainless-sdks/boltz-compute-api-go/packages/respjson"
	"github.com/stainless-sdks/boltz-compute-api-go/shared/constant"
)

// API keys authenticate requests to the Compute API. There are two key types:
// admin keys have full access to all management and compute operations across the
// organization, while workspace keys are scoped to a single workspace and can only
// perform compute operations (predictions, protein design, small molecule design)
// within that workspace. Keys can be created in live or test mode. Test keys
// (prefixed `sk_bc_*_test_`) create test-mode resources with synthetic data and no
// GPU cost. Every resource includes a `livemode` field indicating its mode.
//
// AdminAPIKeyService contains methods and other services that help with
// interacting with the boltz-compute API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminAPIKeyService] method instead.
type AdminAPIKeyService struct {
	options []option.RequestOption
}

// NewAdminAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminAPIKeyService(opts ...option.RequestOption) (r AdminAPIKeyService) {
	r = AdminAPIKeyService{}
	r.options = opts
	return
}

// Create a workspace API key
func (r *AdminAPIKeyService) New(ctx context.Context, body AdminAPIKeyNewParams, opts ...option.RequestOption) (res *AdminAPIKeyNewResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "compute/v1/admin/api-keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List API keys
func (r *AdminAPIKeyService) List(ctx context.Context, query AdminAPIKeyListParams, opts ...option.RequestOption) (res *pagination.CursorPage[AdminAPIKeyListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "compute/v1/admin/api-keys"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List API keys
func (r *AdminAPIKeyService) ListAutoPaging(ctx context.Context, query AdminAPIKeyListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[AdminAPIKeyListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Revoke an API key
func (r *AdminAPIKeyService) Revoke(ctx context.Context, apiKeyID string, opts ...option.RequestOption) (res *AdminAPIKeyRevokeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("compute/v1/admin/api-keys/%s/revoke", url.PathEscape(apiKeyID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type AdminAPIKeyNewResponse struct {
	// The full API key. This is only shown once — store it securely.
	Key        string                           `json:"key" api:"required"`
	KeyDetails AdminAPIKeyNewResponseKeyDetails `json:"key_details" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		KeyDetails  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminAPIKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminAPIKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminAPIKeyNewResponseKeyDetails struct {
	// API key ID
	ID string `json:"id" api:"required"`
	// IP addresses allowed to use this key. An empty array means all IPs are allowed.
	AllowedIPs []string  `json:"allowed_ips" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the key expires. Null if the key does not expire.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	IsActive  bool      `json:"is_active" api:"required"`
	// First 12 characters of the key
	KeyPrefix  string             `json:"key_prefix" api:"required"`
	KeyType    constant.Workspace `json:"key_type" default:"workspace"`
	LastUsedAt time.Time          `json:"last_used_at" api:"required" format:"date-time"`
	// Whether this is a live API key (false for test keys).
	Livemode bool `json:"livemode" api:"required"`
	// API key name
	Name string `json:"name" api:"required"`
	// Workspace this key is scoped to
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AllowedIPs  respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		IsActive    respjson.Field
		KeyPrefix   respjson.Field
		KeyType     respjson.Field
		LastUsedAt  respjson.Field
		Livemode    respjson.Field
		Name        respjson.Field
		WorkspaceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminAPIKeyNewResponseKeyDetails) RawJSON() string { return r.JSON.raw }
func (r *AdminAPIKeyNewResponseKeyDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminAPIKeyListResponse struct {
	// API key ID
	ID string `json:"id" api:"required"`
	// IP addresses allowed to use this key. An empty array means all IPs are allowed.
	AllowedIPs []string  `json:"allowed_ips" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the key expires. Null if the key does not expire.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	IsActive  bool      `json:"is_active" api:"required"`
	// First 12 characters of the key
	KeyPrefix string `json:"key_prefix" api:"required"`
	// Any of "admin", "workspace".
	KeyType    AdminAPIKeyListResponseKeyType `json:"key_type" api:"required"`
	LastUsedAt time.Time                      `json:"last_used_at" api:"required" format:"date-time"`
	// Whether this is a live API key (false for test keys).
	Livemode bool `json:"livemode" api:"required"`
	// API key name
	Name string `json:"name" api:"required"`
	// Workspace ID if workspace-scoped
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AllowedIPs  respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		IsActive    respjson.Field
		KeyPrefix   respjson.Field
		KeyType     respjson.Field
		LastUsedAt  respjson.Field
		Livemode    respjson.Field
		Name        respjson.Field
		WorkspaceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminAPIKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminAPIKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminAPIKeyListResponseKeyType string

const (
	AdminAPIKeyListResponseKeyTypeAdmin     AdminAPIKeyListResponseKeyType = "admin"
	AdminAPIKeyListResponseKeyTypeWorkspace AdminAPIKeyListResponseKeyType = "workspace"
)

type AdminAPIKeyRevokeResponse struct {
	// API key ID
	ID string `json:"id" api:"required"`
	// IP addresses allowed to use this key. An empty array means all IPs are allowed.
	AllowedIPs []string  `json:"allowed_ips" api:"required"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	// When the key expires. Null if the key does not expire.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	IsActive  bool      `json:"is_active" api:"required"`
	// First 12 characters of the key
	KeyPrefix string `json:"key_prefix" api:"required"`
	// Any of "admin", "workspace".
	KeyType    AdminAPIKeyRevokeResponseKeyType `json:"key_type" api:"required"`
	LastUsedAt time.Time                        `json:"last_used_at" api:"required" format:"date-time"`
	// Whether this is a live API key (false for test keys).
	Livemode bool `json:"livemode" api:"required"`
	// API key name
	Name string `json:"name" api:"required"`
	// Workspace ID if workspace-scoped
	WorkspaceID string `json:"workspace_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AllowedIPs  respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		IsActive    respjson.Field
		KeyPrefix   respjson.Field
		KeyType     respjson.Field
		LastUsedAt  respjson.Field
		Livemode    respjson.Field
		Name        respjson.Field
		WorkspaceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AdminAPIKeyRevokeResponse) RawJSON() string { return r.JSON.raw }
func (r *AdminAPIKeyRevokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AdminAPIKeyRevokeResponseKeyType string

const (
	AdminAPIKeyRevokeResponseKeyTypeAdmin     AdminAPIKeyRevokeResponseKeyType = "admin"
	AdminAPIKeyRevokeResponseKeyTypeWorkspace AdminAPIKeyRevokeResponseKeyType = "workspace"
)

type AdminAPIKeyNewParams struct {
	// API key name
	Name string `json:"name" api:"required"`
	// Days until the key expires. Omit for a key that does not expire.
	ExpiresInDays param.Opt[int64] `json:"expires_in_days,omitzero"`
	// Workspace ID to scope this key to. Omit for default workspace.
	WorkspaceID param.Opt[string] `json:"workspace_id,omitzero"`
	// IP addresses allowed to use this key (IPv4 or IPv6). An empty array (the
	// default) means all IPs are allowed.
	AllowedIPs []string `json:"allowed_ips,omitzero"`
	// Key mode. Test keys create test-mode resources with synthetic data.
	//
	// Any of "live", "test".
	Mode AdminAPIKeyNewParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r AdminAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AdminAPIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AdminAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Key mode. Test keys create test-mode resources with synthetic data.
type AdminAPIKeyNewParamsMode string

const (
	AdminAPIKeyNewParamsModeLive AdminAPIKeyNewParamsMode = "live"
	AdminAPIKeyNewParamsModeTest AdminAPIKeyNewParamsMode = "test"
)

type AdminAPIKeyListParams struct {
	// Return results after this ID
	AfterID param.Opt[string] `query:"after_id,omitzero" json:"-"`
	// Return results before this ID
	BeforeID param.Opt[string] `query:"before_id,omitzero" json:"-"`
	// Max items to return
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter by workspace ID. If not provided, returns keys across all workspaces.
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AdminAPIKeyListParams]'s query parameters as `url.Values`.
func (r AdminAPIKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
