// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"encoding/json"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Configuration for an agent.
type AgentConfig struct {
	// The system instructions for the agent
	Instructions string `json:"instructions,required"`
	// The model identifier to use for the agent
	Model       string    `json:"model,required"`
	ClientTools []ToolDef `json:"client_tools"`
	// Optional flag indicating whether session data has to be persisted
	EnableSessionPersistence bool     `json:"enable_session_persistence"`
	InputShields             []string `json:"input_shields"`
	MaxInferIters            int64    `json:"max_infer_iters"`
	// Optional name for the agent, used in telemetry and identification
	Name          string   `json:"name"`
	OutputShields []string `json:"output_shields"`
	// Optional response format configuration
	ResponseFormat ResponseFormatUnion `json:"response_format"`
	// Sampling parameters.
	SamplingParams SamplingParamsResp `json:"sampling_params"`
	// Whether tool use is required or automatic. This is a hint to the model which may
	// not be followed. It depends on the Instruction Following capabilities of the
	// model.
	//
	// Any of "auto", "required", "none".
	//
	// Deprecated: deprecated
	ToolChoice AgentConfigToolChoice `json:"tool_choice"`
	// Configuration for tool use.
	ToolConfig AgentConfigToolConfig `json:"tool_config"`
	// Prompt format for calling custom / zero shot tools.
	//
	// Any of "json", "function_tag", "python_list".
	//
	// Deprecated: deprecated
	ToolPromptFormat AgentConfigToolPromptFormat `json:"tool_prompt_format"`
	Toolgroups       []AgentConfigToolgroupUnion `json:"toolgroups"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Instructions             respjson.Field
		Model                    respjson.Field
		ClientTools              respjson.Field
		EnableSessionPersistence respjson.Field
		InputShields             respjson.Field
		MaxInferIters            respjson.Field
		Name                     respjson.Field
		OutputShields            respjson.Field
		ResponseFormat           respjson.Field
		SamplingParams           respjson.Field
		ToolChoice               respjson.Field
		ToolConfig               respjson.Field
		ToolPromptFormat         respjson.Field
		Toolgroups               respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentConfig to a AgentConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentConfigParam.Overrides()
func (r AgentConfig) ToParam() AgentConfigParam {
	return param.Override[AgentConfigParam](json.RawMessage(r.RawJSON()))
}

// Whether tool use is required or automatic. This is a hint to the model which may
// not be followed. It depends on the Instruction Following capabilities of the
// model.
type AgentConfigToolChoice string

const (
	AgentConfigToolChoiceAuto     AgentConfigToolChoice = "auto"
	AgentConfigToolChoiceRequired AgentConfigToolChoice = "required"
	AgentConfigToolChoiceNone     AgentConfigToolChoice = "none"
)

// Configuration for tool use.
type AgentConfigToolConfig struct {
	// (Optional) Config for how to override the default system prompt. -
	// `SystemMessageBehavior.append`: Appends the provided system message to the
	// default system prompt. - `SystemMessageBehavior.replace`: Replaces the default
	// system prompt with the provided system message. The system message can include
	// the string '{{function_definitions}}' to indicate where the function definitions
	// should be inserted.
	//
	// Any of "append", "replace".
	SystemMessageBehavior string `json:"system_message_behavior"`
	// (Optional) Whether tool use is automatic, required, or none. Can also specify a
	// tool name to use a specific tool. Defaults to ToolChoice.auto.
	ToolChoice string `json:"tool_choice"`
	// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
	// will attempt to use a format that is best adapted to the model. -
	// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
	// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
	// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
	// are output as Python syntax -- a list of function calls.
	//
	// Any of "json", "function_tag", "python_list".
	ToolPromptFormat string `json:"tool_prompt_format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SystemMessageBehavior respjson.Field
		ToolChoice            respjson.Field
		ToolPromptFormat      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConfigToolConfig) RawJSON() string { return r.JSON.raw }
func (r *AgentConfigToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Prompt format for calling custom / zero shot tools.
type AgentConfigToolPromptFormat string

const (
	AgentConfigToolPromptFormatJson        AgentConfigToolPromptFormat = "json"
	AgentConfigToolPromptFormatFunctionTag AgentConfigToolPromptFormat = "function_tag"
	AgentConfigToolPromptFormatPythonList  AgentConfigToolPromptFormat = "python_list"
)

// AgentConfigToolgroupUnion contains all possible properties and values from
// [string], [AgentConfigToolgroupAgentToolGroupWithArgs].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type AgentConfigToolgroupUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [AgentConfigToolgroupAgentToolGroupWithArgs].
	Args map[string]AgentConfigToolgroupAgentToolGroupWithArgsArgUnion `json:"args"`
	// This field is from variant [AgentConfigToolgroupAgentToolGroupWithArgs].
	Name string `json:"name"`
	JSON struct {
		OfString respjson.Field
		Args     respjson.Field
		Name     respjson.Field
		raw      string
	} `json:"-"`
}

func (u AgentConfigToolgroupUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentConfigToolgroupUnion) AsAgentToolGroupWithArgs() (v AgentConfigToolgroupAgentToolGroupWithArgs) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentConfigToolgroupUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentConfigToolgroupUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentConfigToolgroupAgentToolGroupWithArgs struct {
	Args map[string]AgentConfigToolgroupAgentToolGroupWithArgsArgUnion `json:"args,required"`
	Name string                                                        `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentConfigToolgroupAgentToolGroupWithArgs) RawJSON() string { return r.JSON.raw }
func (r *AgentConfigToolgroupAgentToolGroupWithArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentConfigToolgroupAgentToolGroupWithArgsArgUnion contains all possible
// properties and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type AgentConfigToolgroupAgentToolGroupWithArgsArgUnion struct {
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

func (u AgentConfigToolgroupAgentToolGroupWithArgsArgUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentConfigToolgroupAgentToolGroupWithArgsArgUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentConfigToolgroupAgentToolGroupWithArgsArgUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentConfigToolgroupAgentToolGroupWithArgsArgUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentConfigToolgroupAgentToolGroupWithArgsArgUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentConfigToolgroupAgentToolGroupWithArgsArgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for an agent.
//
// The properties Instructions, Model are required.
type AgentConfigParam struct {
	// The system instructions for the agent
	Instructions string `json:"instructions,required"`
	// The model identifier to use for the agent
	Model string `json:"model,required"`
	// Optional flag indicating whether session data has to be persisted
	EnableSessionPersistence param.Opt[bool]  `json:"enable_session_persistence,omitzero"`
	MaxInferIters            param.Opt[int64] `json:"max_infer_iters,omitzero"`
	// Optional name for the agent, used in telemetry and identification
	Name          param.Opt[string] `json:"name,omitzero"`
	ClientTools   []ToolDefParam    `json:"client_tools,omitzero"`
	InputShields  []string          `json:"input_shields,omitzero"`
	OutputShields []string          `json:"output_shields,omitzero"`
	// Optional response format configuration
	ResponseFormat ResponseFormatUnionParam `json:"response_format,omitzero"`
	// Sampling parameters.
	SamplingParams SamplingParams `json:"sampling_params,omitzero"`
	// Whether tool use is required or automatic. This is a hint to the model which may
	// not be followed. It depends on the Instruction Following capabilities of the
	// model.
	//
	// Any of "auto", "required", "none".
	//
	// Deprecated: deprecated
	ToolChoice AgentConfigToolChoice `json:"tool_choice,omitzero"`
	// Configuration for tool use.
	ToolConfig AgentConfigToolConfigParam `json:"tool_config,omitzero"`
	// Prompt format for calling custom / zero shot tools.
	//
	// Any of "json", "function_tag", "python_list".
	//
	// Deprecated: deprecated
	ToolPromptFormat AgentConfigToolPromptFormat      `json:"tool_prompt_format,omitzero"`
	Toolgroups       []AgentConfigToolgroupUnionParam `json:"toolgroups,omitzero"`
	paramObj
}

func (r AgentConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for tool use.
type AgentConfigToolConfigParam struct {
	// (Optional) Config for how to override the default system prompt. -
	// `SystemMessageBehavior.append`: Appends the provided system message to the
	// default system prompt. - `SystemMessageBehavior.replace`: Replaces the default
	// system prompt with the provided system message. The system message can include
	// the string '{{function_definitions}}' to indicate where the function definitions
	// should be inserted.
	//
	// Any of "append", "replace".
	SystemMessageBehavior string `json:"system_message_behavior,omitzero"`
	// (Optional) Whether tool use is automatic, required, or none. Can also specify a
	// tool name to use a specific tool. Defaults to ToolChoice.auto.
	ToolChoice string `json:"tool_choice,omitzero"`
	// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
	// will attempt to use a format that is best adapted to the model. -
	// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
	// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
	// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
	// are output as Python syntax -- a list of function calls.
	//
	// Any of "json", "function_tag", "python_list".
	ToolPromptFormat string `json:"tool_prompt_format,omitzero"`
	paramObj
}

func (r AgentConfigToolConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConfigToolConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConfigToolConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentConfigToolConfigParam](
		"system_message_behavior", "append", "replace",
	)
	apijson.RegisterFieldValidator[AgentConfigToolConfigParam](
		"tool_prompt_format", "json", "function_tag", "python_list",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentConfigToolgroupUnionParam struct {
	OfString                 param.Opt[string]                                `json:",omitzero,inline"`
	OfAgentToolGroupWithArgs *AgentConfigToolgroupAgentToolGroupWithArgsParam `json:",omitzero,inline"`
	paramUnion
}

func (u AgentConfigToolgroupUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAgentToolGroupWithArgs)
}
func (u *AgentConfigToolgroupUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentConfigToolgroupUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAgentToolGroupWithArgs) {
		return u.OfAgentToolGroupWithArgs
	}
	return nil
}

// The properties Args, Name are required.
type AgentConfigToolgroupAgentToolGroupWithArgsParam struct {
	Args map[string]AgentConfigToolgroupAgentToolGroupWithArgsArgUnionParam `json:"args,omitzero,required"`
	Name string                                                             `json:"name,required"`
	paramObj
}

func (r AgentConfigToolgroupAgentToolGroupWithArgsParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentConfigToolgroupAgentToolGroupWithArgsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentConfigToolgroupAgentToolGroupWithArgsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentConfigToolgroupAgentToolGroupWithArgsArgUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u AgentConfigToolgroupAgentToolGroupWithArgsArgUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *AgentConfigToolgroupAgentToolGroupWithArgsArgUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentConfigToolgroupAgentToolGroupWithArgsArgUnionParam) asAny() any {
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

// A message containing the model's (assistant) response in a chat conversation.
type CompletionMessage struct {
	// The content of the model's response
	Content InterleavedContentUnion `json:"content,required"`
	// Must be "assistant" to identify this as the model's response
	Role constant.Assistant `json:"role,required"`
	// Reason why the model stopped generating. Options are: -
	// `StopReason.end_of_turn`: The model finished generating the entire response. -
	// `StopReason.end_of_message`: The model finished generating but generated a
	// partial response -- usually, a tool call. The user may call the tool and
	// continue the conversation with the tool's response. -
	// `StopReason.out_of_tokens`: The model ran out of token budget.
	//
	// Any of "end_of_turn", "end_of_message", "out_of_tokens".
	StopReason CompletionMessageStopReason `json:"stop_reason,required"`
	// List of tool calls. Each tool call is a ToolCall object.
	ToolCalls []ToolCall `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		StopReason  respjson.Field
		ToolCalls   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionMessage) RawJSON() string { return r.JSON.raw }
func (r *CompletionMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CompletionMessage to a CompletionMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CompletionMessageParam.Overrides()
func (r CompletionMessage) ToParam() CompletionMessageParam {
	return param.Override[CompletionMessageParam](json.RawMessage(r.RawJSON()))
}

// Reason why the model stopped generating. Options are: -
// `StopReason.end_of_turn`: The model finished generating the entire response. -
// `StopReason.end_of_message`: The model finished generating but generated a
// partial response -- usually, a tool call. The user may call the tool and
// continue the conversation with the tool's response. -
// `StopReason.out_of_tokens`: The model ran out of token budget.
type CompletionMessageStopReason string

const (
	CompletionMessageStopReasonEndOfTurn    CompletionMessageStopReason = "end_of_turn"
	CompletionMessageStopReasonEndOfMessage CompletionMessageStopReason = "end_of_message"
	CompletionMessageStopReasonOutOfTokens  CompletionMessageStopReason = "out_of_tokens"
)

// A message containing the model's (assistant) response in a chat conversation.
//
// The properties Content, Role, StopReason are required.
type CompletionMessageParam struct {
	// The content of the model's response
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Reason why the model stopped generating. Options are: -
	// `StopReason.end_of_turn`: The model finished generating the entire response. -
	// `StopReason.end_of_message`: The model finished generating but generated a
	// partial response -- usually, a tool call. The user may call the tool and
	// continue the conversation with the tool's response. -
	// `StopReason.out_of_tokens`: The model ran out of token budget.
	//
	// Any of "end_of_turn", "end_of_message", "out_of_tokens".
	StopReason CompletionMessageStopReason `json:"stop_reason,omitzero,required"`
	// List of tool calls. Each tool call is a ToolCall object.
	ToolCalls []ToolCallParam `json:"tool_calls,omitzero"`
	// Must be "assistant" to identify this as the model's response
	//
	// This field can be elided, and will marshal its zero value as "assistant".
	Role constant.Assistant `json:"role,required"`
	paramObj
}

func (r CompletionMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow CompletionMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ContentDeltaUnion contains all possible properties and values from
// [ContentDeltaText], [ContentDeltaImage], [ContentDeltaToolCall].
//
// Use the [ContentDeltaUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ContentDeltaUnion struct {
	// This field is from variant [ContentDeltaText].
	Text string `json:"text"`
	// Any of "text", "image", "tool_call".
	Type string `json:"type"`
	// This field is from variant [ContentDeltaImage].
	Image string `json:"image"`
	// This field is from variant [ContentDeltaToolCall].
	ParseStatus string `json:"parse_status"`
	// This field is from variant [ContentDeltaToolCall].
	ToolCall ContentDeltaToolCallToolCallUnion `json:"tool_call"`
	JSON     struct {
		Text        respjson.Field
		Type        respjson.Field
		Image       respjson.Field
		ParseStatus respjson.Field
		ToolCall    respjson.Field
		raw         string
	} `json:"-"`
}

// anyContentDelta is implemented by each variant of [ContentDeltaUnion] to add
// type safety for the return type of [ContentDeltaUnion.AsAny]
type anyContentDelta interface {
	implContentDeltaUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ContentDeltaUnion.AsAny().(type) {
//	case llamastackclient.ContentDeltaText:
//	case llamastackclient.ContentDeltaImage:
//	case llamastackclient.ContentDeltaToolCall:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ContentDeltaUnion) AsAny() anyContentDelta {
	switch u.Type {
	case "text":
		return u.AsText()
	case "image":
		return u.AsImage()
	case "tool_call":
		return u.AsToolCall()
	}
	return nil
}

func (u ContentDeltaUnion) AsText() (v ContentDeltaText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentDeltaUnion) AsImage() (v ContentDeltaImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentDeltaUnion) AsToolCall() (v ContentDeltaToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContentDeltaUnion) RawJSON() string { return u.JSON.raw }

func (r *ContentDeltaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content delta for streaming responses.
type ContentDeltaText struct {
	// The incremental text content
	Text string `json:"text,required"`
	// Discriminator type of the delta. Always "text"
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentDeltaText) RawJSON() string { return r.JSON.raw }
func (r *ContentDeltaText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ContentDeltaText) implContentDeltaUnion() {}

// An image content delta for streaming responses.
type ContentDeltaImage struct {
	// The incremental image data as bytes
	Image string `json:"image,required"`
	// Discriminator type of the delta. Always "image"
	Type constant.Image `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentDeltaImage) RawJSON() string { return r.JSON.raw }
func (r *ContentDeltaImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ContentDeltaImage) implContentDeltaUnion() {}

// A tool call content delta for streaming responses.
type ContentDeltaToolCall struct {
	// Current parsing status of the tool call
	//
	// Any of "started", "in_progress", "failed", "succeeded".
	ParseStatus string `json:"parse_status,required"`
	// Either an in-progress tool call string or the final parsed tool call
	ToolCall ContentDeltaToolCallToolCallUnion `json:"tool_call,required"`
	// Discriminator type of the delta. Always "tool_call"
	Type constant.ToolCall `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ParseStatus respjson.Field
		ToolCall    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentDeltaToolCall) RawJSON() string { return r.JSON.raw }
func (r *ContentDeltaToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ContentDeltaToolCall) implContentDeltaUnion() {}

// ContentDeltaToolCallToolCallUnion contains all possible properties and values
// from [string], [ToolCall].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type ContentDeltaToolCallToolCallUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [ToolCall].
	Arguments ToolCallArgumentsUnion `json:"arguments"`
	// This field is from variant [ToolCall].
	CallID string `json:"call_id"`
	// This field is from variant [ToolCall].
	ToolName ToolCallToolName `json:"tool_name"`
	// This field is from variant [ToolCall].
	ArgumentsJson string `json:"arguments_json"`
	JSON          struct {
		OfString      respjson.Field
		Arguments     respjson.Field
		CallID        respjson.Field
		ToolName      respjson.Field
		ArgumentsJson respjson.Field
		raw           string
	} `json:"-"`
}

func (u ContentDeltaToolCallToolCallUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ContentDeltaToolCallToolCallUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ContentDeltaToolCallToolCallUnion) RawJSON() string { return u.JSON.raw }

func (r *ContentDeltaToolCallToolCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document to be used for document ingestion in the RAG Tool.
//
// The properties Content, DocumentID, Metadata are required.
type DocumentParam struct {
	// The content of the document.
	Content DocumentContentUnionParam `json:"content,omitzero,required"`
	// The unique identifier for the document.
	DocumentID string `json:"document_id,required"`
	// Additional metadata for the document.
	Metadata map[string]DocumentMetadataUnionParam `json:"metadata,omitzero,required"`
	// The MIME type of the document.
	MimeType param.Opt[string] `json:"mime_type,omitzero"`
	paramObj
}

func (r DocumentParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DocumentContentUnionParam struct {
	OfString                      param.Opt[string]                     `json:",omitzero,inline"`
	OfImageContentItem            *DocumentContentImageContentItemParam `json:",omitzero,inline"`
	OfTextContentItem             *DocumentContentTextContentItemParam  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam    `json:",omitzero,inline"`
	OfURL                         *DocumentContentURLParam              `json:",omitzero,inline"`
	paramUnion
}

func (u DocumentContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfImageContentItem,
		u.OfTextContentItem,
		u.OfInterleavedContentItemArray,
		u.OfURL)
}
func (u *DocumentContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DocumentContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItem) {
		return u.OfImageContentItem
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	} else if !param.IsOmitted(u.OfURL) {
		return u.OfURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DocumentContentUnionParam) GetImage() *DocumentContentImageContentItemImageParam {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DocumentContentUnionParam) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DocumentContentUnionParam) GetUri() *string {
	if vt := u.OfURL; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u DocumentContentUnionParam) GetType() *string {
	if vt := u.OfImageContentItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The properties Image, Type are required.
type DocumentContentImageContentItemParam struct {
	// Image as a base64 encoded string or an URL
	Image DocumentContentImageContentItemImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r DocumentContentImageContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentImageContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentImageContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type DocumentContentImageContentItemImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL DocumentContentImageContentItemImageURLParam `json:"url,omitzero"`
	paramObj
}

func (r DocumentContentImageContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentImageContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentImageContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
//
// The property Uri is required.
type DocumentContentImageContentItemImageURLParam struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r DocumentContentImageContentItemImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentImageContentItemImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentImageContentItemImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type DocumentContentTextContentItemParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r DocumentContentTextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentTextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentTextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type DocumentContentURLParam struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r DocumentContentURLParam) MarshalJSON() (data []byte, err error) {
	type shadow DocumentContentURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DocumentContentURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DocumentMetadataUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u DocumentMetadataUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *DocumentMetadataUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DocumentMetadataUnionParam) asAny() any {
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

// InterleavedContentUnion contains all possible properties and values from
// [string], [InterleavedContentImageContentItem],
// [InterleavedContentTextContentItem], [[]InterleavedContentItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInterleavedContentItemArray]
type InterleavedContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]InterleavedContentItemUnion]
	// instead of an object.
	OfInterleavedContentItemArray []InterleavedContentItemUnion `json:",inline"`
	// This field is from variant [InterleavedContentImageContentItem].
	Image InterleavedContentImageContentItemImage `json:"image"`
	Type  string                                  `json:"type"`
	// This field is from variant [InterleavedContentTextContentItem].
	Text string `json:"text"`
	JSON struct {
		OfString                      respjson.Field
		OfInterleavedContentItemArray respjson.Field
		Image                         respjson.Field
		Type                          respjson.Field
		Text                          respjson.Field
		raw                           string
	} `json:"-"`
}

func (u InterleavedContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsImageContentItem() (v InterleavedContentImageContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsTextContentItem() (v InterleavedContentTextContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentUnion) AsInterleavedContentItemArray() (v []InterleavedContentItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InterleavedContentUnion) RawJSON() string { return u.JSON.raw }

func (r *InterleavedContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InterleavedContentUnion to a InterleavedContentUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InterleavedContentUnionParam.Overrides()
func (r InterleavedContentUnion) ToParam() InterleavedContentUnionParam {
	return param.Override[InterleavedContentUnionParam](json.RawMessage(r.RawJSON()))
}

// A image content item
type InterleavedContentImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentImageContentItemImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	Type constant.Image `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentImageContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentImageContentItemImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL InterleavedContentImageContentItemImageURL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentImageContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
type InterleavedContentImageContentItemImageURL struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentImageContentItemImageURL) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentImageContentItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type InterleavedContentTextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InterleavedContentParamOfImageContentItem(image InterleavedContentImageContentItemImageParam) InterleavedContentUnionParam {
	var variant InterleavedContentImageContentItemParam
	variant.Image = image
	return InterleavedContentUnionParam{OfImageContentItem: &variant}
}

func InterleavedContentParamOfTextContentItem(text string) InterleavedContentUnionParam {
	var variant InterleavedContentTextContentItemParam
	variant.Text = text
	return InterleavedContentUnionParam{OfTextContentItem: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InterleavedContentUnionParam struct {
	OfString                      param.Opt[string]                        `json:",omitzero,inline"`
	OfImageContentItem            *InterleavedContentImageContentItemParam `json:",omitzero,inline"`
	OfTextContentItem             *InterleavedContentTextContentItemParam  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam       `json:",omitzero,inline"`
	paramUnion
}

func (u InterleavedContentUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfImageContentItem, u.OfTextContentItem, u.OfInterleavedContentItemArray)
}
func (u *InterleavedContentUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InterleavedContentUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfImageContentItem) {
		return u.OfImageContentItem
	} else if !param.IsOmitted(u.OfTextContentItem) {
		return u.OfTextContentItem
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetImage() *InterleavedContentImageContentItemImageParam {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentUnionParam) GetType() *string {
	if vt := u.OfImageContentItem; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTextContentItem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// A image content item
//
// The properties Image, Type are required.
type InterleavedContentImageContentItemParam struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentImageContentItemImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r InterleavedContentImageContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentImageContentItemImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL InterleavedContentImageContentItemImageURLParam `json:"url,omitzero"`
	paramObj
}

func (r InterleavedContentImageContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
//
// The property Uri is required.
type InterleavedContentImageContentItemImageURLParam struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r InterleavedContentImageContentItemImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentImageContentItemImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentImageContentItemImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type InterleavedContentTextContentItemParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InterleavedContentTextContentItemParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentTextContentItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentTextContentItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// InterleavedContentItemUnion contains all possible properties and values from
// [InterleavedContentItemImage], [InterleavedContentItemText].
//
// Use the [InterleavedContentItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type InterleavedContentItemUnion struct {
	// This field is from variant [InterleavedContentItemImage].
	Image InterleavedContentItemImageImage `json:"image"`
	// Any of "image", "text".
	Type string `json:"type"`
	// This field is from variant [InterleavedContentItemText].
	Text string `json:"text"`
	JSON struct {
		Image respjson.Field
		Type  respjson.Field
		Text  respjson.Field
		raw   string
	} `json:"-"`
}

// anyInterleavedContentItem is implemented by each variant of
// [InterleavedContentItemUnion] to add type safety for the return type of
// [InterleavedContentItemUnion.AsAny]
type anyInterleavedContentItem interface {
	implInterleavedContentItemUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := InterleavedContentItemUnion.AsAny().(type) {
//	case llamastackclient.InterleavedContentItemImage:
//	case llamastackclient.InterleavedContentItemText:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u InterleavedContentItemUnion) AsAny() anyInterleavedContentItem {
	switch u.Type {
	case "image":
		return u.AsImage()
	case "text":
		return u.AsText()
	}
	return nil
}

func (u InterleavedContentItemUnion) AsImage() (v InterleavedContentItemImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u InterleavedContentItemUnion) AsText() (v InterleavedContentItemText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u InterleavedContentItemUnion) RawJSON() string { return u.JSON.raw }

func (r *InterleavedContentItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this InterleavedContentItemUnion to a
// InterleavedContentItemUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// InterleavedContentItemUnionParam.Overrides()
func (r InterleavedContentItemUnion) ToParam() InterleavedContentItemUnionParam {
	return param.Override[InterleavedContentItemUnionParam](json.RawMessage(r.RawJSON()))
}

// A image content item
type InterleavedContentItemImage struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentItemImageImage `json:"image,required"`
	// Discriminator type of the content item. Always "image"
	Type constant.Image `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (InterleavedContentItemImage) implInterleavedContentItemUnion() {}

// Image as a base64 encoded string or an URL
type InterleavedContentItemImageImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL InterleavedContentItemImageImageURL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImageImage) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImageImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
type InterleavedContentItemImageImageURL struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemImageImageURL) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemImageImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type InterleavedContentItemText struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	Type constant.Text `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InterleavedContentItemText) RawJSON() string { return r.JSON.raw }
func (r *InterleavedContentItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (InterleavedContentItemText) implInterleavedContentItemUnion() {}

func InterleavedContentItemParamOfImage(image InterleavedContentItemImageImageParam) InterleavedContentItemUnionParam {
	var variant InterleavedContentItemImageParam
	variant.Image = image
	return InterleavedContentItemUnionParam{OfImage: &variant}
}

func InterleavedContentItemParamOfText(text string) InterleavedContentItemUnionParam {
	var variant InterleavedContentItemTextParam
	variant.Text = text
	return InterleavedContentItemUnionParam{OfText: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InterleavedContentItemUnionParam struct {
	OfImage *InterleavedContentItemImageParam `json:",omitzero,inline"`
	OfText  *InterleavedContentItemTextParam  `json:",omitzero,inline"`
	paramUnion
}

func (u InterleavedContentItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfImage, u.OfText)
}
func (u *InterleavedContentItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InterleavedContentItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfImage) {
		return u.OfImage
	} else if !param.IsOmitted(u.OfText) {
		return u.OfText
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetImage() *InterleavedContentItemImageImageParam {
	if vt := u.OfImage; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InterleavedContentItemUnionParam) GetType() *string {
	if vt := u.OfImage; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[InterleavedContentItemUnionParam](
		"type",
		apijson.Discriminator[InterleavedContentItemImageParam]("image"),
		apijson.Discriminator[InterleavedContentItemTextParam]("text"),
	)
}

// A image content item
//
// The properties Image, Type are required.
type InterleavedContentItemImageParam struct {
	// Image as a base64 encoded string or an URL
	Image InterleavedContentItemImageImageParam `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r InterleavedContentItemImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type InterleavedContentItemImageImageParam struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL InterleavedContentItemImageImageURLParam `json:"url,omitzero"`
	paramObj
}

func (r InterleavedContentItemImageImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
//
// The property Uri is required.
type InterleavedContentItemImageImageURLParam struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r InterleavedContentItemImageImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemImageImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemImageImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type InterleavedContentItemTextParam struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InterleavedContentItemTextParam) MarshalJSON() (data []byte, err error) {
	type shadow InterleavedContentItemTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InterleavedContentItemTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func MessageParamOfUser[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T) MessageUnionParam {
	var user UserMessageParam
	switch v := any(content).(type) {
	case string:
		user.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		user.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		user.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		user.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfUser: &user}
}

func MessageParamOfSystem[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T) MessageUnionParam {
	var system SystemMessageParam
	switch v := any(content).(type) {
	case string:
		system.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		system.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		system.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		system.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfSystem: &system}
}

func MessageParamOfTool[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](callID string, content T) MessageUnionParam {
	var tool ToolResponseMessageParam
	tool.CallID = callID
	switch v := any(content).(type) {
	case string:
		tool.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		tool.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		tool.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		tool.Content.OfInterleavedContentItemArray = v
	}
	return MessageUnionParam{OfTool: &tool}
}

func MessageParamOfAssistant[
	T string | InterleavedContentImageContentItemParam | InterleavedContentTextContentItemParam | []InterleavedContentItemUnionParam,
](content T, stopReason CompletionMessageStopReason) MessageUnionParam {
	var assistant CompletionMessageParam
	switch v := any(content).(type) {
	case string:
		assistant.Content.OfString = param.NewOpt(v)
	case InterleavedContentImageContentItemParam:
		assistant.Content.OfImageContentItem = &v
	case InterleavedContentTextContentItemParam:
		assistant.Content.OfTextContentItem = &v
	case []InterleavedContentItemUnionParam:
		assistant.Content.OfInterleavedContentItemArray = v
	}
	assistant.StopReason = stopReason
	return MessageUnionParam{OfAssistant: &assistant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MessageUnionParam struct {
	OfUser      *UserMessageParam         `json:",omitzero,inline"`
	OfSystem    *SystemMessageParam       `json:",omitzero,inline"`
	OfTool      *ToolResponseMessageParam `json:",omitzero,inline"`
	OfAssistant *CompletionMessageParam   `json:",omitzero,inline"`
	paramUnion
}

func (u MessageUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser, u.OfSystem, u.OfTool, u.OfAssistant)
}
func (u *MessageUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *MessageUnionParam) asAny() any {
	if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetContext() *InterleavedContentUnionParam {
	if vt := u.OfUser; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.CallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetStopReason() *string {
	if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.StopReason)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetToolCalls() []ToolCallParam {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u MessageUnionParam) GetRole() *string {
	if vt := u.OfUser; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's Content property, if present.
func (u MessageUnionParam) GetContent() *InterleavedContentUnionParam {
	if vt := u.OfUser; vt != nil {
		return &vt.Content
	} else if vt := u.OfSystem; vt != nil {
		return &vt.Content
	} else if vt := u.OfTool; vt != nil {
		return &vt.Content
	} else if vt := u.OfAssistant; vt != nil {
		return &vt.Content
	}
	return nil
}

func init() {
	apijson.RegisterUnion[MessageUnionParam](
		"role",
		apijson.Discriminator[UserMessageParam]("user"),
		apijson.Discriminator[SystemMessageParam]("system"),
		apijson.Discriminator[ToolResponseMessageParam]("tool"),
		apijson.Discriminator[CompletionMessageParam]("assistant"),
	)
}

// Configuration for the RAG query generation.
//
// The properties ChunkTemplate, MaxChunks, MaxTokensInContext,
// QueryGeneratorConfig are required.
type QueryConfigParam struct {
	// Template for formatting each retrieved chunk in the context. Available
	// placeholders: {index} (1-based chunk ordinal), {chunk.content} (chunk content
	// string), {metadata} (chunk metadata dict). Default: "Result {index}\nContent:
	// {chunk.content}\nMetadata: {metadata}\n"
	ChunkTemplate string `json:"chunk_template,required"`
	// Maximum number of chunks to retrieve.
	MaxChunks int64 `json:"max_chunks,required"`
	// Maximum number of tokens in the context.
	MaxTokensInContext int64 `json:"max_tokens_in_context,required"`
	// Configuration for the query generator.
	QueryGeneratorConfig QueryGeneratorConfigUnionParam `json:"query_generator_config,omitzero,required"`
	// Search mode for retrieval—either "vector", "keyword", or "hybrid". Default
	// "vector".
	//
	// Any of "vector", "keyword", "hybrid".
	Mode QueryConfigMode `json:"mode,omitzero"`
	// Configuration for the ranker to use in hybrid search. Defaults to RRF ranker.
	Ranker QueryConfigRankerUnionParam `json:"ranker,omitzero"`
	paramObj
}

func (r QueryConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search mode for retrieval—either "vector", "keyword", or "hybrid". Default
// "vector".
type QueryConfigMode string

const (
	QueryConfigModeVector  QueryConfigMode = "vector"
	QueryConfigModeKeyword QueryConfigMode = "keyword"
	QueryConfigModeHybrid  QueryConfigMode = "hybrid"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryConfigRankerUnionParam struct {
	OfRrf      *QueryConfigRankerRrfParam      `json:",omitzero,inline"`
	OfWeighted *QueryConfigRankerWeightedParam `json:",omitzero,inline"`
	paramUnion
}

func (u QueryConfigRankerUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfRrf, u.OfWeighted)
}
func (u *QueryConfigRankerUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryConfigRankerUnionParam) asAny() any {
	if !param.IsOmitted(u.OfRrf) {
		return u.OfRrf
	} else if !param.IsOmitted(u.OfWeighted) {
		return u.OfWeighted
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigRankerUnionParam) GetImpactFactor() *float64 {
	if vt := u.OfRrf; vt != nil {
		return &vt.ImpactFactor
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigRankerUnionParam) GetAlpha() *float64 {
	if vt := u.OfWeighted; vt != nil {
		return &vt.Alpha
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryConfigRankerUnionParam) GetType() *string {
	if vt := u.OfRrf; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfWeighted; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[QueryConfigRankerUnionParam](
		"type",
		apijson.Discriminator[QueryConfigRankerRrfParam]("rrf"),
		apijson.Discriminator[QueryConfigRankerWeightedParam]("weighted"),
	)
}

// Reciprocal Rank Fusion (RRF) ranker configuration.
//
// The properties ImpactFactor, Type are required.
type QueryConfigRankerRrfParam struct {
	// The impact factor for RRF scoring. Higher values give more weight to
	// higher-ranked results. Must be greater than 0
	ImpactFactor float64 `json:"impact_factor,required"`
	// The type of ranker, always "rrf"
	//
	// This field can be elided, and will marshal its zero value as "rrf".
	Type constant.Rrf `json:"type,required"`
	paramObj
}

func (r QueryConfigRankerRrfParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConfigRankerRrfParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConfigRankerRrfParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Weighted ranker configuration that combines vector and keyword scores.
//
// The properties Alpha, Type are required.
type QueryConfigRankerWeightedParam struct {
	// Weight factor between 0 and 1. 0 means only use keyword scores, 1 means only use
	// vector scores, values in between blend both scores.
	Alpha float64 `json:"alpha,required"`
	// The type of ranker, always "weighted"
	//
	// This field can be elided, and will marshal its zero value as "weighted".
	Type constant.Weighted `json:"type,required"`
	paramObj
}

func (r QueryConfigRankerWeightedParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConfigRankerWeightedParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConfigRankerWeightedParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func QueryGeneratorConfigParamOfDefault(separator string) QueryGeneratorConfigUnionParam {
	var default_ QueryGeneratorConfigDefaultParam
	default_.Separator = separator
	return QueryGeneratorConfigUnionParam{OfDefault: &default_}
}

func QueryGeneratorConfigParamOfLlm(model string, template string) QueryGeneratorConfigUnionParam {
	var llm QueryGeneratorConfigLlmParam
	llm.Model = model
	llm.Template = template
	return QueryGeneratorConfigUnionParam{OfLlm: &llm}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryGeneratorConfigUnionParam struct {
	OfDefault *QueryGeneratorConfigDefaultParam `json:",omitzero,inline"`
	OfLlm     *QueryGeneratorConfigLlmParam     `json:",omitzero,inline"`
	paramUnion
}

func (u QueryGeneratorConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfDefault, u.OfLlm)
}
func (u *QueryGeneratorConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryGeneratorConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfDefault) {
		return u.OfDefault
	} else if !param.IsOmitted(u.OfLlm) {
		return u.OfLlm
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryGeneratorConfigUnionParam) GetSeparator() *string {
	if vt := u.OfDefault; vt != nil {
		return &vt.Separator
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryGeneratorConfigUnionParam) GetModel() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Model
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryGeneratorConfigUnionParam) GetTemplate() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Template
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u QueryGeneratorConfigUnionParam) GetType() *string {
	if vt := u.OfDefault; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfLlm; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[QueryGeneratorConfigUnionParam](
		"type",
		apijson.Discriminator[QueryGeneratorConfigDefaultParam]("default"),
		apijson.Discriminator[QueryGeneratorConfigLlmParam]("llm"),
	)
}

// Configuration for the default RAG query generator.
//
// The properties Separator, Type are required.
type QueryGeneratorConfigDefaultParam struct {
	// String separator used to join query terms
	Separator string `json:"separator,required"`
	// Type of query generator, always 'default'
	//
	// This field can be elided, and will marshal its zero value as "default".
	Type constant.Default `json:"type,required"`
	paramObj
}

func (r QueryGeneratorConfigDefaultParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryGeneratorConfigDefaultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryGeneratorConfigDefaultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the LLM-based RAG query generator.
//
// The properties Model, Template, Type are required.
type QueryGeneratorConfigLlmParam struct {
	// Name of the language model to use for query generation
	Model string `json:"model,required"`
	// Template string for formatting the query generation prompt
	Template string `json:"template,required"`
	// Type of query generator, always 'llm'
	//
	// This field can be elided, and will marshal its zero value as "llm".
	Type constant.Llm `json:"type,required"`
	paramObj
}

func (r QueryGeneratorConfigLlmParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryGeneratorConfigLlmParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryGeneratorConfigLlmParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of a RAG query containing retrieved content and metadata.
type QueryResult struct {
	// Additional metadata about the query result
	Metadata map[string]QueryResultMetadataUnion `json:"metadata,required"`
	// (Optional) The retrieved content from the query
	Content InterleavedContentUnion `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryResult) RawJSON() string { return r.JSON.raw }
func (r *QueryResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QueryResultMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type QueryResultMetadataUnion struct {
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

func (u QueryResultMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryResultMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryResultMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryResultMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QueryResultMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *QueryResultMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ResponseFormatUnion contains all possible properties and values from
// [ResponseFormatJsonSchema], [ResponseFormatGrammar].
//
// Use the [ResponseFormatUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ResponseFormatUnion struct {
	// This field is from variant [ResponseFormatJsonSchema].
	JsonSchema map[string]ResponseFormatJsonSchemaJsonSchemaUnion `json:"json_schema"`
	// Any of "json_schema", "grammar".
	Type string `json:"type"`
	// This field is from variant [ResponseFormatGrammar].
	Bnf  map[string]ResponseFormatGrammarBnfUnion `json:"bnf"`
	JSON struct {
		JsonSchema respjson.Field
		Type       respjson.Field
		Bnf        respjson.Field
		raw        string
	} `json:"-"`
}

// anyResponseFormat is implemented by each variant of [ResponseFormatUnion] to add
// type safety for the return type of [ResponseFormatUnion.AsAny]
type anyResponseFormat interface {
	implResponseFormatUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ResponseFormatUnion.AsAny().(type) {
//	case llamastackclient.ResponseFormatJsonSchema:
//	case llamastackclient.ResponseFormatGrammar:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ResponseFormatUnion) AsAny() anyResponseFormat {
	switch u.Type {
	case "json_schema":
		return u.AsJsonSchema()
	case "grammar":
		return u.AsGrammar()
	}
	return nil
}

func (u ResponseFormatUnion) AsJsonSchema() (v ResponseFormatJsonSchema) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatUnion) AsGrammar() (v ResponseFormatGrammar) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseFormatUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ResponseFormatUnion to a ResponseFormatUnionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ResponseFormatUnionParam.Overrides()
func (r ResponseFormatUnion) ToParam() ResponseFormatUnionParam {
	return param.Override[ResponseFormatUnionParam](json.RawMessage(r.RawJSON()))
}

// Configuration for JSON schema-guided response generation.
type ResponseFormatJsonSchema struct {
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model.
	JsonSchema map[string]ResponseFormatJsonSchemaJsonSchemaUnion `json:"json_schema,required"`
	// Must be "json_schema" to identify this format type
	Type constant.JsonSchema `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JsonSchema  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseFormatJsonSchema) RawJSON() string { return r.JSON.raw }
func (r *ResponseFormatJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ResponseFormatJsonSchema) implResponseFormatUnion() {}

// ResponseFormatJsonSchemaJsonSchemaUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseFormatJsonSchemaJsonSchemaUnion struct {
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

func (u ResponseFormatJsonSchemaJsonSchemaUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatJsonSchemaJsonSchemaUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatJsonSchemaJsonSchemaUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatJsonSchemaJsonSchemaUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseFormatJsonSchemaJsonSchemaUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseFormatJsonSchemaJsonSchemaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for grammar-guided response generation.
type ResponseFormatGrammar struct {
	// The BNF grammar specification the response should conform to
	Bnf map[string]ResponseFormatGrammarBnfUnion `json:"bnf,required"`
	// Must be "grammar" to identify this format type
	Type constant.Grammar `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bnf         respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResponseFormatGrammar) RawJSON() string { return r.JSON.raw }
func (r *ResponseFormatGrammar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ResponseFormatGrammar) implResponseFormatUnion() {}

// ResponseFormatGrammarBnfUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ResponseFormatGrammarBnfUnion struct {
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

func (u ResponseFormatGrammarBnfUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatGrammarBnfUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatGrammarBnfUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ResponseFormatGrammarBnfUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ResponseFormatGrammarBnfUnion) RawJSON() string { return u.JSON.raw }

func (r *ResponseFormatGrammarBnfUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ResponseFormatParamOfJsonSchema(jsonSchema map[string]ResponseFormatJsonSchemaJsonSchemaUnionParam) ResponseFormatUnionParam {
	var variant ResponseFormatJsonSchemaParam
	variant.JsonSchema = jsonSchema
	return ResponseFormatUnionParam{OfJsonSchema: &variant}
}

func ResponseFormatParamOfGrammar(bnf map[string]ResponseFormatGrammarBnfUnionParam) ResponseFormatUnionParam {
	var grammar ResponseFormatGrammarParam
	grammar.Bnf = bnf
	return ResponseFormatUnionParam{OfGrammar: &grammar}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseFormatUnionParam struct {
	OfJsonSchema *ResponseFormatJsonSchemaParam `json:",omitzero,inline"`
	OfGrammar    *ResponseFormatGrammarParam    `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseFormatUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfJsonSchema, u.OfGrammar)
}
func (u *ResponseFormatUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseFormatUnionParam) asAny() any {
	if !param.IsOmitted(u.OfJsonSchema) {
		return u.OfJsonSchema
	} else if !param.IsOmitted(u.OfGrammar) {
		return u.OfGrammar
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetJsonSchema() map[string]ResponseFormatJsonSchemaJsonSchemaUnionParam {
	if vt := u.OfJsonSchema; vt != nil {
		return vt.JsonSchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetBnf() map[string]ResponseFormatGrammarBnfUnionParam {
	if vt := u.OfGrammar; vt != nil {
		return vt.Bnf
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ResponseFormatUnionParam) GetType() *string {
	if vt := u.OfJsonSchema; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfGrammar; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ResponseFormatUnionParam](
		"type",
		apijson.Discriminator[ResponseFormatJsonSchemaParam]("json_schema"),
		apijson.Discriminator[ResponseFormatGrammarParam]("grammar"),
	)
}

// Configuration for JSON schema-guided response generation.
//
// The properties JsonSchema, Type are required.
type ResponseFormatJsonSchemaParam struct {
	// The JSON schema the response should conform to. In a Python SDK, this is often a
	// `pydantic` model.
	JsonSchema map[string]ResponseFormatJsonSchemaJsonSchemaUnionParam `json:"json_schema,omitzero,required"`
	// Must be "json_schema" to identify this format type
	//
	// This field can be elided, and will marshal its zero value as "json_schema".
	Type constant.JsonSchema `json:"type,required"`
	paramObj
}

func (r ResponseFormatJsonSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseFormatJsonSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseFormatJsonSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseFormatJsonSchemaJsonSchemaUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseFormatJsonSchemaJsonSchemaUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseFormatJsonSchemaJsonSchemaUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseFormatJsonSchemaJsonSchemaUnionParam) asAny() any {
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

// Configuration for grammar-guided response generation.
//
// The properties Bnf, Type are required.
type ResponseFormatGrammarParam struct {
	// The BNF grammar specification the response should conform to
	Bnf map[string]ResponseFormatGrammarBnfUnionParam `json:"bnf,omitzero,required"`
	// Must be "grammar" to identify this format type
	//
	// This field can be elided, and will marshal its zero value as "grammar".
	Type constant.Grammar `json:"type,required"`
	paramObj
}

func (r ResponseFormatGrammarParam) MarshalJSON() (data []byte, err error) {
	type shadow ResponseFormatGrammarParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResponseFormatGrammarParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ResponseFormatGrammarBnfUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u ResponseFormatGrammarBnfUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *ResponseFormatGrammarBnfUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ResponseFormatGrammarBnfUnionParam) asAny() any {
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

// Details of a safety violation detected by content moderation.
type SafetyViolation struct {
	// Additional metadata including specific violation codes for debugging and
	// telemetry
	Metadata map[string]SafetyViolationMetadataUnion `json:"metadata,required"`
	// Severity level of the violation
	//
	// Any of "info", "warn", "error".
	ViolationLevel SafetyViolationViolationLevel `json:"violation_level,required"`
	// (Optional) Message to convey to the user about the violation
	UserMessage string `json:"user_message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata       respjson.Field
		ViolationLevel respjson.Field
		UserMessage    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SafetyViolation) RawJSON() string { return r.JSON.raw }
func (r *SafetyViolation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SafetyViolationMetadataUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type SafetyViolationMetadataUnion struct {
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

func (u SafetyViolationMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SafetyViolationMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SafetyViolationMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *SafetyViolationMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Severity level of the violation
type SafetyViolationViolationLevel string

const (
	SafetyViolationViolationLevelInfo  SafetyViolationViolationLevel = "info"
	SafetyViolationViolationLevelWarn  SafetyViolationViolationLevel = "warn"
	SafetyViolationViolationLevelError SafetyViolationViolationLevel = "error"
)

// Sampling parameters.
type SamplingParamsResp struct {
	// The sampling strategy.
	Strategy SamplingParamsStrategyUnionResp `json:"strategy,required"`
	// The maximum number of tokens that can be generated in the completion. The token
	// count of your prompt plus max_tokens cannot exceed the model's context length.
	MaxTokens int64 `json:"max_tokens"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	RepetitionPenalty float64 `json:"repetition_penalty"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop []string `json:"stop"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Strategy          respjson.Field
		MaxTokens         respjson.Field
		RepetitionPenalty respjson.Field
		Stop              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SamplingParamsResp) RawJSON() string { return r.JSON.raw }
func (r *SamplingParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SamplingParamsResp to a SamplingParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SamplingParams.Overrides()
func (r SamplingParamsResp) ToParam() SamplingParams {
	return param.Override[SamplingParams](json.RawMessage(r.RawJSON()))
}

// SamplingParamsStrategyUnionResp contains all possible properties and values from
// [SamplingParamsStrategyGreedyResp], [SamplingParamsStrategyTopPResp],
// [SamplingParamsStrategyTopKResp].
//
// Use the [SamplingParamsStrategyUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SamplingParamsStrategyUnionResp struct {
	// Any of "greedy", "top_p", "top_k".
	Type string `json:"type"`
	// This field is from variant [SamplingParamsStrategyTopPResp].
	Temperature float64 `json:"temperature"`
	// This field is from variant [SamplingParamsStrategyTopPResp].
	TopP float64 `json:"top_p"`
	// This field is from variant [SamplingParamsStrategyTopKResp].
	TopK int64 `json:"top_k"`
	JSON struct {
		Type        respjson.Field
		Temperature respjson.Field
		TopP        respjson.Field
		TopK        respjson.Field
		raw         string
	} `json:"-"`
}

// anySamplingParamsStrategyResp is implemented by each variant of
// [SamplingParamsStrategyUnionResp] to add type safety for the return type of
// [SamplingParamsStrategyUnionResp.AsAny]
type anySamplingParamsStrategyResp interface {
	implSamplingParamsStrategyUnionResp()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := SamplingParamsStrategyUnionResp.AsAny().(type) {
//	case llamastackclient.SamplingParamsStrategyGreedyResp:
//	case llamastackclient.SamplingParamsStrategyTopPResp:
//	case llamastackclient.SamplingParamsStrategyTopKResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u SamplingParamsStrategyUnionResp) AsAny() anySamplingParamsStrategyResp {
	switch u.Type {
	case "greedy":
		return u.AsGreedy()
	case "top_p":
		return u.AsTopP()
	case "top_k":
		return u.AsTopK()
	}
	return nil
}

func (u SamplingParamsStrategyUnionResp) AsGreedy() (v SamplingParamsStrategyGreedyResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SamplingParamsStrategyUnionResp) AsTopP() (v SamplingParamsStrategyTopPResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SamplingParamsStrategyUnionResp) AsTopK() (v SamplingParamsStrategyTopKResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SamplingParamsStrategyUnionResp) RawJSON() string { return u.JSON.raw }

func (r *SamplingParamsStrategyUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Greedy sampling strategy that selects the highest probability token at each
// step.
type SamplingParamsStrategyGreedyResp struct {
	// Must be "greedy" to identify this sampling strategy
	Type constant.Greedy `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SamplingParamsStrategyGreedyResp) RawJSON() string { return r.JSON.raw }
func (r *SamplingParamsStrategyGreedyResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (SamplingParamsStrategyGreedyResp) implSamplingParamsStrategyUnionResp() {}

// Top-p (nucleus) sampling strategy that samples from the smallest set of tokens
// with cumulative probability >= p.
type SamplingParamsStrategyTopPResp struct {
	// Must be "top_p" to identify this sampling strategy
	Type constant.TopP `json:"type,required"`
	// Controls randomness in sampling. Higher values increase randomness
	Temperature float64 `json:"temperature"`
	// Cumulative probability threshold for nucleus sampling. Defaults to 0.95
	TopP float64 `json:"top_p"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Temperature respjson.Field
		TopP        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SamplingParamsStrategyTopPResp) RawJSON() string { return r.JSON.raw }
func (r *SamplingParamsStrategyTopPResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (SamplingParamsStrategyTopPResp) implSamplingParamsStrategyUnionResp() {}

// Top-k sampling strategy that restricts sampling to the k most likely tokens.
type SamplingParamsStrategyTopKResp struct {
	// Number of top tokens to consider for sampling. Must be at least 1
	TopK int64 `json:"top_k,required"`
	// Must be "top_k" to identify this sampling strategy
	Type constant.TopK `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TopK        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SamplingParamsStrategyTopKResp) RawJSON() string { return r.JSON.raw }
func (r *SamplingParamsStrategyTopKResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (SamplingParamsStrategyTopKResp) implSamplingParamsStrategyUnionResp() {}

// Sampling parameters.
//
// The property Strategy is required.
type SamplingParams struct {
	// The sampling strategy.
	Strategy SamplingParamsStrategyUnion `json:"strategy,omitzero,required"`
	// The maximum number of tokens that can be generated in the completion. The token
	// count of your prompt plus max_tokens cannot exceed the model's context length.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// Number between -2.0 and 2.0. Positive values penalize new tokens based on
	// whether they appear in the text so far, increasing the model's likelihood to
	// talk about new topics.
	RepetitionPenalty param.Opt[float64] `json:"repetition_penalty,omitzero"`
	// Up to 4 sequences where the API will stop generating further tokens. The
	// returned text will not contain the stop sequence.
	Stop []string `json:"stop,omitzero"`
	paramObj
}

func (r SamplingParams) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SamplingParamsStrategyUnion struct {
	OfGreedy *SamplingParamsStrategyGreedy `json:",omitzero,inline"`
	OfTopP   *SamplingParamsStrategyTopP   `json:",omitzero,inline"`
	OfTopK   *SamplingParamsStrategyTopK   `json:",omitzero,inline"`
	paramUnion
}

func (u SamplingParamsStrategyUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGreedy, u.OfTopP, u.OfTopK)
}
func (u *SamplingParamsStrategyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SamplingParamsStrategyUnion) asAny() any {
	if !param.IsOmitted(u.OfGreedy) {
		return u.OfGreedy
	} else if !param.IsOmitted(u.OfTopP) {
		return u.OfTopP
	} else if !param.IsOmitted(u.OfTopK) {
		return u.OfTopK
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTemperature() *float64 {
	if vt := u.OfTopP; vt != nil && vt.Temperature.Valid() {
		return &vt.Temperature.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTopP() *float64 {
	if vt := u.OfTopP; vt != nil && vt.TopP.Valid() {
		return &vt.TopP.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetTopK() *int64 {
	if vt := u.OfTopK; vt != nil {
		return &vt.TopK
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SamplingParamsStrategyUnion) GetType() *string {
	if vt := u.OfGreedy; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTopP; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfTopK; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SamplingParamsStrategyUnion](
		"type",
		apijson.Discriminator[SamplingParamsStrategyGreedy]("greedy"),
		apijson.Discriminator[SamplingParamsStrategyTopP]("top_p"),
		apijson.Discriminator[SamplingParamsStrategyTopK]("top_k"),
	)
}

func NewSamplingParamsStrategyGreedy() SamplingParamsStrategyGreedy {
	return SamplingParamsStrategyGreedy{
		Type: "greedy",
	}
}

// Greedy sampling strategy that selects the highest probability token at each
// step.
//
// This struct has a constant value, construct it with
// [NewSamplingParamsStrategyGreedy].
type SamplingParamsStrategyGreedy struct {
	// Must be "greedy" to identify this sampling strategy
	Type constant.Greedy `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyGreedy) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyGreedy
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyGreedy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-p (nucleus) sampling strategy that samples from the smallest set of tokens
// with cumulative probability >= p.
//
// The property Type is required.
type SamplingParamsStrategyTopP struct {
	// Controls randomness in sampling. Higher values increase randomness
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Cumulative probability threshold for nucleus sampling. Defaults to 0.95
	TopP param.Opt[float64] `json:"top_p,omitzero"`
	// Must be "top_p" to identify this sampling strategy
	//
	// This field can be elided, and will marshal its zero value as "top_p".
	Type constant.TopP `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyTopP) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyTopP
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyTopP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Top-k sampling strategy that restricts sampling to the k most likely tokens.
//
// The properties TopK, Type are required.
type SamplingParamsStrategyTopK struct {
	// Number of top tokens to consider for sampling. Must be at least 1
	TopK int64 `json:"top_k,required"`
	// Must be "top_k" to identify this sampling strategy
	//
	// This field can be elided, and will marshal its zero value as "top_k".
	Type constant.TopK `json:"type,required"`
	paramObj
}

func (r SamplingParamsStrategyTopK) MarshalJSON() (data []byte, err error) {
	type shadow SamplingParamsStrategyTopK
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SamplingParamsStrategyTopK) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A scoring result for a single row.
type ScoringResult struct {
	// Map of metric name to aggregated value
	AggregatedResults map[string]ScoringResultAggregatedResultUnion `json:"aggregated_results,required"`
	// The scoring result for each row. Each row is a map of column name to value.
	ScoreRows []map[string]ScoringResultScoreRowUnion `json:"score_rows,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregatedResults respjson.Field
		ScoreRows         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringResult) RawJSON() string { return r.JSON.raw }
func (r *ScoringResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringResultAggregatedResultUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ScoringResultAggregatedResultUnion struct {
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

func (u ScoringResultAggregatedResultUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultAggregatedResultUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringResultAggregatedResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringResultAggregatedResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringResultScoreRowUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ScoringResultScoreRowUnion struct {
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

func (u ScoringResultScoreRowUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringResultScoreRowUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringResultScoreRowUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringResultScoreRowUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The properties Content, Role are required.
type SystemMessageParam struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Must be "system" to identify this as a system message
	//
	// This field can be elided, and will marshal its zero value as "system".
	Role constant.System `json:"role,required"`
	paramObj
}

func (r SystemMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolCall struct {
	Arguments     ToolCallArgumentsUnion `json:"arguments,required"`
	CallID        string                 `json:"call_id,required"`
	ToolName      ToolCallToolName       `json:"tool_name,required"`
	ArgumentsJson string                 `json:"arguments_json"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments     respjson.Field
		CallID        respjson.Field
		ToolName      respjson.Field
		ArgumentsJson respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolCall) RawJSON() string { return r.JSON.raw }
func (r *ToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolCall to a ToolCallParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolCallParam.Overrides()
func (r ToolCall) ToParam() ToolCallParam {
	return param.Override[ToolCallParam](json.RawMessage(r.RawJSON()))
}

// ToolCallArgumentsUnion contains all possible properties and values from
// [string], [map[string]ToolCallArgumentsMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfToolCallArgumentsMapItemArray]
type ToolCallArgumentsUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ToolCallArgumentsMapItemArrayItemUnion] instead of an object.
	OfToolCallArgumentsMapItemArray []ToolCallArgumentsMapItemArrayItemUnion `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfFloat                         respjson.Field
		OfBool                          respjson.Field
		OfToolCallArgumentsMapItemArray respjson.Field
		raw                             string
	} `json:"-"`
}

func (u ToolCallArgumentsUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsUnion) AsToolCallArgumentsMapMap() (v map[string]ToolCallArgumentsMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolCallArgumentsUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolCallArgumentsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolCallArgumentsMapItemUnion contains all possible properties and values from
// [string], [float64], [bool], [[]ToolCallArgumentsMapItemArrayItemUnion],
// [map[string]ToolCallArgumentsMapItemMapItemUnion].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool OfToolCallArgumentsMapItemArray]
type ToolCallArgumentsMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a
	// [[]ToolCallArgumentsMapItemArrayItemUnion] instead of an object.
	OfToolCallArgumentsMapItemArray []ToolCallArgumentsMapItemArrayItemUnion `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfFloat                         respjson.Field
		OfBool                          respjson.Field
		OfToolCallArgumentsMapItemArray respjson.Field
		raw                             string
	} `json:"-"`
}

func (u ToolCallArgumentsMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemUnion) AsToolCallArgumentsMapItemArray() (v []ToolCallArgumentsMapItemArrayItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemUnion) AsToolCallArgumentsMapItemMapMap() (v map[string]ToolCallArgumentsMapItemMapItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolCallArgumentsMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolCallArgumentsMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolCallArgumentsMapItemArrayItemUnion contains all possible properties and
// values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ToolCallArgumentsMapItemArrayItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ToolCallArgumentsMapItemArrayItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemArrayItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemArrayItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolCallArgumentsMapItemArrayItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolCallArgumentsMapItemArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolCallArgumentsMapItemMapItemUnion contains all possible properties and values
// from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ToolCallArgumentsMapItemMapItemUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ToolCallArgumentsMapItemMapItemUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemMapItemUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolCallArgumentsMapItemMapItemUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolCallArgumentsMapItemMapItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolCallArgumentsMapItemMapItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolCallToolName string

const (
	ToolCallToolNameBraveSearch     ToolCallToolName = "brave_search"
	ToolCallToolNameWolframAlpha    ToolCallToolName = "wolfram_alpha"
	ToolCallToolNamePhotogen        ToolCallToolName = "photogen"
	ToolCallToolNameCodeInterpreter ToolCallToolName = "code_interpreter"
)

// The properties Arguments, CallID, ToolName are required.
type ToolCallParam struct {
	Arguments     ToolCallArgumentsUnionParam `json:"arguments,omitzero,required"`
	CallID        string                      `json:"call_id,required"`
	ToolName      ToolCallToolName            `json:"tool_name,omitzero,required"`
	ArgumentsJson param.Opt[string]           `json:"arguments_json,omitzero"`
	paramObj
}

func (r ToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolCallArgumentsUnionParam struct {
	OfString                  param.Opt[string]                             `json:",omitzero,inline"`
	OfToolCallArgumentsMapMap map[string]ToolCallArgumentsMapItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ToolCallArgumentsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfToolCallArgumentsMapMap)
}
func (u *ToolCallArgumentsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolCallArgumentsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfToolCallArgumentsMapMap) {
		return &u.OfToolCallArgumentsMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolCallArgumentsMapItemUnionParam struct {
	OfString                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfFloat                          param.Opt[float64]                                   `json:",omitzero,inline"`
	OfBool                           param.Opt[bool]                                      `json:",omitzero,inline"`
	OfToolCallArgumentsMapItemArray  []ToolCallArgumentsMapItemArrayItemUnionParam        `json:",omitzero,inline"`
	OfToolCallArgumentsMapItemMapMap map[string]ToolCallArgumentsMapItemMapItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ToolCallArgumentsMapItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfFloat,
		u.OfBool,
		u.OfToolCallArgumentsMapItemArray,
		u.OfToolCallArgumentsMapItemMapMap)
}
func (u *ToolCallArgumentsMapItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolCallArgumentsMapItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfToolCallArgumentsMapItemArray) {
		return &u.OfToolCallArgumentsMapItemArray
	} else if !param.IsOmitted(u.OfToolCallArgumentsMapItemMapMap) {
		return &u.OfToolCallArgumentsMapItemMapMap
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolCallArgumentsMapItemArrayItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ToolCallArgumentsMapItemArrayItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ToolCallArgumentsMapItemArrayItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolCallArgumentsMapItemArrayItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolCallArgumentsMapItemMapItemUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ToolCallArgumentsMapItemMapItemUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ToolCallArgumentsMapItemMapItemUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolCallArgumentsMapItemMapItemUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// A message representing the result of a tool invocation.
type ToolResponseMessage struct {
	// Unique identifier for the tool call this response is for
	CallID string `json:"call_id,required"`
	// The response content from the tool
	Content InterleavedContentUnion `json:"content,required"`
	// Must be "tool" to identify this as a tool response
	Role constant.Tool `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallID      respjson.Field
		Content     respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolResponseMessage) RawJSON() string { return r.JSON.raw }
func (r *ToolResponseMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ToolResponseMessage to a ToolResponseMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ToolResponseMessageParam.Overrides()
func (r ToolResponseMessage) ToParam() ToolResponseMessageParam {
	return param.Override[ToolResponseMessageParam](json.RawMessage(r.RawJSON()))
}

// A message representing the result of a tool invocation.
//
// The properties CallID, Content, Role are required.
type ToolResponseMessageParam struct {
	// Unique identifier for the tool call this response is for
	CallID string `json:"call_id,required"`
	// The response content from the tool
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// Must be "tool" to identify this as a tool response
	//
	// This field can be elided, and will marshal its zero value as "tool".
	Role constant.Tool `json:"role,required"`
	paramObj
}

func (r ToolResponseMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolResponseMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolResponseMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the user in a chat conversation.
type UserMessage struct {
	// The content of the message, which can include text and other media
	Content InterleavedContentUnion `json:"content,required"`
	// Must be "user" to identify this as a user message
	Role constant.User `json:"role,required"`
	// (Optional) This field is used internally by Llama Stack to pass RAG context.
	// This field may be removed in the API in the future.
	Context InterleavedContentUnion `json:"context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Role        respjson.Field
		Context     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserMessage) RawJSON() string { return r.JSON.raw }
func (r *UserMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserMessage to a UserMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserMessageParam.Overrides()
func (r UserMessage) ToParam() UserMessageParam {
	return param.Override[UserMessageParam](json.RawMessage(r.RawJSON()))
}

// A message from the user in a chat conversation.
//
// The properties Content, Role are required.
type UserMessageParam struct {
	// The content of the message, which can include text and other media
	Content InterleavedContentUnionParam `json:"content,omitzero,required"`
	// (Optional) This field is used internally by Llama Stack to pass RAG context.
	// This field may be removed in the API in the future.
	Context InterleavedContentUnionParam `json:"context,omitzero"`
	// Must be "user" to identify this as a user message
	//
	// This field can be elided, and will marshal its zero value as "user".
	Role constant.User `json:"role,required"`
	paramObj
}

func (r UserMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow UserMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
