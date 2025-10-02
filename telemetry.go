// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// TelemetryService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTelemetryService] method instead.
type TelemetryService struct {
	Options []option.RequestOption
}

// NewTelemetryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTelemetryService(opts ...option.RequestOption) (r TelemetryService) {
	r = TelemetryService{}
	r.Options = opts
	return
}

// Log an event.
func (r *TelemetryService) LogEvent(ctx context.Context, body TelemetryLogEventParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/telemetry/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EventUnionParam struct {
	OfUnstructuredLog *EventUnstructuredLogParam `json:",omitzero,inline"`
	OfMetric          *EventMetricParam          `json:",omitzero,inline"`
	OfStructuredLog   *EventStructuredLogParam   `json:",omitzero,inline"`
	paramUnion
}

func (u EventUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfUnstructuredLog, u.OfMetric, u.OfStructuredLog)
}
func (u *EventUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EventUnionParam) asAny() any {
	if !param.IsOmitted(u.OfUnstructuredLog) {
		return u.OfUnstructuredLog
	} else if !param.IsOmitted(u.OfMetric) {
		return u.OfMetric
	} else if !param.IsOmitted(u.OfStructuredLog) {
		return u.OfStructuredLog
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetMessage() *string {
	if vt := u.OfUnstructuredLog; vt != nil {
		return &vt.Message
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetSeverity() *string {
	if vt := u.OfUnstructuredLog; vt != nil {
		return &vt.Severity
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetMetric() *string {
	if vt := u.OfMetric; vt != nil {
		return &vt.Metric
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetUnit() *string {
	if vt := u.OfMetric; vt != nil {
		return &vt.Unit
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetValue() *float64 {
	if vt := u.OfMetric; vt != nil {
		return &vt.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetPayload() *EventStructuredLogPayloadUnionParam {
	if vt := u.OfStructuredLog; vt != nil {
		return &vt.Payload
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetSpanID() *string {
	if vt := u.OfUnstructuredLog; vt != nil {
		return (*string)(&vt.SpanID)
	} else if vt := u.OfMetric; vt != nil {
		return (*string)(&vt.SpanID)
	} else if vt := u.OfStructuredLog; vt != nil {
		return (*string)(&vt.SpanID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetTraceID() *string {
	if vt := u.OfUnstructuredLog; vt != nil {
		return (*string)(&vt.TraceID)
	} else if vt := u.OfMetric; vt != nil {
		return (*string)(&vt.TraceID)
	} else if vt := u.OfStructuredLog; vt != nil {
		return (*string)(&vt.TraceID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventUnionParam) GetType() *string {
	if vt := u.OfUnstructuredLog; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMetric; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfStructuredLog; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Timestamp property, if present.
func (u EventUnionParam) GetTimestamp() *time.Time {
	if vt := u.OfUnstructuredLog; vt != nil {
		return &vt.Timestamp
	} else if vt := u.OfMetric; vt != nil {
		return &vt.Timestamp
	} else if vt := u.OfStructuredLog; vt != nil {
		return &vt.Timestamp
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u EventUnionParam) GetAttributes() (res eventUnionParamAttributes) {
	if vt := u.OfUnstructuredLog; vt != nil {
		res.any = &vt.Attributes
	} else if vt := u.OfMetric; vt != nil {
		res.any = &vt.Attributes
	} else if vt := u.OfStructuredLog; vt != nil {
		res.any = &vt.Attributes
	}
	return
}

// Can have the runtime types
// [*map[string]EventUnstructuredLogAttributeUnionParam],
// [*map[string]EventMetricAttributeUnionParam],
// [\*map[string]EventStructuredLogAttributeUnionParam]
type eventUnionParamAttributes struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *map[string]llamastackclient.EventUnstructuredLogAttributeUnionParam:
//	case *map[string]llamastackclient.EventMetricAttributeUnionParam:
//	case *map[string]llamastackclient.EventStructuredLogAttributeUnionParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u eventUnionParamAttributes) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[EventUnionParam](
		"type",
		apijson.Discriminator[EventUnstructuredLogParam]("unstructured_log"),
		apijson.Discriminator[EventMetricParam]("metric"),
		apijson.Discriminator[EventStructuredLogParam]("structured_log"),
	)
}

// An unstructured log event containing a simple text message.
//
// The properties Message, Severity, SpanID, Timestamp, TraceID, Type are required.
type EventUnstructuredLogParam struct {
	// The log message text
	Message string `json:"message,required"`
	// The severity level of the log message
	//
	// Any of "verbose", "debug", "info", "warn", "error", "critical".
	Severity string `json:"severity,omitzero,required"`
	// Unique identifier for the span this event belongs to
	SpanID string `json:"span_id,required"`
	// Timestamp when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Unique identifier for the trace this event belongs to
	TraceID string `json:"trace_id,required"`
	// (Optional) Key-value pairs containing additional metadata about the event
	Attributes map[string]EventUnstructuredLogAttributeUnionParam `json:"attributes,omitzero"`
	// Event type identifier set to UNSTRUCTURED_LOG
	//
	// This field can be elided, and will marshal its zero value as "unstructured_log".
	Type constant.UnstructuredLog `json:"type,required"`
	paramObj
}

func (r EventUnstructuredLogParam) MarshalJSON() (data []byte, err error) {
	type shadow EventUnstructuredLogParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventUnstructuredLogParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EventUnstructuredLogParam](
		"severity", "verbose", "debug", "info", "warn", "error", "critical",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EventUnstructuredLogAttributeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u EventUnstructuredLogAttributeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *EventUnstructuredLogAttributeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EventUnstructuredLogAttributeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// A metric event containing a measured value.
//
// The properties Metric, SpanID, Timestamp, TraceID, Type, Unit, Value are
// required.
type EventMetricParam struct {
	// The name of the metric being measured
	Metric string `json:"metric,required"`
	// Unique identifier for the span this event belongs to
	SpanID string `json:"span_id,required"`
	// Timestamp when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Unique identifier for the trace this event belongs to
	TraceID string `json:"trace_id,required"`
	// The unit of measurement for the metric value
	Unit string `json:"unit,required"`
	// The numeric value of the metric measurement
	Value float64 `json:"value,required"`
	// (Optional) Key-value pairs containing additional metadata about the event
	Attributes map[string]EventMetricAttributeUnionParam `json:"attributes,omitzero"`
	// Event type identifier set to METRIC
	//
	// This field can be elided, and will marshal its zero value as "metric".
	Type constant.Metric `json:"type,required"`
	paramObj
}

func (r EventMetricParam) MarshalJSON() (data []byte, err error) {
	type shadow EventMetricParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventMetricParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EventMetricAttributeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u EventMetricAttributeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *EventMetricAttributeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EventMetricAttributeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// A structured log event containing typed payload data.
//
// The properties Payload, SpanID, Timestamp, TraceID, Type are required.
type EventStructuredLogParam struct {
	// The structured payload data for the log event
	Payload EventStructuredLogPayloadUnionParam `json:"payload,omitzero,required"`
	// Unique identifier for the span this event belongs to
	SpanID string `json:"span_id,required"`
	// Timestamp when the event occurred
	Timestamp time.Time `json:"timestamp,required" format:"date-time"`
	// Unique identifier for the trace this event belongs to
	TraceID string `json:"trace_id,required"`
	// (Optional) Key-value pairs containing additional metadata about the event
	Attributes map[string]EventStructuredLogAttributeUnionParam `json:"attributes,omitzero"`
	// Event type identifier set to STRUCTURED_LOG
	//
	// This field can be elided, and will marshal its zero value as "structured_log".
	Type constant.StructuredLog `json:"type,required"`
	paramObj
}

func (r EventStructuredLogParam) MarshalJSON() (data []byte, err error) {
	type shadow EventStructuredLogParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventStructuredLogParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EventStructuredLogPayloadUnionParam struct {
	OfSpanStart *EventStructuredLogPayloadSpanStartParam `json:",omitzero,inline"`
	OfSpanEnd   *EventStructuredLogPayloadSpanEndParam   `json:",omitzero,inline"`
	paramUnion
}

func (u EventStructuredLogPayloadUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSpanStart, u.OfSpanEnd)
}
func (u *EventStructuredLogPayloadUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EventStructuredLogPayloadUnionParam) asAny() any {
	if !param.IsOmitted(u.OfSpanStart) {
		return u.OfSpanStart
	} else if !param.IsOmitted(u.OfSpanEnd) {
		return u.OfSpanEnd
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventStructuredLogPayloadUnionParam) GetName() *string {
	if vt := u.OfSpanStart; vt != nil {
		return &vt.Name
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventStructuredLogPayloadUnionParam) GetParentSpanID() *string {
	if vt := u.OfSpanStart; vt != nil && vt.ParentSpanID.Valid() {
		return &vt.ParentSpanID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventStructuredLogPayloadUnionParam) GetStatus() *string {
	if vt := u.OfSpanEnd; vt != nil {
		return &vt.Status
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u EventStructuredLogPayloadUnionParam) GetType() *string {
	if vt := u.OfSpanStart; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSpanEnd; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[EventStructuredLogPayloadUnionParam](
		"type",
		apijson.Discriminator[EventStructuredLogPayloadSpanStartParam]("span_start"),
		apijson.Discriminator[EventStructuredLogPayloadSpanEndParam]("span_end"),
	)
}

// Payload for a span start event.
//
// The properties Name, Type are required.
type EventStructuredLogPayloadSpanStartParam struct {
	// Human-readable name describing the operation this span represents
	Name string `json:"name,required"`
	// (Optional) Unique identifier for the parent span, if this is a child span
	ParentSpanID param.Opt[string] `json:"parent_span_id,omitzero"`
	// Payload type identifier set to SPAN_START
	//
	// This field can be elided, and will marshal its zero value as "span_start".
	Type constant.SpanStart `json:"type,required"`
	paramObj
}

func (r EventStructuredLogPayloadSpanStartParam) MarshalJSON() (data []byte, err error) {
	type shadow EventStructuredLogPayloadSpanStartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventStructuredLogPayloadSpanStartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Payload for a span end event.
//
// The properties Status, Type are required.
type EventStructuredLogPayloadSpanEndParam struct {
	// The final status of the span indicating success or failure
	//
	// Any of "ok", "error".
	Status string `json:"status,omitzero,required"`
	// Payload type identifier set to SPAN_END
	//
	// This field can be elided, and will marshal its zero value as "span_end".
	Type constant.SpanEnd `json:"type,required"`
	paramObj
}

func (r EventStructuredLogPayloadSpanEndParam) MarshalJSON() (data []byte, err error) {
	type shadow EventStructuredLogPayloadSpanEndParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EventStructuredLogPayloadSpanEndParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[EventStructuredLogPayloadSpanEndParam](
		"status", "ok", "error",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EventStructuredLogAttributeUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u EventStructuredLogAttributeUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *EventStructuredLogAttributeUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EventStructuredLogAttributeUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type TelemetryLogEventParams struct {
	// The event to log.
	Event EventUnionParam `json:"event,omitzero,required"`
	// The time to live of the event.
	TtlSeconds int64 `json:"ttl_seconds,required"`
	paramObj
}

func (r TelemetryLogEventParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryLogEventParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryLogEventParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
