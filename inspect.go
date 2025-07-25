// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// InspectService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInspectService] method instead.
type InspectService struct {
	Options []option.RequestOption
}

// NewInspectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInspectService(opts ...option.RequestOption) (r InspectService) {
	r = InspectService{}
	r.Options = opts
	return
}

// Get the health of the service.
func (r *InspectService) Health(ctx context.Context, opts ...option.RequestOption) (res *HealthInfo, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the version of the service.
func (r *InspectService) Version(ctx context.Context, opts ...option.RequestOption) (res *VersionInfo, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthInfo struct {
	// Any of "OK", "Error", "Not Implemented".
	Status HealthInfoStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthInfo) RawJSON() string { return r.JSON.raw }
func (r *HealthInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HealthInfoStatus string

const (
	HealthInfoStatusOk             HealthInfoStatus = "OK"
	HealthInfoStatusError          HealthInfoStatus = "Error"
	HealthInfoStatusNotImplemented HealthInfoStatus = "Not Implemented"
)

type ProviderInfo struct {
	API          string                             `json:"api,required"`
	Config       map[string]ProviderInfoConfigUnion `json:"config,required"`
	Health       map[string]ProviderInfoHealthUnion `json:"health,required"`
	ProviderID   string                             `json:"provider_id,required"`
	ProviderType string                             `json:"provider_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		API          respjson.Field
		Config       respjson.Field
		Health       respjson.Field
		ProviderID   respjson.Field
		ProviderType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProviderInfo) RawJSON() string { return r.JSON.raw }
func (r *ProviderInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProviderInfoConfigUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ProviderInfoConfigUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ProviderInfoConfigUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoConfigUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoConfigUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoConfigUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProviderInfoConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *ProviderInfoConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ProviderInfoHealthUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ProviderInfoHealthUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ProviderInfoHealthUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoHealthUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoHealthUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ProviderInfoHealthUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ProviderInfoHealthUnion) RawJSON() string { return u.JSON.raw }

func (r *ProviderInfoHealthUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RouteInfo struct {
	Method        string   `json:"method,required"`
	ProviderTypes []string `json:"provider_types,required"`
	Route         string   `json:"route,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method        respjson.Field
		ProviderTypes respjson.Field
		Route         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RouteInfo) RawJSON() string { return r.JSON.raw }
func (r *RouteInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VersionInfo struct {
	Version string `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VersionInfo) RawJSON() string { return r.JSON.raw }
func (r *VersionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
