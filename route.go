// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
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

// List all routes.
func (r *RouteService) List(ctx context.Context, opts ...option.RequestOption) (res *[]RouteInfo, err error) {
	var env ListRoutesResponse
	opts = append(r.Options[:], opts...)
	path := "v1/inspect/routes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

type ListRoutesResponse struct {
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
