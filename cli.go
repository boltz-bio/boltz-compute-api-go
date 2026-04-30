// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package boltzcompute

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/boltz-bio/boltz-compute-api-go/internal/apijson"
	"github.com/boltz-bio/boltz-compute-api-go/internal/apiquery"
	"github.com/boltz-bio/boltz-compute-api-go/internal/requestconfig"
	"github.com/boltz-bio/boltz-compute-api-go/option"
	"github.com/boltz-bio/boltz-compute-api-go/packages/param"
	"github.com/boltz-bio/boltz-compute-api-go/packages/respjson"
)

// Check the installed boltz-api CLI version against the currently published CLI
// release and the minimum version supported by the Compute API.
//
// CliService contains methods and other services that help with interacting with
// the boltz-compute API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCliService] method instead.
type CliService struct {
	Options []option.RequestOption
}

// NewCliService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCliService(opts ...option.RequestOption) (r CliService) {
	r = CliService{}
	r.Options = opts
	return
}

// Returns public boltz-api CLI version metadata for lightweight update checks.
func (r *CliService) Version(ctx context.Context, query CliVersionParams, opts ...option.RequestOption) (res *CliVersionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/v1/cli/version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type CliVersionResponse struct {
	Install          CliVersionResponseInstall `json:"install" api:"required"`
	Latest           string                    `json:"latest" api:"required"`
	Message          string                    `json:"message" api:"required"`
	MinimumSupported string                    `json:"minimum_supported" api:"required"`
	UpdateAvailable  bool                      `json:"update_available" api:"required"`
	UpdateRequired   bool                      `json:"update_required" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Install          respjson.Field
		Latest           respjson.Field
		Message          respjson.Field
		MinimumSupported respjson.Field
		UpdateAvailable  respjson.Field
		UpdateRequired   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CliVersionResponse) RawJSON() string { return r.JSON.raw }
func (r *CliVersionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CliVersionResponseInstall struct {
	MacosLinux string `json:"macos_linux" api:"required"`
	Windows    string `json:"windows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MacosLinux  respjson.Field
		Windows     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CliVersionResponseInstall) RawJSON() string { return r.JSON.raw }
func (r *CliVersionResponseInstall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CliVersionParams struct {
	Current  param.Opt[string] `query:"current,omitzero" json:"-"`
	Platform param.Opt[string] `query:"platform,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CliVersionParams]'s query parameters as `url.Values`.
func (r CliVersionParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
