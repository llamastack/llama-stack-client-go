// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/packages/ssestream"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// AlphaAgentTurnService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaAgentTurnService] method instead.
type AlphaAgentTurnService struct {
	Options []option.RequestOption
}

// NewAlphaAgentTurnService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaAgentTurnService(opts ...option.RequestOption) (r AlphaAgentTurnService) {
	r = AlphaAgentTurnService{}
	r.Options = opts
	return
}

// Create a new turn for an agent.
func (r *AlphaAgentTurnService) New(ctx context.Context, sessionID string, params AlphaAgentTurnNewParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn", params.AgentID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a new turn for an agent.
func (r *AlphaAgentTurnService) NewStreaming(ctx context.Context, sessionID string, params AlphaAgentTurnNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[AgentTurnResponseStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	if params.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn", params.AgentID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &raw, opts...)
	return ssestream.NewStream[AgentTurnResponseStreamChunk](ssestream.NewDecoder(raw), err)
}

// Retrieve an agent turn by its ID.
func (r *AlphaAgentTurnService) Get(ctx context.Context, turnID string, query AlphaAgentTurnGetParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if query.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if turnID == "" {
		err = errors.New("missing required turn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn/%s", query.AgentID, query.SessionID, turnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Resume an agent turn with executed tool call responses. When a Turn has the
// status `awaiting_input` due to pending input from client side tool calls, this
// endpoint can be used to submit the outputs from the tool calls once they are
// ready.
func (r *AlphaAgentTurnService) Resume(ctx context.Context, turnID string, params AlphaAgentTurnResumeParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if params.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if turnID == "" {
		err = errors.New("missing required turn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn/%s/resume", params.AgentID, params.SessionID, turnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Resume an agent turn with executed tool call responses. When a Turn has the
// status `awaiting_input` due to pending input from client side tool calls, this
// endpoint can be used to submit the outputs from the tool calls once they are
// ready.
func (r *AlphaAgentTurnService) ResumeStreaming(ctx context.Context, turnID string, params AlphaAgentTurnResumeParams, opts ...option.RequestOption) (stream *ssestream.Stream[AgentTurnResponseStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithJSONSet("stream", true)}, opts...)
	if params.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	if params.SessionID == "" {
		err = errors.New("missing required session_id parameter")
		return
	}
	if turnID == "" {
		err = errors.New("missing required turn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s/session/%s/turn/%s/resume", params.AgentID, params.SessionID, turnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &raw, opts...)
	return ssestream.NewStream[AgentTurnResponseStreamChunk](ssestream.NewDecoder(raw), err)
}

// Streamed agent turn completion response.
type AgentTurnResponseStreamChunk struct {
	// Individual event in the agent turn response stream
	Event TurnResponseEvent `json:"event,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Event       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTurnResponseStreamChunk) RawJSON() string { return r.JSON.raw }
func (r *AgentTurnResponseStreamChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single turn in an interaction with an Agentic System.
type Turn struct {
	// List of messages that initiated this turn
	InputMessages []TurnInputMessageUnion `json:"input_messages,required"`
	// The model's generated response containing content and metadata
	OutputMessage CompletionMessage `json:"output_message,required"`
	// Unique identifier for the conversation session
	SessionID string `json:"session_id,required"`
	// Timestamp when the turn began
	StartedAt time.Time `json:"started_at,required" format:"date-time"`
	// Ordered list of processing steps executed during this turn
	Steps []TurnStepUnion `json:"steps,required"`
	// Unique identifier for the turn within a session
	TurnID string `json:"turn_id,required"`
	// (Optional) Timestamp when the turn finished, if completed
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// (Optional) Files or media attached to the agent's response
	OutputAttachments []TurnOutputAttachment `json:"output_attachments"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputMessages     respjson.Field
		OutputMessage     respjson.Field
		SessionID         respjson.Field
		StartedAt         respjson.Field
		Steps             respjson.Field
		TurnID            respjson.Field
		CompletedAt       respjson.Field
		OutputAttachments respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Turn) RawJSON() string { return r.JSON.raw }
func (r *Turn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnInputMessageUnion contains all possible properties and values from
// [UserMessage], [ToolResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TurnInputMessageUnion struct {
	// This field is from variant [UserMessage].
	Content InterleavedContentUnion `json:"content"`
	Role    string                  `json:"role"`
	// This field is from variant [UserMessage].
	Context InterleavedContentUnion `json:"context"`
	// This field is from variant [ToolResponseMessage].
	CallID string `json:"call_id"`
	JSON   struct {
		Content respjson.Field
		Role    respjson.Field
		Context respjson.Field
		CallID  respjson.Field
		raw     string
	} `json:"-"`
}

func (u TurnInputMessageUnion) AsUserMessage() (v UserMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnInputMessageUnion) AsToolResponseMessage() (v ToolResponseMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnInputMessageUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnInputMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnStepUnion contains all possible properties and values from [InferenceStep],
// [ToolExecutionStep], [ShieldCallStep], [MemoryRetrievalStep].
//
// Use the [TurnStepUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TurnStepUnion struct {
	// This field is from variant [InferenceStep].
	ModelResponse CompletionMessage `json:"model_response"`
	StepID        string            `json:"step_id"`
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType    string    `json:"step_type"`
	TurnID      string    `json:"turn_id"`
	CompletedAt time.Time `json:"completed_at"`
	StartedAt   time.Time `json:"started_at"`
	// This field is from variant [ToolExecutionStep].
	ToolCalls []ToolCall `json:"tool_calls"`
	// This field is from variant [ToolExecutionStep].
	ToolResponses []ToolResponse `json:"tool_responses"`
	// This field is from variant [ShieldCallStep].
	Violation SafetyViolation `json:"violation"`
	// This field is from variant [MemoryRetrievalStep].
	InsertedContext InterleavedContentUnion `json:"inserted_context"`
	// This field is from variant [MemoryRetrievalStep].
	VectorDBIDs string `json:"vector_db_ids"`
	JSON        struct {
		ModelResponse   respjson.Field
		StepID          respjson.Field
		StepType        respjson.Field
		TurnID          respjson.Field
		CompletedAt     respjson.Field
		StartedAt       respjson.Field
		ToolCalls       respjson.Field
		ToolResponses   respjson.Field
		Violation       respjson.Field
		InsertedContext respjson.Field
		VectorDBIDs     respjson.Field
		raw             string
	} `json:"-"`
}

// anyTurnStep is implemented by each variant of [TurnStepUnion] to add type safety
// for the return type of [TurnStepUnion.AsAny]
type anyTurnStep interface {
	implTurnStepUnion()
}

func (InferenceStep) implTurnStepUnion()       {}
func (ToolExecutionStep) implTurnStepUnion()   {}
func (ShieldCallStep) implTurnStepUnion()      {}
func (MemoryRetrievalStep) implTurnStepUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TurnStepUnion.AsAny().(type) {
//	case llamastackclient.InferenceStep:
//	case llamastackclient.ToolExecutionStep:
//	case llamastackclient.ShieldCallStep:
//	case llamastackclient.MemoryRetrievalStep:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TurnStepUnion) AsAny() anyTurnStep {
	switch u.StepType {
	case "inference":
		return u.AsInference()
	case "tool_execution":
		return u.AsToolExecution()
	case "shield_call":
		return u.AsShieldCall()
	case "memory_retrieval":
		return u.AsMemoryRetrieval()
	}
	return nil
}

func (u TurnStepUnion) AsInference() (v InferenceStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnStepUnion) AsToolExecution() (v ToolExecutionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnStepUnion) AsShieldCall() (v ShieldCallStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnStepUnion) AsMemoryRetrieval() (v MemoryRetrievalStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnStepUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An attachment to an agent turn.
type TurnOutputAttachment struct {
	// The content of the attachment.
	Content TurnOutputAttachmentContentUnion `json:"content,required"`
	// The MIME type of the attachment.
	MimeType string `json:"mime_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		MimeType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnOutputAttachment) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnOutputAttachmentContentUnion contains all possible properties and values
// from [string], [TurnOutputAttachmentContentImageContentItem],
// [TurnOutputAttachmentContentTextContentItem], [[]InterleavedContentItemUnion],
// [TurnOutputAttachmentContentURL].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInterleavedContentItemArray]
type TurnOutputAttachmentContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]InterleavedContentItemUnion]
	// instead of an object.
	OfInterleavedContentItemArray []InterleavedContentItemUnion `json:",inline"`
	// This field is from variant [TurnOutputAttachmentContentImageContentItem].
	Image TurnOutputAttachmentContentImageContentItemImage `json:"image"`
	Type  string                                           `json:"type"`
	// This field is from variant [TurnOutputAttachmentContentTextContentItem].
	Text string `json:"text"`
	// This field is from variant [TurnOutputAttachmentContentURL].
	Uri  string `json:"uri"`
	JSON struct {
		OfString                      respjson.Field
		OfInterleavedContentItemArray respjson.Field
		Image                         respjson.Field
		Type                          respjson.Field
		Text                          respjson.Field
		Uri                           respjson.Field
		raw                           string
	} `json:"-"`
}

func (u TurnOutputAttachmentContentUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnOutputAttachmentContentUnion) AsImageContentItem() (v TurnOutputAttachmentContentImageContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnOutputAttachmentContentUnion) AsTextContentItem() (v TurnOutputAttachmentContentTextContentItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnOutputAttachmentContentUnion) AsInterleavedContentItemArray() (v []InterleavedContentItemUnion) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnOutputAttachmentContentUnion) AsURL() (v TurnOutputAttachmentContentURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnOutputAttachmentContentUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnOutputAttachmentContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A image content item
type TurnOutputAttachmentContentImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image TurnOutputAttachmentContentImageContentItemImage `json:"image,required"`
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
func (r TurnOutputAttachmentContentImageContentItem) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachmentContentImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type TurnOutputAttachmentContentImageContentItemImage struct {
	// base64 encoded image data as string
	Data string `json:"data"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL TurnOutputAttachmentContentImageContentItemImageURL `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnOutputAttachmentContentImageContentItemImage) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachmentContentImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
type TurnOutputAttachmentContentImageContentItemImageURL struct {
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
func (r TurnOutputAttachmentContentImageContentItemImageURL) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachmentContentImageContentItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
type TurnOutputAttachmentContentTextContentItem struct {
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
func (r TurnOutputAttachmentContentTextContentItem) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachmentContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
type TurnOutputAttachmentContentURL struct {
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
func (r TurnOutputAttachmentContentURL) RawJSON() string { return r.JSON.raw }
func (r *TurnOutputAttachmentContentURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An event in an agent turn response stream.
type TurnResponseEvent struct {
	// Event-specific payload containing event data
	Payload TurnResponseEventPayloadUnion `json:"payload,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Payload     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnResponseEvent) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnResponseEventPayloadUnion contains all possible properties and values from
// [TurnResponseEventPayloadStepStart], [TurnResponseEventPayloadStepProgress],
// [TurnResponseEventPayloadStepComplete], [TurnResponseEventPayloadTurnStart],
// [TurnResponseEventPayloadTurnComplete],
// [TurnResponseEventPayloadTurnAwaitingInput].
//
// Use the [TurnResponseEventPayloadUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TurnResponseEventPayloadUnion struct {
	// Any of "step_start", "step_progress", "step_complete", "turn_start",
	// "turn_complete", "turn_awaiting_input".
	EventType string `json:"event_type"`
	StepID    string `json:"step_id"`
	StepType  string `json:"step_type"`
	// This field is from variant [TurnResponseEventPayloadStepStart].
	Metadata map[string]TurnResponseEventPayloadStepStartMetadataUnion `json:"metadata"`
	// This field is from variant [TurnResponseEventPayloadStepProgress].
	Delta TurnResponseEventPayloadStepProgressDeltaUnion `json:"delta"`
	// This field is from variant [TurnResponseEventPayloadStepComplete].
	StepDetails TurnResponseEventPayloadStepCompleteStepDetailsUnion `json:"step_details"`
	// This field is from variant [TurnResponseEventPayloadTurnStart].
	TurnID string `json:"turn_id"`
	// This field is from variant [TurnResponseEventPayloadTurnComplete].
	Turn Turn `json:"turn"`
	JSON struct {
		EventType   respjson.Field
		StepID      respjson.Field
		StepType    respjson.Field
		Metadata    respjson.Field
		Delta       respjson.Field
		StepDetails respjson.Field
		TurnID      respjson.Field
		Turn        respjson.Field
		raw         string
	} `json:"-"`
}

// anyTurnResponseEventPayload is implemented by each variant of
// [TurnResponseEventPayloadUnion] to add type safety for the return type of
// [TurnResponseEventPayloadUnion.AsAny]
type anyTurnResponseEventPayload interface {
	implTurnResponseEventPayloadUnion()
}

func (TurnResponseEventPayloadStepStart) implTurnResponseEventPayloadUnion()         {}
func (TurnResponseEventPayloadStepProgress) implTurnResponseEventPayloadUnion()      {}
func (TurnResponseEventPayloadStepComplete) implTurnResponseEventPayloadUnion()      {}
func (TurnResponseEventPayloadTurnStart) implTurnResponseEventPayloadUnion()         {}
func (TurnResponseEventPayloadTurnComplete) implTurnResponseEventPayloadUnion()      {}
func (TurnResponseEventPayloadTurnAwaitingInput) implTurnResponseEventPayloadUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TurnResponseEventPayloadUnion.AsAny().(type) {
//	case llamastackclient.TurnResponseEventPayloadStepStart:
//	case llamastackclient.TurnResponseEventPayloadStepProgress:
//	case llamastackclient.TurnResponseEventPayloadStepComplete:
//	case llamastackclient.TurnResponseEventPayloadTurnStart:
//	case llamastackclient.TurnResponseEventPayloadTurnComplete:
//	case llamastackclient.TurnResponseEventPayloadTurnAwaitingInput:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TurnResponseEventPayloadUnion) AsAny() anyTurnResponseEventPayload {
	switch u.EventType {
	case "step_start":
		return u.AsStepStart()
	case "step_progress":
		return u.AsStepProgress()
	case "step_complete":
		return u.AsStepComplete()
	case "turn_start":
		return u.AsTurnStart()
	case "turn_complete":
		return u.AsTurnComplete()
	case "turn_awaiting_input":
		return u.AsTurnAwaitingInput()
	}
	return nil
}

func (u TurnResponseEventPayloadUnion) AsStepStart() (v TurnResponseEventPayloadStepStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadUnion) AsStepProgress() (v TurnResponseEventPayloadStepProgress) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadUnion) AsStepComplete() (v TurnResponseEventPayloadStepComplete) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadUnion) AsTurnStart() (v TurnResponseEventPayloadTurnStart) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadUnion) AsTurnComplete() (v TurnResponseEventPayloadTurnComplete) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadUnion) AsTurnAwaitingInput() (v TurnResponseEventPayloadTurnAwaitingInput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnResponseEventPayloadUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnResponseEventPayloadUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for step start events in agent turn responses.
type TurnResponseEventPayloadStepStart struct {
	// Type of event being reported
	EventType constant.StepStart `json:"event_type,required"`
	// Unique identifier for the step within a turn
	StepID string `json:"step_id,required"`
	// Type of step being executed
	//
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType string `json:"step_type,required"`
	// (Optional) Additional metadata for the step
	Metadata map[string]TurnResponseEventPayloadStepStartMetadataUnion `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		StepID      respjson.Field
		StepType    respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnResponseEventPayloadStepStart) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadStepStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnResponseEventPayloadStepStartMetadataUnion contains all possible properties
// and values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type TurnResponseEventPayloadStepStartMetadataUnion struct {
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

func (u TurnResponseEventPayloadStepStartMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepStartMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepStartMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepStartMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnResponseEventPayloadStepStartMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnResponseEventPayloadStepStartMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for step progress events in agent turn responses.
type TurnResponseEventPayloadStepProgress struct {
	// Incremental content changes during step execution
	Delta TurnResponseEventPayloadStepProgressDeltaUnion `json:"delta,required"`
	// Type of event being reported
	EventType constant.StepProgress `json:"event_type,required"`
	// Unique identifier for the step within a turn
	StepID string `json:"step_id,required"`
	// Type of step being executed
	//
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType string `json:"step_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Delta       respjson.Field
		EventType   respjson.Field
		StepID      respjson.Field
		StepType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnResponseEventPayloadStepProgress) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadStepProgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnResponseEventPayloadStepProgressDeltaUnion contains all possible properties
// and values from [TurnResponseEventPayloadStepProgressDeltaText],
// [TurnResponseEventPayloadStepProgressDeltaImage],
// [TurnResponseEventPayloadStepProgressDeltaToolCall].
//
// Use the [TurnResponseEventPayloadStepProgressDeltaUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TurnResponseEventPayloadStepProgressDeltaUnion struct {
	// This field is from variant [TurnResponseEventPayloadStepProgressDeltaText].
	Text string `json:"text"`
	// Any of "text", "image", "tool_call".
	Type string `json:"type"`
	// This field is from variant [TurnResponseEventPayloadStepProgressDeltaImage].
	Image string `json:"image"`
	// This field is from variant [TurnResponseEventPayloadStepProgressDeltaToolCall].
	ParseStatus string `json:"parse_status"`
	// This field is from variant [TurnResponseEventPayloadStepProgressDeltaToolCall].
	ToolCall TurnResponseEventPayloadStepProgressDeltaToolCallToolCallUnion `json:"tool_call"`
	JSON     struct {
		Text        respjson.Field
		Type        respjson.Field
		Image       respjson.Field
		ParseStatus respjson.Field
		ToolCall    respjson.Field
		raw         string
	} `json:"-"`
}

// anyTurnResponseEventPayloadStepProgressDelta is implemented by each variant of
// [TurnResponseEventPayloadStepProgressDeltaUnion] to add type safety for the
// return type of [TurnResponseEventPayloadStepProgressDeltaUnion.AsAny]
type anyTurnResponseEventPayloadStepProgressDelta interface {
	implTurnResponseEventPayloadStepProgressDeltaUnion()
}

func (TurnResponseEventPayloadStepProgressDeltaText) implTurnResponseEventPayloadStepProgressDeltaUnion() {
}
func (TurnResponseEventPayloadStepProgressDeltaImage) implTurnResponseEventPayloadStepProgressDeltaUnion() {
}
func (TurnResponseEventPayloadStepProgressDeltaToolCall) implTurnResponseEventPayloadStepProgressDeltaUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := TurnResponseEventPayloadStepProgressDeltaUnion.AsAny().(type) {
//	case llamastackclient.TurnResponseEventPayloadStepProgressDeltaText:
//	case llamastackclient.TurnResponseEventPayloadStepProgressDeltaImage:
//	case llamastackclient.TurnResponseEventPayloadStepProgressDeltaToolCall:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TurnResponseEventPayloadStepProgressDeltaUnion) AsAny() anyTurnResponseEventPayloadStepProgressDelta {
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

func (u TurnResponseEventPayloadStepProgressDeltaUnion) AsText() (v TurnResponseEventPayloadStepProgressDeltaText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepProgressDeltaUnion) AsImage() (v TurnResponseEventPayloadStepProgressDeltaImage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepProgressDeltaUnion) AsToolCall() (v TurnResponseEventPayloadStepProgressDeltaToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnResponseEventPayloadStepProgressDeltaUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnResponseEventPayloadStepProgressDeltaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content delta for streaming responses.
type TurnResponseEventPayloadStepProgressDeltaText struct {
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
func (r TurnResponseEventPayloadStepProgressDeltaText) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadStepProgressDeltaText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An image content delta for streaming responses.
type TurnResponseEventPayloadStepProgressDeltaImage struct {
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
func (r TurnResponseEventPayloadStepProgressDeltaImage) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadStepProgressDeltaImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tool call content delta for streaming responses.
type TurnResponseEventPayloadStepProgressDeltaToolCall struct {
	// Current parsing status of the tool call
	//
	// Any of "started", "in_progress", "failed", "succeeded".
	ParseStatus string `json:"parse_status,required"`
	// Either an in-progress tool call string or the final parsed tool call
	ToolCall TurnResponseEventPayloadStepProgressDeltaToolCallToolCallUnion `json:"tool_call,required"`
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
func (r TurnResponseEventPayloadStepProgressDeltaToolCall) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadStepProgressDeltaToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnResponseEventPayloadStepProgressDeltaToolCallToolCallUnion contains all
// possible properties and values from [string], [ToolCall].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type TurnResponseEventPayloadStepProgressDeltaToolCallToolCallUnion struct {
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

func (u TurnResponseEventPayloadStepProgressDeltaToolCallToolCallUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepProgressDeltaToolCallToolCallUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnResponseEventPayloadStepProgressDeltaToolCallToolCallUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *TurnResponseEventPayloadStepProgressDeltaToolCallToolCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for step completion events in agent turn responses.
type TurnResponseEventPayloadStepComplete struct {
	// Type of event being reported
	EventType constant.StepComplete `json:"event_type,required"`
	// Complete details of the executed step
	StepDetails TurnResponseEventPayloadStepCompleteStepDetailsUnion `json:"step_details,required"`
	// Unique identifier for the step within a turn
	StepID string `json:"step_id,required"`
	// Type of step being executed
	//
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType string `json:"step_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		StepDetails respjson.Field
		StepID      respjson.Field
		StepType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnResponseEventPayloadStepComplete) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadStepComplete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TurnResponseEventPayloadStepCompleteStepDetailsUnion contains all possible
// properties and values from [InferenceStep], [ToolExecutionStep],
// [ShieldCallStep], [MemoryRetrievalStep].
//
// Use the [TurnResponseEventPayloadStepCompleteStepDetailsUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TurnResponseEventPayloadStepCompleteStepDetailsUnion struct {
	// This field is from variant [InferenceStep].
	ModelResponse CompletionMessage `json:"model_response"`
	StepID        string            `json:"step_id"`
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType    string    `json:"step_type"`
	TurnID      string    `json:"turn_id"`
	CompletedAt time.Time `json:"completed_at"`
	StartedAt   time.Time `json:"started_at"`
	// This field is from variant [ToolExecutionStep].
	ToolCalls []ToolCall `json:"tool_calls"`
	// This field is from variant [ToolExecutionStep].
	ToolResponses []ToolResponse `json:"tool_responses"`
	// This field is from variant [ShieldCallStep].
	Violation SafetyViolation `json:"violation"`
	// This field is from variant [MemoryRetrievalStep].
	InsertedContext InterleavedContentUnion `json:"inserted_context"`
	// This field is from variant [MemoryRetrievalStep].
	VectorDBIDs string `json:"vector_db_ids"`
	JSON        struct {
		ModelResponse   respjson.Field
		StepID          respjson.Field
		StepType        respjson.Field
		TurnID          respjson.Field
		CompletedAt     respjson.Field
		StartedAt       respjson.Field
		ToolCalls       respjson.Field
		ToolResponses   respjson.Field
		Violation       respjson.Field
		InsertedContext respjson.Field
		VectorDBIDs     respjson.Field
		raw             string
	} `json:"-"`
}

// anyTurnResponseEventPayloadStepCompleteStepDetails is implemented by each
// variant of [TurnResponseEventPayloadStepCompleteStepDetailsUnion] to add type
// safety for the return type of
// [TurnResponseEventPayloadStepCompleteStepDetailsUnion.AsAny]
type anyTurnResponseEventPayloadStepCompleteStepDetails interface {
	implTurnResponseEventPayloadStepCompleteStepDetailsUnion()
}

func (InferenceStep) implTurnResponseEventPayloadStepCompleteStepDetailsUnion()       {}
func (ToolExecutionStep) implTurnResponseEventPayloadStepCompleteStepDetailsUnion()   {}
func (ShieldCallStep) implTurnResponseEventPayloadStepCompleteStepDetailsUnion()      {}
func (MemoryRetrievalStep) implTurnResponseEventPayloadStepCompleteStepDetailsUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TurnResponseEventPayloadStepCompleteStepDetailsUnion.AsAny().(type) {
//	case llamastackclient.InferenceStep:
//	case llamastackclient.ToolExecutionStep:
//	case llamastackclient.ShieldCallStep:
//	case llamastackclient.MemoryRetrievalStep:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TurnResponseEventPayloadStepCompleteStepDetailsUnion) AsAny() anyTurnResponseEventPayloadStepCompleteStepDetails {
	switch u.StepType {
	case "inference":
		return u.AsInference()
	case "tool_execution":
		return u.AsToolExecution()
	case "shield_call":
		return u.AsShieldCall()
	case "memory_retrieval":
		return u.AsMemoryRetrieval()
	}
	return nil
}

func (u TurnResponseEventPayloadStepCompleteStepDetailsUnion) AsInference() (v InferenceStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepCompleteStepDetailsUnion) AsToolExecution() (v ToolExecutionStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepCompleteStepDetailsUnion) AsShieldCall() (v ShieldCallStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnResponseEventPayloadStepCompleteStepDetailsUnion) AsMemoryRetrieval() (v MemoryRetrievalStep) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TurnResponseEventPayloadStepCompleteStepDetailsUnion) RawJSON() string { return u.JSON.raw }

func (r *TurnResponseEventPayloadStepCompleteStepDetailsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for turn start events in agent turn responses.
type TurnResponseEventPayloadTurnStart struct {
	// Type of event being reported
	EventType constant.TurnStart `json:"event_type,required"`
	// Unique identifier for the turn within a session
	TurnID string `json:"turn_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		TurnID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnResponseEventPayloadTurnStart) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadTurnStart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for turn completion events in agent turn responses.
type TurnResponseEventPayloadTurnComplete struct {
	// Type of event being reported
	EventType constant.TurnComplete `json:"event_type,required"`
	// Complete turn data including all steps and results
	Turn Turn `json:"turn,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		Turn        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnResponseEventPayloadTurnComplete) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadTurnComplete) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for turn awaiting input events in agent turn responses.
type TurnResponseEventPayloadTurnAwaitingInput struct {
	// Type of event being reported
	EventType constant.TurnAwaitingInput `json:"event_type,required"`
	// Turn data when waiting for external tool responses
	Turn Turn `json:"turn,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EventType   respjson.Field
		Turn        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnResponseEventPayloadTurnAwaitingInput) RawJSON() string { return r.JSON.raw }
func (r *TurnResponseEventPayloadTurnAwaitingInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AlphaAgentTurnNewParams struct {
	AgentID string `path:"agent_id,required" json:"-"`
	// List of messages to start the turn with.
	Messages []AlphaAgentTurnNewParamsMessageUnion `json:"messages,omitzero,required"`
	// (Optional) List of documents to create the turn with.
	Documents []AlphaAgentTurnNewParamsDocument `json:"documents,omitzero"`
	// (Optional) The tool configuration to create the turn with, will be used to
	// override the agent's tool_config.
	ToolConfig AlphaAgentTurnNewParamsToolConfig `json:"tool_config,omitzero"`
	// (Optional) List of toolgroups to create the turn with, will be used in addition
	// to the agent's config toolgroups for the request.
	Toolgroups []AlphaAgentTurnNewParamsToolgroupUnion `json:"toolgroups,omitzero"`
	paramObj
}

func (r AlphaAgentTurnNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaAgentTurnNewParamsMessageUnion struct {
	OfUserMessage         *UserMessageParam         `json:",omitzero,inline"`
	OfToolResponseMessage *ToolResponseMessageParam `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaAgentTurnNewParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUserMessage, u.OfToolResponseMessage)
}
func (u *AlphaAgentTurnNewParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaAgentTurnNewParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfUserMessage) {
		return u.OfUserMessage
	} else if !param.IsOmitted(u.OfToolResponseMessage) {
		return u.OfToolResponseMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaAgentTurnNewParamsMessageUnion) GetContext() *InterleavedContentUnionParam {
	if vt := u.OfUserMessage; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaAgentTurnNewParamsMessageUnion) GetCallID() *string {
	if vt := u.OfToolResponseMessage; vt != nil {
		return &vt.CallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaAgentTurnNewParamsMessageUnion) GetRole() *string {
	if vt := u.OfUserMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfToolResponseMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's Content property, if present.
func (u AlphaAgentTurnNewParamsMessageUnion) GetContent() *InterleavedContentUnionParam {
	if vt := u.OfUserMessage; vt != nil {
		return &vt.Content
	} else if vt := u.OfToolResponseMessage; vt != nil {
		return &vt.Content
	}
	return nil
}

// A document to be used by an agent.
//
// The properties Content, MimeType are required.
type AlphaAgentTurnNewParamsDocument struct {
	// The content of the document.
	Content AlphaAgentTurnNewParamsDocumentContentUnion `json:"content,omitzero,required"`
	// The MIME type of the document.
	MimeType string `json:"mime_type,required"`
	paramObj
}

func (r AlphaAgentTurnNewParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaAgentTurnNewParamsDocumentContentUnion struct {
	OfString                      param.Opt[string]                                       `json:",omitzero,inline"`
	OfImageContentItem            *AlphaAgentTurnNewParamsDocumentContentImageContentItem `json:",omitzero,inline"`
	OfTextContentItem             *AlphaAgentTurnNewParamsDocumentContentTextContentItem  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []InterleavedContentItemUnionParam                      `json:",omitzero,inline"`
	OfURL                         *AlphaAgentTurnNewParamsDocumentContentURL              `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaAgentTurnNewParamsDocumentContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfImageContentItem,
		u.OfTextContentItem,
		u.OfInterleavedContentItemArray,
		u.OfURL)
}
func (u *AlphaAgentTurnNewParamsDocumentContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaAgentTurnNewParamsDocumentContentUnion) asAny() any {
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
func (u AlphaAgentTurnNewParamsDocumentContentUnion) GetImage() *AlphaAgentTurnNewParamsDocumentContentImageContentItemImage {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaAgentTurnNewParamsDocumentContentUnion) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaAgentTurnNewParamsDocumentContentUnion) GetUri() *string {
	if vt := u.OfURL; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AlphaAgentTurnNewParamsDocumentContentUnion) GetType() *string {
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
type AlphaAgentTurnNewParamsDocumentContentImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image AlphaAgentTurnNewParamsDocumentContentImageContentItemImage `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r AlphaAgentTurnNewParamsDocumentContentImageContentItem) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParamsDocumentContentImageContentItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParamsDocumentContentImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type AlphaAgentTurnNewParamsDocumentContentImageContentItemImage struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL AlphaAgentTurnNewParamsDocumentContentImageContentItemImageURL `json:"url,omitzero"`
	paramObj
}

func (r AlphaAgentTurnNewParamsDocumentContentImageContentItemImage) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParamsDocumentContentImageContentItemImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParamsDocumentContentImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
//
// The property Uri is required.
type AlphaAgentTurnNewParamsDocumentContentImageContentItemImageURL struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r AlphaAgentTurnNewParamsDocumentContentImageContentItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParamsDocumentContentImageContentItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParamsDocumentContentImageContentItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type AlphaAgentTurnNewParamsDocumentContentTextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r AlphaAgentTurnNewParamsDocumentContentTextContentItem) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParamsDocumentContentTextContentItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParamsDocumentContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type AlphaAgentTurnNewParamsDocumentContentURL struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r AlphaAgentTurnNewParamsDocumentContentURL) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParamsDocumentContentURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParamsDocumentContentURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) The tool configuration to create the turn with, will be used to
// override the agent's tool_config.
type AlphaAgentTurnNewParamsToolConfig struct {
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

func (r AlphaAgentTurnNewParamsToolConfig) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParamsToolConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParamsToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AlphaAgentTurnNewParamsToolConfig](
		"system_message_behavior", "append", "replace",
	)
	apijson.RegisterFieldValidator[AlphaAgentTurnNewParamsToolConfig](
		"tool_prompt_format", "json", "function_tag", "python_list",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaAgentTurnNewParamsToolgroupUnion struct {
	OfString                 param.Opt[string]                                       `json:",omitzero,inline"`
	OfAgentToolGroupWithArgs *AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgs `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaAgentTurnNewParamsToolgroupUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAgentToolGroupWithArgs)
}
func (u *AlphaAgentTurnNewParamsToolgroupUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaAgentTurnNewParamsToolgroupUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAgentToolGroupWithArgs) {
		return u.OfAgentToolGroupWithArgs
	}
	return nil
}

// The properties Args, Name are required.
type AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgs struct {
	Args map[string]AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion `json:"args,omitzero,required"`
	Name string                                                                    `json:"name,required"`
	paramObj
}

func (r AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgs) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AlphaAgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion) asAny() any {
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

type AlphaAgentTurnGetParams struct {
	AgentID   string `path:"agent_id,required" json:"-"`
	SessionID string `path:"session_id,required" json:"-"`
	paramObj
}

type AlphaAgentTurnResumeParams struct {
	AgentID   string `path:"agent_id,required" json:"-"`
	SessionID string `path:"session_id,required" json:"-"`
	// The tool call responses to resume the turn with.
	ToolResponses []ToolResponseParam `json:"tool_responses,omitzero,required"`
	paramObj
}

func (r AlphaAgentTurnResumeParams) MarshalJSON() (data []byte, err error) {
	type shadow AlphaAgentTurnResumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AlphaAgentTurnResumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
