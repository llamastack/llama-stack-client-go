// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/packages/ssestream"
)

// InferenceService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInferenceService] method instead.
type InferenceService struct {
	Options []option.RequestOption
}

// NewInferenceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInferenceService(opts ...option.RequestOption) (r InferenceService) {
	r = InferenceService{}
	r.Options = opts
	return
}

// Generate a chat completion for the given messages using the specified model.
//
// Deprecated: /v1/inference/chat-completion is deprecated. Please use
// /v1/chat/completions.
func (r *InferenceService) ChatCompletion(ctx context.Context, body InferenceChatCompletionParams, opts ...option.RequestOption) (res *ChatCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/inference/chat-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a chat completion for the given messages using the specified model.
//
// Deprecated: /v1/inference/chat-completion is deprecated. Please use
// /v1/chat/completions.
func (r *InferenceService) ChatCompletionStreaming(ctx context.Context, body InferenceChatCompletionParams, opts ...option.RequestOption) (stream *ssestream.Stream[ChatCompletionResponseStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/inference/chat-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ChatCompletionResponseStreamChunk](ssestream.NewDecoder(raw), err)
}

// A chunk of a streamed chat completion response.
type ChatCompletionResponseStreamChunk struct {
	// The event containing the new content
	Event ChatCompletionResponseStreamChunkEvent `json:"event,required"`
	// (Optional) List of metrics associated with the API response
	Metrics []Metric `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponseStreamChunk) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponseStreamChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The event containing the new content
type ChatCompletionResponseStreamChunkEvent struct {
	// Content generated since last event. This can be one or more tokens, or a tool
	// call.
	Delta ContentDeltaUnion `json:"delta,required"`
	// Type of the event
	//
	// Any of "start", "complete", "progress".
	EventType string `json:"event_type,required"`
	// Optional log probabilities for generated tokens
	Logprobs []TokenLogProbs `json:"logprobs"`
	// Optional reason why generation stopped, if complete
	//
	// Any of "end_of_turn", "end_of_message", "out_of_tokens".
	StopReason string `json:"stop_reason"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta       respjson.Field
		EventType   respjson.Field
		Logprobs    respjson.Field
		StopReason  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponseStreamChunkEvent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponseStreamChunkEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Log probabilities for generated tokens.
type TokenLogProbs struct {
	// Dictionary mapping tokens to their log probabilities
	LogprobsByToken map[string]float64 `json:"logprobs_by_token,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LogprobsByToken respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TokenLogProbs) RawJSON() string { return r.JSON.raw }
func (r *TokenLogProbs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceChatCompletionParams struct {
	// List of messages in the conversation.
	Messages []MessageUnionParam `json:"messages,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) If specified, log probabilities for each token position will be
	// returned.
	Logprobs InferenceChatCompletionParamsLogprobs `json:"logprobs,omitzero"`
	// (Optional) Grammar specification for guided (structured) decoding. There are two
	// options: - `ResponseFormat.json_schema`: The grammar is a JSON schema. Most
	// providers support this format. - `ResponseFormat.grammar`: The grammar is a BNF
	// grammar. This format is more flexible, but not all providers support it.
	ResponseFormat ResponseFormatUnionParam `json:"response_format,omitzero"`
	// Parameters to control the sampling strategy.
	SamplingParams SamplingParams `json:"sampling_params,omitzero"`
	// (Optional) Whether tool use is required or automatic. Defaults to
	// ToolChoice.auto. .. deprecated:: Use tool_config instead.
	//
	// Any of "auto", "required", "none".
	ToolChoice InferenceChatCompletionParamsToolChoice `json:"tool_choice,omitzero"`
	// (Optional) Configuration for tool use.
	ToolConfig InferenceChatCompletionParamsToolConfig `json:"tool_config,omitzero"`
	// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
	// will attempt to use a format that is best adapted to the model. -
	// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
	// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
	// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
	// are output as Python syntax -- a list of function calls. .. deprecated:: Use
	// tool_config instead.
	//
	// Any of "json", "function_tag", "python_list".
	ToolPromptFormat InferenceChatCompletionParamsToolPromptFormat `json:"tool_prompt_format,omitzero"`
	// (Optional) List of tool definitions available to the model.
	Tools []InferenceChatCompletionParamsTool `json:"tools,omitzero"`
	paramObj
}

func (r InferenceChatCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceChatCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceChatCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) If specified, log probabilities for each token position will be
// returned.
type InferenceChatCompletionParamsLogprobs struct {
	// How many tokens (for each position) to return log probabilities for.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r InferenceChatCompletionParamsLogprobs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceChatCompletionParamsLogprobs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceChatCompletionParamsLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Whether tool use is required or automatic. Defaults to
// ToolChoice.auto. .. deprecated:: Use tool_config instead.
type InferenceChatCompletionParamsToolChoice string

const (
	InferenceChatCompletionParamsToolChoiceAuto     InferenceChatCompletionParamsToolChoice = "auto"
	InferenceChatCompletionParamsToolChoiceRequired InferenceChatCompletionParamsToolChoice = "required"
	InferenceChatCompletionParamsToolChoiceNone     InferenceChatCompletionParamsToolChoice = "none"
)

// (Optional) Configuration for tool use.
type InferenceChatCompletionParamsToolConfig struct {
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

func (r InferenceChatCompletionParamsToolConfig) MarshalJSON() (data []byte, err error) {
	type shadow InferenceChatCompletionParamsToolConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceChatCompletionParamsToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceChatCompletionParamsToolConfig](
		"system_message_behavior", "append", "replace",
	)
	apijson.RegisterFieldValidator[InferenceChatCompletionParamsToolConfig](
		"tool_prompt_format", "json", "function_tag", "python_list",
	)
}

// (Optional) Instructs the model how to format tool calls. By default, Llama Stack
// will attempt to use a format that is best adapted to the model. -
// `ToolPromptFormat.json`: The tool calls are formatted as a JSON object. -
// `ToolPromptFormat.function_tag`: The tool calls are enclosed in a
// <function=function_name> tag. - `ToolPromptFormat.python_list`: The tool calls
// are output as Python syntax -- a list of function calls. .. deprecated:: Use
// tool_config instead.
type InferenceChatCompletionParamsToolPromptFormat string

const (
	InferenceChatCompletionParamsToolPromptFormatJson        InferenceChatCompletionParamsToolPromptFormat = "json"
	InferenceChatCompletionParamsToolPromptFormatFunctionTag InferenceChatCompletionParamsToolPromptFormat = "function_tag"
	InferenceChatCompletionParamsToolPromptFormatPythonList  InferenceChatCompletionParamsToolPromptFormat = "python_list"
)

// The property ToolName is required.
type InferenceChatCompletionParamsTool struct {
	ToolName    string                         `json:"tool_name,omitzero,required"`
	Description param.Opt[string]              `json:"description,omitzero"`
	Parameters  map[string]ToolParamDefinition `json:"parameters,omitzero"`
	paramObj
}

func (r InferenceChatCompletionParamsTool) MarshalJSON() (data []byte, err error) {
	type shadow InferenceChatCompletionParamsTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceChatCompletionParamsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
