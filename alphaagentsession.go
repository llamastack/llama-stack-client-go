// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// AlphaAgentSessionService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaAgentSessionService] method instead.
type AlphaAgentSessionService struct {
	Options []option.RequestOption
}

// NewAlphaAgentSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlphaAgentSessionService(opts ...option.RequestOption) (r AlphaAgentSessionService) {
	r = AlphaAgentSessionService{}
	r.Options = opts
	return
}

// Create a new session for an agent.
func (r *AlphaAgentSessionService) New(ctx context.Context, agentID string, body AlphaAgentSessionNewParams, opts ...option.RequestOption) (res *AlphaAgentSessionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an agent session by its ID.
func (r *AlphaAgentSessionService) Get(ctx context.Context, sessionID string, params AlphaAgentSessionGetParams, opts ...option.RequestOption) (res *Session, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s", params.AgentID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// List all session(s) of a given agent.
func (r *AlphaAgentSessionService) List(ctx context.Context, agentID string, query AlphaAgentSessionListParams, opts ...option.RequestOption) (res *AlphaAgentSessionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/sessions", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an agent session by its ID and its associated turns.
func (r *AlphaAgentSessionService) Delete(ctx context.Context, sessionID string, body AlphaAgentSessionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s", body.AgentID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// A single session of an interaction with an Agentic System.
type Session struct {
	// Unique identifier for the conversation session
	SessionID string `json:"session_id,required"`
	// Human-readable name for the session
	SessionName string `json:"session_name,required"`
	// Timestamp when the session was created
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// List of all turns that have occurred in this session
	Turns []Turn `json:"turns,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SessionID   respjson.Field
		SessionName respjson.Field
		StartedAt   respjson.Field
		Turns       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Session) RawJSON() string { return r.JSON.raw }
func (r *Session) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response returned when creating a new agent session.
type AlphaAgentSessionNewResponse struct {
	// Unique identifier for the created session
	SessionID string `json:"session_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SessionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaAgentSessionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaAgentSessionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A generic paginated response that follows a simple format.
type AlphaAgentSessionListResponse struct {
	// The list of items for the current page
	Data []map[string]AlphaAgentSessionListResponseDataUnion `json:"data,required"`
	// Whether there are more items available after this set
	HasMore bool `json:"has_more,required"`
	// The URL for accessing this list
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AlphaAgentSessionListResponse) RawJSON() string { return r.JSON.raw }
func (r *AlphaAgentSessionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AlphaAgentSessionListResponseDataUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type AlphaAgentSessionListResponseDataUnion struct {
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

func (u AlphaAgentSessionListResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AlphaAgentSessionListResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AlphaAgentSessionListResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AlphaAgentSessionListResponseDataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AlphaAgentSessionListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *AlphaAgentSessionListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaAgentSessionNewParams struct {
	// The name of the session to create.
	SessionName string `json:"session_name,required"`
	paramObj
}

func (r AlphaAgentSessionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentSessionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaAgentSessionGetParams struct {
	AgentID string `path:"agent_id,required" json:"-"`
	// (Optional) List of turn IDs to filter the session by.
	TurnIDs []string `query:"turn_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaAgentSessionGetParams]'s query parameters as
// `url.Values`.
func (r AlphaAgentSessionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlphaAgentSessionListParams struct {
	// The number of sessions to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index to start the pagination from.
	StartIndex param.Opt[int64] `query:"start_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AlphaAgentSessionListParams]'s query parameters as
// `url.Values`.
func (r AlphaAgentSessionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AlphaAgentSessionDeleteParams struct {
	AgentID string `path:"agent_id,required" json:"-"`
	paramObj
}
