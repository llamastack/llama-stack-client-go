// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/packages/ssestream"
	"github.com/llamastack/llama-stack-client-go/shared"
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

// Generate chat completions for a batch of messages using the specified model.
func (r *InferenceService) BatchChatCompletion(ctx context.Context, body InferenceBatchChatCompletionParams, opts ...option.RequestOption) (res *InferenceBatchChatCompletionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/batch-chat-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate completions for a batch of content using the specified model.
func (r *InferenceService) BatchCompletion(ctx context.Context, body InferenceBatchCompletionParams, opts ...option.RequestOption) (res *shared.BatchCompletion, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/batch-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a chat completion for the given messages using the specified model.
//
// Deprecated: /v1/inference/chat-completion is deprecated. Please use
// /v1/openai/v1/chat/completions.
func (r *InferenceService) ChatCompletion(ctx context.Context, body InferenceChatCompletionParams, opts ...option.RequestOption) (res *shared.ChatCompletionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/chat-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a chat completion for the given messages using the specified model.
//
// Deprecated: /v1/inference/chat-completion is deprecated. Please use
// /v1/openai/v1/chat/completions.
func (r *InferenceService) ChatCompletionStreaming(ctx context.Context, body InferenceChatCompletionParams, opts ...option.RequestOption) (stream *ssestream.Stream[ChatCompletionResponseStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/inference/chat-completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[ChatCompletionResponseStreamChunk](ssestream.NewDecoder(raw), err)
}

// Generate a completion for the given content using the specified model.
//
// Deprecated: /v1/inference/completion is deprecated. Please use
// /v1/openai/v1/completions.
func (r *InferenceService) Completion(ctx context.Context, body InferenceCompletionParams, opts ...option.RequestOption) (res *CompletionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Generate a completion for the given content using the specified model.
//
// Deprecated: /v1/inference/completion is deprecated. Please use
// /v1/openai/v1/completions.
func (r *InferenceService) CompletionStreaming(ctx context.Context, body InferenceCompletionParams, opts ...option.RequestOption) (stream *ssestream.Stream[CompletionResponse]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	path := "v1/inference/completion"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &raw, opts...)
	return ssestream.NewStream[CompletionResponse](ssestream.NewDecoder(raw), err)
}

// Generate embeddings for content pieces using the specified model.
//
// Deprecated: /v1/inference/embeddings is deprecated. Please use
// /v1/openai/v1/embeddings.
func (r *InferenceService) Embeddings(ctx context.Context, body InferenceEmbeddingsParams, opts ...option.RequestOption) (res *EmbeddingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/inference/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// A chunk of a streamed chat completion response.
type ChatCompletionResponseStreamChunk struct {
	// The event containing the new content
	Event ChatCompletionResponseStreamChunkEvent `json:"event,required"`
	// (Optional) List of metrics associated with the API response
	Metrics []ChatCompletionResponseStreamChunkMetric `json:"metrics"`
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
	Delta shared.ContentDeltaUnion `json:"delta,required"`
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

// A metric value included in API responses.
type ChatCompletionResponseStreamChunkMetric struct {
	// The name of the metric
	Metric string `json:"metric,required"`
	// The numeric value of the metric
	Value float64 `json:"value,required"`
	// (Optional) The unit of measurement for the metric value
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metric      respjson.Field
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponseStreamChunkMetric) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponseStreamChunkMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from a completion request.
type CompletionResponse struct {
	// The generated completion text
	Content string `json:"content,required"`
	// Reason why generation stopped
	//
	// Any of "end_of_turn", "end_of_message", "out_of_tokens".
	StopReason CompletionResponseStopReason `json:"stop_reason,required"`
	// Optional log probabilities for generated tokens
	Logprobs []TokenLogProbs `json:"logprobs"`
	// (Optional) List of metrics associated with the API response
	Metrics []CompletionResponseMetric `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		StopReason  respjson.Field
		Logprobs    respjson.Field
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *CompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reason why generation stopped
type CompletionResponseStopReason string

const (
	CompletionResponseStopReasonEndOfTurn    CompletionResponseStopReason = "end_of_turn"
	CompletionResponseStopReasonEndOfMessage CompletionResponseStopReason = "end_of_message"
	CompletionResponseStopReasonOutOfTokens  CompletionResponseStopReason = "out_of_tokens"
)

// A metric value included in API responses.
type CompletionResponseMetric struct {
	// The name of the metric
	Metric string `json:"metric,required"`
	// The numeric value of the metric
	Value float64 `json:"value,required"`
	// (Optional) The unit of measurement for the metric value
	Unit string `json:"unit"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metric      respjson.Field
		Value       respjson.Field
		Unit        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionResponseMetric) RawJSON() string { return r.JSON.raw }
func (r *CompletionResponseMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing generated embeddings.
type EmbeddingsResponse struct {
	// List of embedding vectors, one per input content. Each embedding is a list of
	// floats. The dimensionality of the embedding is model-specific; you can check
	// model metadata using /models/{model_id}
	Embeddings [][]float64 `json:"embeddings,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embeddings  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddingsResponse) RawJSON() string { return r.JSON.raw }
func (r *EmbeddingsResponse) UnmarshalJSON(data []byte) error {
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

// Response from a batch chat completion request.
type InferenceBatchChatCompletionResponse struct {
	// List of chat completion responses, one for each conversation in the batch
	Batch []shared.ChatCompletionResponse `json:"batch,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Batch       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceBatchChatCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *InferenceBatchChatCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceBatchChatCompletionParams struct {
	// The messages to generate completions for.
	MessagesBatch [][]shared.MessageUnionParam `json:"messages_batch,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) If specified, log probabilities for each token position will be
	// returned.
	Logprobs InferenceBatchChatCompletionParamsLogprobs `json:"logprobs,omitzero"`
	// (Optional) Grammar specification for guided (structured) decoding.
	ResponseFormat shared.ResponseFormatUnionParam `json:"response_format,omitzero"`
	// (Optional) Parameters to control the sampling strategy.
	SamplingParams shared.SamplingParams `json:"sampling_params,omitzero"`
	// (Optional) Configuration for tool use.
	ToolConfig InferenceBatchChatCompletionParamsToolConfig `json:"tool_config,omitzero"`
	// (Optional) List of tool definitions available to the model.
	Tools []InferenceBatchChatCompletionParamsTool `json:"tools,omitzero"`
	paramObj
}

func (r InferenceBatchChatCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchChatCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchChatCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) If specified, log probabilities for each token position will be
// returned.
type InferenceBatchChatCompletionParamsLogprobs struct {
	// How many tokens (for each position) to return log probabilities for.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r InferenceBatchChatCompletionParamsLogprobs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchChatCompletionParamsLogprobs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchChatCompletionParamsLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Configuration for tool use.
type InferenceBatchChatCompletionParamsToolConfig struct {
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

func (r InferenceBatchChatCompletionParamsToolConfig) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchChatCompletionParamsToolConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchChatCompletionParamsToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[InferenceBatchChatCompletionParamsToolConfig](
		"system_message_behavior", "append", "replace",
	)
	apijson.RegisterFieldValidator[InferenceBatchChatCompletionParamsToolConfig](
		"tool_prompt_format", "json", "function_tag", "python_list",
	)
}

// The property ToolName is required.
type InferenceBatchChatCompletionParamsTool struct {
	ToolName    string                                `json:"tool_name,omitzero,required"`
	Description param.Opt[string]                     `json:"description,omitzero"`
	Parameters  map[string]shared.ToolParamDefinition `json:"parameters,omitzero"`
	paramObj
}

func (r InferenceBatchChatCompletionParamsTool) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchChatCompletionParamsTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchChatCompletionParamsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceBatchCompletionParams struct {
	// The content to generate completions for.
	ContentBatch []shared.InterleavedContentUnionParam `json:"content_batch,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) If specified, log probabilities for each token position will be
	// returned.
	Logprobs InferenceBatchCompletionParamsLogprobs `json:"logprobs,omitzero"`
	// (Optional) Grammar specification for guided (structured) decoding.
	ResponseFormat shared.ResponseFormatUnionParam `json:"response_format,omitzero"`
	// (Optional) Parameters to control the sampling strategy.
	SamplingParams shared.SamplingParams `json:"sampling_params,omitzero"`
	paramObj
}

func (r InferenceBatchCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) If specified, log probabilities for each token position will be
// returned.
type InferenceBatchCompletionParamsLogprobs struct {
	// How many tokens (for each position) to return log probabilities for.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r InferenceBatchCompletionParamsLogprobs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceBatchCompletionParamsLogprobs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceBatchCompletionParamsLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceChatCompletionParams struct {
	// List of messages in the conversation.
	Messages []shared.MessageUnionParam `json:"messages,omitzero,required"`
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
	ResponseFormat shared.ResponseFormatUnionParam `json:"response_format,omitzero"`
	// Parameters to control the sampling strategy.
	SamplingParams shared.SamplingParams `json:"sampling_params,omitzero"`
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
	ToolName    string                                `json:"tool_name,omitzero,required"`
	Description param.Opt[string]                     `json:"description,omitzero"`
	Parameters  map[string]shared.ToolParamDefinition `json:"parameters,omitzero"`
	paramObj
}

func (r InferenceChatCompletionParamsTool) MarshalJSON() (data []byte, err error) {
	type shadow InferenceChatCompletionParamsTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceChatCompletionParamsTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceCompletionParams struct {
	// The content to generate a completion for.
	Content shared.InterleavedContentUnionParam `json:"content,omitzero,required"`
	// The identifier of the model to use. The model must be registered with Llama
	// Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) If specified, log probabilities for each token position will be
	// returned.
	Logprobs InferenceCompletionParamsLogprobs `json:"logprobs,omitzero"`
	// (Optional) Grammar specification for guided (structured) decoding.
	ResponseFormat shared.ResponseFormatUnionParam `json:"response_format,omitzero"`
	// (Optional) Parameters to control the sampling strategy.
	SamplingParams shared.SamplingParams `json:"sampling_params,omitzero"`
	paramObj
}

func (r InferenceCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) If specified, log probabilities for each token position will be
// returned.
type InferenceCompletionParamsLogprobs struct {
	// How many tokens (for each position) to return log probabilities for.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	paramObj
}

func (r InferenceCompletionParamsLogprobs) MarshalJSON() (data []byte, err error) {
	type shadow InferenceCompletionParamsLogprobs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceCompletionParamsLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InferenceEmbeddingsParams struct {
	// List of contents to generate embeddings for. Each content can be a string or an
	// InterleavedContentItem (and hence can be multimodal). The behavior depends on
	// the model and provider. Some models may only support text.
	Contents InferenceEmbeddingsParamsContentsUnion `json:"contents,omitzero,required"`
	// The identifier of the model to use. The model must be an embedding model
	// registered with Llama Stack and available via the /models endpoint.
	ModelID string `json:"model_id,required"`
	// (Optional) Output dimensionality for the embeddings. Only supported by
	// Matryoshka models.
	OutputDimension param.Opt[int64] `json:"output_dimension,omitzero"`
	// (Optional) How is the embedding being used? This is only supported by asymmetric
	// embedding models.
	//
	// Any of "query", "document".
	TaskType InferenceEmbeddingsParamsTaskType `json:"task_type,omitzero"`
	// (Optional) Config for how to truncate text for embedding when text is longer
	// than the model's max sequence length.
	//
	// Any of "none", "start", "end".
	TextTruncation InferenceEmbeddingsParamsTextTruncation `json:"text_truncation,omitzero"`
	paramObj
}

func (r InferenceEmbeddingsParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceEmbeddingsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceEmbeddingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InferenceEmbeddingsParamsContentsUnion struct {
	OfStringArray                 []string                                  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []shared.InterleavedContentItemUnionParam `json:",omitzero,inline"`
	paramUnion
}

func (u InferenceEmbeddingsParamsContentsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringArray, u.OfInterleavedContentItemArray)
}
func (u *InferenceEmbeddingsParamsContentsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InferenceEmbeddingsParamsContentsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfInterleavedContentItemArray) {
		return &u.OfInterleavedContentItemArray
	}
	return nil
}

// (Optional) How is the embedding being used? This is only supported by asymmetric
// embedding models.
type InferenceEmbeddingsParamsTaskType string

const (
	InferenceEmbeddingsParamsTaskTypeQuery    InferenceEmbeddingsParamsTaskType = "query"
	InferenceEmbeddingsParamsTaskTypeDocument InferenceEmbeddingsParamsTaskType = "document"
)

// (Optional) Config for how to truncate text for embedding when text is longer
// than the model's max sequence length.
type InferenceEmbeddingsParamsTextTruncation string

const (
	InferenceEmbeddingsParamsTextTruncationNone  InferenceEmbeddingsParamsTextTruncation = "none"
	InferenceEmbeddingsParamsTextTruncationStart InferenceEmbeddingsParamsTextTruncation = "start"
	InferenceEmbeddingsParamsTextTruncationEnd   InferenceEmbeddingsParamsTextTruncation = "end"
)
