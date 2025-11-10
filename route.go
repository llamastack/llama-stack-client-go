// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// RouteService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRouteService] method instead.
type RouteService struct {
	Options []option.RequestOption
}

// NewRouteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRouteService(opts ...option.RequestOption) (r RouteService) {
	r = RouteService{}
	r.Options = opts
	return
}

// List routes. List all available API routes with their methods and implementing
// providers.
func (r *RouteService) List(ctx context.Context, query RouteListParams, opts ...option.RequestOption) (res *[]RouteInfo, err error) {
	var env ListRoutesResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1/inspect/routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Response containing a list of all available API routes.
type ListRoutesResponse struct {
	// List of available route information objects
	Data []RouteInfo `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListRoutesResponse) RawJSON() string { return r.JSON.raw }
func (r *ListRoutesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteListParams struct {
	// Optional filter to control which routes are returned. Can be an API level ('v1',
	// 'v1alpha', 'v1beta') to show non-deprecated routes at that level, or
	// 'deprecated' to show deprecated routes across all levels. If not specified,
	// returns all non-deprecated routes.
	//
	// Any of "v1", "v1alpha", "v1beta", "deprecated".
	APIFilter RouteListParamsAPIFilter `query:"api_filter,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [RouteListParams]'s query parameters as `url.Values`.
func (r RouteListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Optional filter to control which routes are returned. Can be an API level ('v1',
// 'v1alpha', 'v1beta') to show non-deprecated routes at that level, or
// 'deprecated' to show deprecated routes across all levels. If not specified,
// returns all non-deprecated routes.
type RouteListParamsAPIFilter string

const (
	RouteListParamsAPIFilterV1         RouteListParamsAPIFilter = "v1"
	RouteListParamsAPIFilterV1alpha    RouteListParamsAPIFilter = "v1alpha"
	RouteListParamsAPIFilterV1beta     RouteListParamsAPIFilter = "v1beta"
	RouteListParamsAPIFilterDeprecated RouteListParamsAPIFilter = "deprecated"
)
