// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/packages/ssestream"
	"github.com/llamastack/llama-stack-client-go/shared"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// AgentTurnService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentTurnService] method instead.
type AgentTurnService struct {
	Options []option.RequestOption
}

// NewAgentTurnService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentTurnService(opts ...option.RequestOption) (r AgentTurnService) {
	r = AgentTurnService{}
	r.Options = opts
	return
}

// Create a new turn for an agent.
func (r *AgentTurnService) New(ctx context.Context, sessionID string, params AgentTurnNewParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *AgentTurnService) NewStreaming(ctx context.Context, sessionID string, params AgentTurnNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[AgentTurnResponseStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
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
func (r *AgentTurnService) Get(ctx context.Context, turnID string, query AgentTurnGetParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *AgentTurnService) Resume(ctx context.Context, turnID string, params AgentTurnResumeParams, opts ...option.RequestOption) (res *Turn, err error) {
	opts = append(r.Options[:], opts...)
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
func (r *AgentTurnService) ResumeStreaming(ctx context.Context, turnID string, params AgentTurnResumeParams, opts ...option.RequestOption) (stream *ssestream.Stream[AgentTurnResponseStreamChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = append(r.Options[:], opts...)
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
	OutputMessage shared.CompletionMessage `json:"output_message,required"`
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
// [shared.UserMessage], [shared.ToolResponseMessage].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TurnInputMessageUnion struct {
	// This field is from variant [shared.UserMessage].
	Content shared.InterleavedContentUnion `json:"content"`
	Role    string                         `json:"role"`
	// This field is from variant [shared.UserMessage].
	Context shared.InterleavedContentUnion `json:"context"`
	// This field is from variant [shared.ToolResponseMessage].
	CallID string `json:"call_id"`
	JSON   struct {
		Content respjson.Field
		Role    respjson.Field
		Context respjson.Field
		CallID  respjson.Field
		raw     string
	} `json:"-"`
}

func (u TurnInputMessageUnion) AsUserMessage() (v shared.UserMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TurnInputMessageUnion) AsToolResponseMessage() (v shared.ToolResponseMessage) {
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
	ModelResponse shared.CompletionMessage `json:"model_response"`
	StepID        string                   `json:"step_id"`
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType    string    `json:"step_type"`
	TurnID      string    `json:"turn_id"`
	CompletedAt time.Time `json:"completed_at"`
	StartedAt   time.Time `json:"started_at"`
	// This field is from variant [ToolExecutionStep].
	ToolCalls []shared.ToolCall `json:"tool_calls"`
	// This field is from variant [ToolExecutionStep].
	ToolResponses []ToolResponse `json:"tool_responses"`
	// This field is from variant [ShieldCallStep].
	Violation shared.SafetyViolation `json:"violation"`
	// This field is from variant [MemoryRetrievalStep].
	InsertedContext shared.InterleavedContentUnion `json:"inserted_context"`
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
// [TurnOutputAttachmentContentTextContentItem],
// [[]shared.InterleavedContentItemUnion], [TurnOutputAttachmentContentURL].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfInterleavedContentItemArray]
type TurnOutputAttachmentContentUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]shared.InterleavedContentItemUnion] instead of an object.
	OfInterleavedContentItemArray []shared.InterleavedContentItemUnion `json:",inline"`
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

func (u TurnOutputAttachmentContentUnion) AsInterleavedContentItemArray() (v []shared.InterleavedContentItemUnion) {
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
	Delta shared.ContentDeltaUnion `json:"delta"`
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
	Delta shared.ContentDeltaUnion `json:"delta,required"`
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
	ModelResponse shared.CompletionMessage `json:"model_response"`
	StepID        string                   `json:"step_id"`
	// Any of "inference", "tool_execution", "shield_call", "memory_retrieval".
	StepType    string    `json:"step_type"`
	TurnID      string    `json:"turn_id"`
	CompletedAt time.Time `json:"completed_at"`
	StartedAt   time.Time `json:"started_at"`
	// This field is from variant [ToolExecutionStep].
	ToolCalls []shared.ToolCall `json:"tool_calls"`
	// This field is from variant [ToolExecutionStep].
	ToolResponses []ToolResponse `json:"tool_responses"`
	// This field is from variant [ShieldCallStep].
	Violation shared.SafetyViolation `json:"violation"`
	// This field is from variant [MemoryRetrievalStep].
	InsertedContext shared.InterleavedContentUnion `json:"inserted_context"`
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

type AgentTurnNewParams struct {
	AgentID string `path:"agent_id,required" json:"-"`
	// List of messages to start the turn with.
	Messages []AgentTurnNewParamsMessageUnion `json:"messages,omitzero,required"`
	// (Optional) List of documents to create the turn with.
	Documents []AgentTurnNewParamsDocument `json:"documents,omitzero"`
	// (Optional) The tool configuration to create the turn with, will be used to
	// override the agent's tool_config.
	ToolConfig AgentTurnNewParamsToolConfig `json:"tool_config,omitzero"`
	// (Optional) List of toolgroups to create the turn with, will be used in addition
	// to the agent's config toolgroups for the request.
	Toolgroups []AgentTurnNewParamsToolgroupUnion `json:"toolgroups,omitzero"`
	paramObj
}

func (r AgentTurnNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentTurnNewParamsMessageUnion struct {
	OfUserMessage         *shared.UserMessageParam         `json:",omitzero,inline"`
	OfToolResponseMessage *shared.ToolResponseMessageParam `json:",omitzero,inline"`
	paramUnion
}

func (u AgentTurnNewParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUserMessage, u.OfToolResponseMessage)
}
func (u *AgentTurnNewParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentTurnNewParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfUserMessage) {
		return u.OfUserMessage
	} else if !param.IsOmitted(u.OfToolResponseMessage) {
		return u.OfToolResponseMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentTurnNewParamsMessageUnion) GetContext() *shared.InterleavedContentUnionParam {
	if vt := u.OfUserMessage; vt != nil {
		return &vt.Context
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentTurnNewParamsMessageUnion) GetCallID() *string {
	if vt := u.OfToolResponseMessage; vt != nil {
		return &vt.CallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentTurnNewParamsMessageUnion) GetRole() *string {
	if vt := u.OfUserMessage; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfToolResponseMessage; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's Content property, if present.
func (u AgentTurnNewParamsMessageUnion) GetContent() *shared.InterleavedContentUnionParam {
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
type AgentTurnNewParamsDocument struct {
	// The content of the document.
	Content AgentTurnNewParamsDocumentContentUnion `json:"content,omitzero,required"`
	// The MIME type of the document.
	MimeType string `json:"mime_type,required"`
	paramObj
}

func (r AgentTurnNewParamsDocument) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParamsDocument
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParamsDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentTurnNewParamsDocumentContentUnion struct {
	OfString                      param.Opt[string]                                  `json:",omitzero,inline"`
	OfImageContentItem            *AgentTurnNewParamsDocumentContentImageContentItem `json:",omitzero,inline"`
	OfTextContentItem             *AgentTurnNewParamsDocumentContentTextContentItem  `json:",omitzero,inline"`
	OfInterleavedContentItemArray []shared.InterleavedContentItemUnionParam          `json:",omitzero,inline"`
	OfURL                         *AgentTurnNewParamsDocumentContentURL              `json:",omitzero,inline"`
	paramUnion
}

func (u AgentTurnNewParamsDocumentContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString,
		u.OfImageContentItem,
		u.OfTextContentItem,
		u.OfInterleavedContentItemArray,
		u.OfURL)
}
func (u *AgentTurnNewParamsDocumentContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentTurnNewParamsDocumentContentUnion) asAny() any {
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
func (u AgentTurnNewParamsDocumentContentUnion) GetImage() *AgentTurnNewParamsDocumentContentImageContentItemImage {
	if vt := u.OfImageContentItem; vt != nil {
		return &vt.Image
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentTurnNewParamsDocumentContentUnion) GetText() *string {
	if vt := u.OfTextContentItem; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentTurnNewParamsDocumentContentUnion) GetUri() *string {
	if vt := u.OfURL; vt != nil {
		return &vt.Uri
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AgentTurnNewParamsDocumentContentUnion) GetType() *string {
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
type AgentTurnNewParamsDocumentContentImageContentItem struct {
	// Image as a base64 encoded string or an URL
	Image AgentTurnNewParamsDocumentContentImageContentItemImage `json:"image,omitzero,required"`
	// Discriminator type of the content item. Always "image"
	//
	// This field can be elided, and will marshal its zero value as "image".
	Type constant.Image `json:"type,required"`
	paramObj
}

func (r AgentTurnNewParamsDocumentContentImageContentItem) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParamsDocumentContentImageContentItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParamsDocumentContentImageContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image as a base64 encoded string or an URL
type AgentTurnNewParamsDocumentContentImageContentItemImage struct {
	// base64 encoded image data as string
	Data param.Opt[string] `json:"data,omitzero"`
	// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
	// Note that URL could have length limits.
	URL AgentTurnNewParamsDocumentContentImageContentItemImageURL `json:"url,omitzero"`
	paramObj
}

func (r AgentTurnNewParamsDocumentContentImageContentItemImage) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParamsDocumentContentImageContentItemImage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParamsDocumentContentImageContentItemImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL of the image or data URL in the format of data:image/{type};base64,{data}.
// Note that URL could have length limits.
//
// The property Uri is required.
type AgentTurnNewParamsDocumentContentImageContentItemImageURL struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r AgentTurnNewParamsDocumentContentImageContentItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParamsDocumentContentImageContentItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParamsDocumentContentImageContentItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text content item
//
// The properties Text, Type are required.
type AgentTurnNewParamsDocumentContentTextContentItem struct {
	// Text content
	Text string `json:"text,required"`
	// Discriminator type of the content item. Always "text"
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r AgentTurnNewParamsDocumentContentTextContentItem) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParamsDocumentContentTextContentItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParamsDocumentContentTextContentItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A URL reference to external content.
//
// The property Uri is required.
type AgentTurnNewParamsDocumentContentURL struct {
	// The URL string pointing to the resource
	Uri string `json:"uri,required"`
	paramObj
}

func (r AgentTurnNewParamsDocumentContentURL) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParamsDocumentContentURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParamsDocumentContentURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) The tool configuration to create the turn with, will be used to
// override the agent's tool_config.
type AgentTurnNewParamsToolConfig struct {
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

func (r AgentTurnNewParamsToolConfig) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParamsToolConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParamsToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentTurnNewParamsToolConfig](
		"system_message_behavior", "append", "replace",
	)
	apijson.RegisterFieldValidator[AgentTurnNewParamsToolConfig](
		"tool_prompt_format", "json", "function_tag", "python_list",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentTurnNewParamsToolgroupUnion struct {
	OfString                 param.Opt[string]                                  `json:",omitzero,inline"`
	OfAgentToolGroupWithArgs *AgentTurnNewParamsToolgroupAgentToolGroupWithArgs `json:",omitzero,inline"`
	paramUnion
}

func (u AgentTurnNewParamsToolgroupUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAgentToolGroupWithArgs)
}
func (u *AgentTurnNewParamsToolgroupUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentTurnNewParamsToolgroupUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAgentToolGroupWithArgs) {
		return u.OfAgentToolGroupWithArgs
	}
	return nil
}

// The properties Args, Name are required.
type AgentTurnNewParamsToolgroupAgentToolGroupWithArgs struct {
	Args map[string]AgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion `json:"args,omitzero,required"`
	Name string                                                               `json:"name,required"`
	paramObj
}

func (r AgentTurnNewParamsToolgroupAgentToolGroupWithArgs) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnNewParamsToolgroupAgentToolGroupWithArgs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnNewParamsToolgroupAgentToolGroupWithArgs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u AgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *AgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentTurnNewParamsToolgroupAgentToolGroupWithArgsArgUnion) asAny() any {
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

type AgentTurnGetParams struct {
	AgentID   string `path:"agent_id,required" json:"-"`
	SessionID string `path:"session_id,required" json:"-"`
	paramObj
}

type AgentTurnResumeParams struct {
	AgentID   string `path:"agent_id,required" json:"-"`
	SessionID string `path:"session_id,required" json:"-"`
	// The tool call responses to resume the turn with.
	ToolResponses []ToolResponseParam `json:"tool_responses,omitzero,required"`
	paramObj
}

func (r AgentTurnResumeParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentTurnResumeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentTurnResumeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
