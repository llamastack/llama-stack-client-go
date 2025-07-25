// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// AgentService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
	Session AgentSessionService
	Steps   AgentStepService
	Turn    AgentTurnService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	r.Session = NewAgentSessionService(opts...)
	r.Steps = NewAgentStepService(opts...)
	r.Turn = NewAgentTurnService(opts...)
	return
}

// Create an agent with the given configuration.
func (r *AgentService) New(ctx context.Context, body AgentNewParams, opts ...option.RequestOption) (res *AgentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Describe an agent by its ID.
func (r *AgentService) Get(ctx context.Context, agentID string, opts ...option.RequestOption) (res *AgentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all agents.
func (r *AgentService) List(ctx context.Context, query AgentListParams, opts ...option.RequestOption) (res *AgentListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an agent by its ID and its associated sessions and turns.
func (r *AgentService) Delete(ctx context.Context, agentID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// An inference step in an agent turn.
type InferenceStep struct {
	// The response from the LLM.
	ModelResponse shared.CompletionMessage `json:"model_response,required"`
	// The ID of the step.
	StepID string `json:"step_id,required"`
	// Type of the step in an agent turn.
	StepType constant.Inference `json:"step_type,required"`
	// The ID of the turn.
	TurnID string `json:"turn_id,required"`
	// The time the step completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// The time the step started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelResponse respjson.Field
		StepID        respjson.Field
		StepType      respjson.Field
		TurnID        respjson.Field
		CompletedAt   respjson.Field
		StartedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceStep) RawJSON() string { return r.JSON.raw }
func (r *InferenceStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A memory retrieval step in an agent turn.
type MemoryRetrievalStep struct {
	// The context retrieved from the vector databases.
	InsertedContext shared.InterleavedContentUnion `json:"inserted_context,required"`
	// The ID of the step.
	StepID string `json:"step_id,required"`
	// Type of the step in an agent turn.
	StepType constant.MemoryRetrieval `json:"step_type,required"`
	// The ID of the turn.
	TurnID string `json:"turn_id,required"`
	// The IDs of the vector databases to retrieve context from.
	VectorDBIDs string `json:"vector_db_ids,required"`
	// The time the step completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// The time the step started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InsertedContext respjson.Field
		StepID          respjson.Field
		StepType        respjson.Field
		TurnID          respjson.Field
		VectorDBIDs     respjson.Field
		CompletedAt     respjson.Field
		StartedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryRetrievalStep) RawJSON() string { return r.JSON.raw }
func (r *MemoryRetrievalStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A shield call step in an agent turn.
type ShieldCallStep struct {
	// The ID of the step.
	StepID string `json:"step_id,required"`
	// Type of the step in an agent turn.
	StepType constant.ShieldCall `json:"step_type,required"`
	// The ID of the turn.
	TurnID string `json:"turn_id,required"`
	// The time the step completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// The time the step started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// The violation from the shield call.
	Violation shared.SafetyViolation `json:"violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StepID      respjson.Field
		StepType    respjson.Field
		TurnID      respjson.Field
		CompletedAt respjson.Field
		StartedAt   respjson.Field
		Violation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ShieldCallStep) RawJSON() string { return r.JSON.raw }
func (r *ShieldCallStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tool execution step in an agent turn.
type ToolExecutionStep struct {
	// The ID of the step.
	StepID string `json:"step_id,required"`
	// Type of the step in an agent turn.
	StepType constant.ToolExecution `json:"step_type,required"`
	// The tool calls to execute.
	ToolCalls []shared.ToolCall `json:"tool_calls,required"`
	// The tool responses from the tool calls.
	ToolResponses []ToolResponse `json:"tool_responses,required"`
	// The ID of the turn.
	TurnID string `json:"turn_id,required"`
	// The time the step completed.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// The time the step started.
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StepID        respjson.Field
		StepType      respjson.Field
		ToolCalls     respjson.Field
		ToolResponses respjson.Field
		TurnID        respjson.Field
		CompletedAt   respjson.Field
		StartedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolExecutionStep) RawJSON() string { return r.JSON.raw }
func (r *ToolExecutionStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolResponse struct {
	CallID string `json:"call_id,required"`
	// A image content item
	Content  shared.InterleavedContentUnion       `json:"content,required"`
	ToolName ToolResponseToolName                 `json:"tool_name,required"`
	Metadata map[string]ToolResponseMetadataUnion `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Content     respjson.Field
		ToolName    respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolResponse to a ToolResponseParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolResponseParam.Overrides()
func (r ToolResponse) ToParam() ToolResponseParam {
	return param.Override[ToolResponseParam](json.RawMessage(r.RawJSON()))
}

type ToolResponseToolName string

const (
	ToolResponseToolNameBraveSearch     ToolResponseToolName = "brave_search"
	ToolResponseToolNameWolframAlpha    ToolResponseToolName = "wolfram_alpha"
	ToolResponseToolNamePhotogen        ToolResponseToolName = "photogen"
	ToolResponseToolNameCodeInterpreter ToolResponseToolName = "code_interpreter"
)

// ToolResponseMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ToolResponseMetadataUnion struct {
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

func (u ToolResponseMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResponseMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResponseMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResponseMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolResponseMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolResponseMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CallID, Content, ToolName are required.
type ToolResponseParam struct {
	CallID string `json:"call_id,required"`
	// A image content item
	Content  shared.InterleavedContentUnionParam       `json:"content,omitzero,required"`
	ToolName ToolResponseToolName                      `json:"tool_name,omitzero,required"`
	Metadata map[string]ToolResponseMetadataUnionParam `json:"metadata,omitzero"`
	paramObj
}

func (r ToolResponseParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolResponseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolResponseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolResponseMetadataUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ToolResponseMetadataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ToolResponseMetadataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolResponseMetadataUnionParam) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

type AgentNewResponse struct {
	AgentID string `json:"agent_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponse struct {
	// Configuration for an agent.
	AgentConfig shared.AgentConfig `json:"agent_config,required"`
	AgentID     string             `json:"agent_id,required"`
	CreatedAt   time.Time          `json:"created_at,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentConfig respjson.Field
		AgentID     respjson.Field
		CreatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A generic paginated response that follows a simple format.
type AgentListResponse struct {
	// The list of items for the current page
	Data []map[string]AgentListResponseDataUnion `json:"data,required"`
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
func (r AgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentListResponseDataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type AgentListResponseDataUnion struct {
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

func (u AgentListResponseDataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentListResponseDataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentListResponseDataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentListResponseDataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentListResponseDataUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentListResponseDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewParams struct {
	// The configuration for the agent.
	AgentConfig shared.AgentConfigParam `json:"agent_config,omitzero,required"`
	paramObj
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListParams struct {
	// The number of agents to return.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The index to start the pagination from.
	StartIndex param.Opt[int64] `query:"start_index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
