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
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// SafetyService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSafetyService] method instead.
type SafetyService struct {
	Options []option.RequestOption
}

// NewSafetyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSafetyService(opts ...option.RequestOption) (r SafetyService) {
	r = SafetyService{}
	r.Options = opts
	return
}

// Run shield. Run a shield.
func (r *SafetyService) RunShield(ctx context.Context, body SafetyRunShieldParams, opts ...option.RequestOption) (res *RunShieldResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/safety/run-shield"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response from running a safety shield.
type RunShieldResponse struct {
	// (Optional) Safety violation detected by the shield, if any
	Violation SafetyViolation `json:"violation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Violation   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RunShieldResponse) RawJSON() string { return r.JSON.raw }
func (r *RunShieldResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SafetyRunShieldParams struct {
	// The messages to run the shield on.
	Messages []SafetyRunShieldParamsMessageUnion `json:"messages,omitzero,required"`
	// The parameters of the shield.
	Params map[string]SafetyRunShieldParamsParamUnion `json:"params,omitzero,required"`
	// The identifier of the shield to run.
	ShieldID string `json:"shield_id,required"`
	paramObj
}

func (r SafetyRunShieldParams) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageUnion struct {
	OfUser      *SafetyRunShieldParamsMessageUser      `json:",omitzero,inline"`
	OfSystem    *SafetyRunShieldParamsMessageSystem    `json:",omitzero,inline"`
	OfAssistant *SafetyRunShieldParamsMessageAssistant `json:",omitzero,inline"`
	OfTool      *SafetyRunShieldParamsMessageTool      `json:",omitzero,inline"`
	OfDeveloper *SafetyRunShieldParamsMessageDeveloper `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUser,
		u.OfSystem,
		u.OfAssistant,
		u.OfTool,
		u.OfDeveloper)
}
func (u *SafetyRunShieldParamsMessageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageUnion) asAny() any {
	if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	} else if !param.IsOmitted(u.OfAssistant) {
		return u.OfAssistant
	} else if !param.IsOmitted(u.OfTool) {
		return u.OfTool
	} else if !param.IsOmitted(u.OfDeveloper) {
		return u.OfDeveloper
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUnion) GetToolCalls() []SafetyRunShieldParamsMessageAssistantToolCall {
	if vt := u.OfAssistant; vt != nil {
		return vt.ToolCalls
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUnion) GetToolCallID() *string {
	if vt := u.OfTool; vt != nil {
		return &vt.ToolCallID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUnion) GetRole() *string {
	if vt := u.OfUser; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfAssistant; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfTool; vt != nil {
		return (*string)(&vt.Role)
	} else if vt := u.OfDeveloper; vt != nil {
		return (*string)(&vt.Role)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUnion) GetName() *string {
	if vt := u.OfUser; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfSystem; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfAssistant; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	} else if vt := u.OfDeveloper; vt != nil && vt.Name.Valid() {
		return &vt.Name.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u SafetyRunShieldParamsMessageUnion) GetContent() (res safetyRunShieldParamsMessageUnionContent) {
	if vt := u.OfUser; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfSystem; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfAssistant; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfTool; vt != nil {
		res.any = vt.Content.asAny()
	} else if vt := u.OfDeveloper; vt != nil {
		res.any = vt.Content.asAny()
	}
	return
}

// Can have the runtime types [*string],
// [_[]SafetyRunShieldParamsMessageUserContentArrayItemUnion],
// [_[]SafetyRunShieldParamsMessageSystemContentArrayItem],
// [_[]SafetyRunShieldParamsMessageAssistantContentArrayItem],
// [_[]SafetyRunShieldParamsMessageToolContentArrayItem],
// [\*[]SafetyRunShieldParamsMessageDeveloperContentArrayItem]
type safetyRunShieldParamsMessageUnionContent struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageUserContentArrayItemUnion:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageSystemContentArrayItem:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageAssistantContentArrayItem:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageToolContentArrayItem:
//	case *[]llamastackclient.SafetyRunShieldParamsMessageDeveloperContentArrayItem:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u safetyRunShieldParamsMessageUnionContent) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[SafetyRunShieldParamsMessageUnion](
		"role",
		apijson.Discriminator[SafetyRunShieldParamsMessageUser]("user"),
		apijson.Discriminator[SafetyRunShieldParamsMessageSystem]("system"),
		apijson.Discriminator[SafetyRunShieldParamsMessageAssistant]("assistant"),
		apijson.Discriminator[SafetyRunShieldParamsMessageTool]("tool"),
		apijson.Discriminator[SafetyRunShieldParamsMessageDeveloper]("developer"),
	)
}

// A message from the user in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type SafetyRunShieldParamsMessageUser struct {
	// The content of the message, which can include text and other media
	Content SafetyRunShieldParamsMessageUserContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the user message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "user" to identify this as a user message
	//
	// This field can be elided, and will marshal its zero value as "user".
	Role constant.User `json:"role,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUser) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageUserContentUnion struct {
	OfString                                  param.Opt[string]                                       `json:",omitzero,inline"`
	OfSafetyRunShieldsMessageUserContentArray []SafetyRunShieldParamsMessageUserContentArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageUserContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfSafetyRunShieldsMessageUserContentArray)
}
func (u *SafetyRunShieldParamsMessageUserContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageUserContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfSafetyRunShieldsMessageUserContentArray) {
		return &u.OfSafetyRunShieldsMessageUserContentArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageUserContentArrayItemUnion struct {
	OfText     *SafetyRunShieldParamsMessageUserContentArrayItemText     `json:",omitzero,inline"`
	OfImageURL *SafetyRunShieldParamsMessageUserContentArrayItemImageURL `json:",omitzero,inline"`
	OfFile     *SafetyRunShieldParamsMessageUserContentArrayItemFile     `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageUserContentArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfText, u.OfImageURL, u.OfFile)
}
func (u *SafetyRunShieldParamsMessageUserContentArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageUserContentArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfText) {
		return u.OfText
	} else if !param.IsOmitted(u.OfImageURL) {
		return u.OfImageURL
	} else if !param.IsOmitted(u.OfFile) {
		return u.OfFile
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUserContentArrayItemUnion) GetText() *string {
	if vt := u.OfText; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUserContentArrayItemUnion) GetImageURL() *SafetyRunShieldParamsMessageUserContentArrayItemImageURLImageURL {
	if vt := u.OfImageURL; vt != nil {
		return &vt.ImageURL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUserContentArrayItemUnion) GetFile() *SafetyRunShieldParamsMessageUserContentArrayItemFileFile {
	if vt := u.OfFile; vt != nil {
		return &vt.File
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SafetyRunShieldParamsMessageUserContentArrayItemUnion) GetType() *string {
	if vt := u.OfText; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfImageURL; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfFile; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SafetyRunShieldParamsMessageUserContentArrayItemUnion](
		"type",
		apijson.Discriminator[SafetyRunShieldParamsMessageUserContentArrayItemText]("text"),
		apijson.Discriminator[SafetyRunShieldParamsMessageUserContentArrayItemImageURL]("image_url"),
		apijson.Discriminator[SafetyRunShieldParamsMessageUserContentArrayItemFile]("file"),
	)
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type SafetyRunShieldParamsMessageUserContentArrayItemText struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentArrayItemText) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentArrayItemText
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentArrayItemText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image content part for OpenAI-compatible chat completion messages.
//
// The properties ImageURL, Type are required.
type SafetyRunShieldParamsMessageUserContentArrayItemImageURL struct {
	// Image URL specification and processing details
	ImageURL SafetyRunShieldParamsMessageUserContentArrayItemImageURLImageURL `json:"image_url,omitzero,required"`
	// Must be "image_url" to identify this as image content
	//
	// This field can be elided, and will marshal its zero value as "image_url".
	Type constant.ImageURL `json:"type,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentArrayItemImageURL) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentArrayItemImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentArrayItemImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Image URL specification and processing details
//
// The property URL is required.
type SafetyRunShieldParamsMessageUserContentArrayItemImageURLImageURL struct {
	// URL of the image to include in the message
	URL string `json:"url,required"`
	// (Optional) Level of detail for image processing. Can be "low", "high", or "auto"
	Detail param.Opt[string] `json:"detail,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentArrayItemImageURLImageURL) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentArrayItemImageURLImageURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentArrayItemImageURLImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties File, Type are required.
type SafetyRunShieldParamsMessageUserContentArrayItemFile struct {
	File SafetyRunShieldParamsMessageUserContentArrayItemFileFile `json:"file,omitzero,required"`
	// This field can be elided, and will marshal its zero value as "file".
	Type constant.File `json:"type,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentArrayItemFile) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentArrayItemFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentArrayItemFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SafetyRunShieldParamsMessageUserContentArrayItemFileFile struct {
	FileData param.Opt[string] `json:"file_data,omitzero"`
	FileID   param.Opt[string] `json:"file_id,omitzero"`
	Filename param.Opt[string] `json:"filename,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageUserContentArrayItemFileFile) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageUserContentArrayItemFileFile
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageUserContentArrayItemFileFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system message providing instructions or context to the model.
//
// The properties Content, Role are required.
type SafetyRunShieldParamsMessageSystem struct {
	// The content of the "system prompt". If multiple system messages are provided,
	// they are concatenated. The underlying Llama Stack code may also add other system
	// messages (for example, for formatting tool definitions).
	Content SafetyRunShieldParamsMessageSystemContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the system message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "system" to identify this as a system message
	//
	// This field can be elided, and will marshal its zero value as "system".
	Role constant.System `json:"role,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageSystem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageSystem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageSystem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageSystemContentUnion struct {
	OfString                                    param.Opt[string]                                    `json:",omitzero,inline"`
	OfSafetyRunShieldsMessageSystemContentArray []SafetyRunShieldParamsMessageSystemContentArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageSystemContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfSafetyRunShieldsMessageSystemContentArray)
}
func (u *SafetyRunShieldParamsMessageSystemContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageSystemContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfSafetyRunShieldsMessageSystemContentArray) {
		return &u.OfSafetyRunShieldsMessageSystemContentArray
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type SafetyRunShieldParamsMessageSystemContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageSystemContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageSystemContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageSystemContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message containing the model's (assistant) response in an OpenAI-compatible
// chat completion request.
//
// The property Role is required.
type SafetyRunShieldParamsMessageAssistant struct {
	// (Optional) The name of the assistant message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// The content of the model's response
	Content SafetyRunShieldParamsMessageAssistantContentUnion `json:"content,omitzero"`
	// List of tool calls. Each tool call is an OpenAIChatCompletionToolCall object.
	ToolCalls []SafetyRunShieldParamsMessageAssistantToolCall `json:"tool_calls,omitzero"`
	// Must be "assistant" to identify this as the model's response
	//
	// This field can be elided, and will marshal its zero value as "assistant".
	Role constant.Assistant `json:"role,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageAssistant) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageAssistant
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageAssistantContentUnion struct {
	OfString                                       param.Opt[string]                                       `json:",omitzero,inline"`
	OfSafetyRunShieldsMessageAssistantContentArray []SafetyRunShieldParamsMessageAssistantContentArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageAssistantContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfSafetyRunShieldsMessageAssistantContentArray)
}
func (u *SafetyRunShieldParamsMessageAssistantContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageAssistantContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfSafetyRunShieldsMessageAssistantContentArray) {
		return &u.OfSafetyRunShieldsMessageAssistantContentArray
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type SafetyRunShieldParamsMessageAssistantContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageAssistantContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageAssistantContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageAssistantContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tool call specification for OpenAI-compatible chat completion responses.
//
// The property Type is required.
type SafetyRunShieldParamsMessageAssistantToolCall struct {
	// (Optional) Unique identifier for the tool call
	ID param.Opt[string] `json:"id,omitzero"`
	// (Optional) Index of the tool call in the list
	Index param.Opt[int64] `json:"index,omitzero"`
	// (Optional) Function call details
	Function SafetyRunShieldParamsMessageAssistantToolCallFunction `json:"function,omitzero"`
	// Must be "function" to identify this as a function call
	//
	// This field can be elided, and will marshal its zero value as "function".
	Type constant.Function `json:"type,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageAssistantToolCall) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageAssistantToolCall
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageAssistantToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) Function call details
type SafetyRunShieldParamsMessageAssistantToolCallFunction struct {
	// (Optional) Arguments to pass to the function as a JSON string
	Arguments param.Opt[string] `json:"arguments,omitzero"`
	// (Optional) Name of the function to call
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r SafetyRunShieldParamsMessageAssistantToolCallFunction) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageAssistantToolCallFunction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageAssistantToolCallFunction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message representing the result of a tool invocation in an OpenAI-compatible
// chat completion request.
//
// The properties Content, Role, ToolCallID are required.
type SafetyRunShieldParamsMessageTool struct {
	// The response content from the tool
	Content SafetyRunShieldParamsMessageToolContentUnion `json:"content,omitzero,required"`
	// Unique identifier for the tool call this response is for
	ToolCallID string `json:"tool_call_id,required"`
	// Must be "tool" to identify this as a tool response
	//
	// This field can be elided, and will marshal its zero value as "tool".
	Role constant.Tool `json:"role,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageTool) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageTool
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageToolContentUnion struct {
	OfString                                  param.Opt[string]                                  `json:",omitzero,inline"`
	OfSafetyRunShieldsMessageToolContentArray []SafetyRunShieldParamsMessageToolContentArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageToolContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfSafetyRunShieldsMessageToolContentArray)
}
func (u *SafetyRunShieldParamsMessageToolContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageToolContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfSafetyRunShieldsMessageToolContentArray) {
		return &u.OfSafetyRunShieldsMessageToolContentArray
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type SafetyRunShieldParamsMessageToolContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageToolContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageToolContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageToolContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A message from the developer in an OpenAI-compatible chat completion request.
//
// The properties Content, Role are required.
type SafetyRunShieldParamsMessageDeveloper struct {
	// The content of the developer message
	Content SafetyRunShieldParamsMessageDeveloperContentUnion `json:"content,omitzero,required"`
	// (Optional) The name of the developer message participant.
	Name param.Opt[string] `json:"name,omitzero"`
	// Must be "developer" to identify this as a developer message
	//
	// This field can be elided, and will marshal its zero value as "developer".
	Role constant.Developer `json:"role,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageDeveloper) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageDeveloper
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageDeveloper) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsMessageDeveloperContentUnion struct {
	OfString                                       param.Opt[string]                                       `json:",omitzero,inline"`
	OfSafetyRunShieldsMessageDeveloperContentArray []SafetyRunShieldParamsMessageDeveloperContentArrayItem `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsMessageDeveloperContentUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfSafetyRunShieldsMessageDeveloperContentArray)
}
func (u *SafetyRunShieldParamsMessageDeveloperContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsMessageDeveloperContentUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfSafetyRunShieldsMessageDeveloperContentArray) {
		return &u.OfSafetyRunShieldsMessageDeveloperContentArray
	}
	return nil
}

// Text content part for OpenAI-compatible chat completion messages.
//
// The properties Text, Type are required.
type SafetyRunShieldParamsMessageDeveloperContentArrayItem struct {
	// The text content of the message
	Text string `json:"text,required"`
	// Must be "text" to identify this as text content
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r SafetyRunShieldParamsMessageDeveloperContentArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow SafetyRunShieldParamsMessageDeveloperContentArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SafetyRunShieldParamsMessageDeveloperContentArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SafetyRunShieldParamsParamUnion struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u SafetyRunShieldParamsParamUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *SafetyRunShieldParamsParamUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SafetyRunShieldParamsParamUnion) asAny() any {
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
