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
	"github.com/llamastack/llama-stack-client-go/shared/constant"
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

// Generate embeddings for content pieces using the specified model.
//
// Deprecated: /v1/inference/embeddings is deprecated. Please use /v1/embeddings.
func (r *InferenceService) Embeddings(ctx context.Context, body InferenceEmbeddingsParams, opts ...option.RequestOption) (res *EmbeddingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/inference/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Rerank a list of documents based on their relevance to a query.
func (r *InferenceService) Rerank(ctx context.Context, body InferenceRerankParams, opts ...option.RequestOption) (res *[]InferenceRerankResponse, err error) {
	var env InferenceRerankResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v1/inference/rerank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
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

// A single rerank result from a reranking response.
type InferenceRerankResponse struct {
	// The original index of the document in the input list
	Index int64 `json:"index,required"`
	// The relevance score from the model output. Values are inverted when applicable
	// so that higher scores indicate greater relevance.
	RelevanceScore float64 `json:"relevance_score,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index          respjson.Field
		RelevanceScore respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceRerankResponse) RawJSON() string { return r.JSON.raw }
func (r *InferenceRerankResponse) UnmarshalJSON(data []byte) error {
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
	OfStringArray                 []string                           `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam `json:",omitzero,inline"`
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

type InferenceRerankParams struct {
	// List of items to rerank. Each item can be a string, text content part, or image
	// content part. Each input must not exceed the model's max input token length.
	Items []InferenceRerankParamsItemUnion `json:"items,omitzero,required"`
	// The identifier of the reranking model to use.
	Model string `json:"model,required"`
	// The search query to rank items against. Can be a string, text content part, or
	// image content part. The input must not exceed the model's max input token
	// length.
	Query InferenceRerankParamsQueryUnion `json:"query,omitzero,required"`
	// (Optional) Maximum number of results to return. Default: returns all.
	MaxNumResults param.Opt[int64] `json:"max_num_results,omitzero"`
	paramObj
}

func (r InferenceRerankParams) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InferenceRerankParamsItemUnion struct {
	OfString                               param.Opt[string]                                                   `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartText  *InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam  `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u InferenceRerankParamsItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *InferenceRerankParamsItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InferenceRerankParamsItemUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsItemUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsItemUnion) GetImageURL() *InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsItemUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsItemOpenAIChatCompletionContentPartTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam struct {
	// Image URL specification and processing details
	ImageURL InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsItemOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InferenceRerankParamsQueryUnion struct {
	OfString                               param.Opt[string]                                                    `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartText  *InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam  `json:",omitzero,inline"`
	OfOpenAIChatCompletionContentPartImage *InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam `json:",omitzero,inline"`
	paramUnion
}

func (u InferenceRerankParamsQueryUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfOpenAIChatCompletionContentPartText, u.OfOpenAIChatCompletionContentPartImage)
}
func (u *InferenceRerankParamsQueryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InferenceRerankParamsQueryUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartText) {
		return u.OfOpenAIChatCompletionContentPartText
	} else if !param.IsOmitted(u.OfOpenAIChatCompletionContentPartImage) {
		return u.OfOpenAIChatCompletionContentPartImage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsQueryUnion) GetText() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsQueryUnion) GetImageURL() *InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL {
	if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u InferenceRerankParamsQueryUnion) GetType() *string {
	if vt := u.OfOpenAIChatCompletionContentPartText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfOpenAIChatCompletionContentPartImage; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsQueryOpenAIChatCompletionContentPartTextParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam struct {
	// Image URL specification and processing details
	ImageURL InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL) MarshalJSON() (data []byte, err error) {
	type shadow InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InferenceRerankParamsQueryOpenAIChatCompletionContentPartImageParamImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response from a reranking request.
type InferenceRerankResponseEnvelope struct {
	// List of rerank result objects, sorted by relevance score (descending)
	Data []InferenceRerankResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InferenceRerankResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *InferenceRerankResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
