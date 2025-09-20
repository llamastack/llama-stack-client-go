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

// Get a span by its ID.
func (r *TelemetryService) GetSpan(ctx context.Context, spanID string, query TelemetryGetSpanParams, opts ...option.RequestOption) (res *TelemetryGetSpanResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.TraceID == "" {
		err = errors.New("missing required trace_id parameter")
		return
	}
	if spanID == "" {
		err = errors.New("missing required span_id parameter")
		return
	}
	path := fmt.Sprintf("v1/telemetry/traces/%s/spans/%s", query.TraceID, spanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a span tree by its ID.
func (r *TelemetryService) GetSpanTree(ctx context.Context, spanID string, body TelemetryGetSpanTreeParams, opts ...option.RequestOption) (res *TelemetryGetSpanTreeResponse, err error) {
	var env TelemetryGetSpanTreeResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if spanID == "" {
		err = errors.New("missing required span_id parameter")
		return
	}
	path := fmt.Sprintf("v1/telemetry/spans/%s/tree", spanID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a trace by its ID.
func (r *TelemetryService) GetTrace(ctx context.Context, traceID string, opts ...option.RequestOption) (res *Trace, err error) {
	opts = slices.Concat(r.Options, opts)
	if traceID == "" {
		err = errors.New("missing required trace_id parameter")
		return
	}
	path := fmt.Sprintf("v1/telemetry/traces/%s", traceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Query metrics.
func (r *TelemetryService) QueryMetrics(ctx context.Context, metricName string, body TelemetryQueryMetricsParams, opts ...option.RequestOption) (res *[]TelemetryQueryMetricsResponse, err error) {
	var env TelemetryQueryMetricsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if metricName == "" {
		err = errors.New("missing required metric_name parameter")
		return
	}
	path := fmt.Sprintf("v1/telemetry/metrics/%s", metricName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Query spans.
func (r *TelemetryService) QuerySpans(ctx context.Context, body TelemetryQuerySpansParams, opts ...option.RequestOption) (res *[]QuerySpansResponseData, err error) {
	var env QuerySpansResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1/telemetry/spans"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Query traces.
func (r *TelemetryService) QueryTraces(ctx context.Context, body TelemetryQueryTracesParams, opts ...option.RequestOption) (res *[]Trace, err error) {
	var env TelemetryQueryTracesResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v1/telemetry/traces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Save spans to a dataset.
func (r *TelemetryService) SaveSpansToDataset(ctx context.Context, body TelemetrySaveSpansToDatasetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/telemetry/spans/export"
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

// A condition for filtering query results.
//
// The properties Key, Op, Value are required.
type QueryConditionParam struct {
	// The value to compare against
	Value QueryConditionValueUnionParam `json:"value,omitzero,required"`
	// The attribute key to filter on
	Key string `json:"key,required"`
	// The comparison operator to apply
	//
	// Any of "eq", "ne", "gt", "lt".
	Op QueryConditionOp `json:"op,omitzero,required"`
	paramObj
}

func (r QueryConditionParam) MarshalJSON() (data []byte, err error) {
	type shadow QueryConditionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryConditionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The comparison operator to apply
type QueryConditionOp string

const (
	QueryConditionOpEq QueryConditionOp = "eq"
	QueryConditionOpNe QueryConditionOp = "ne"
	QueryConditionOpGt QueryConditionOp = "gt"
	QueryConditionOpLt QueryConditionOp = "lt"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryConditionValueUnionParam struct {
	OfBool     param.Opt[bool]    `json:",omitzero,inline"`
	OfFloat    param.Opt[float64] `json:",omitzero,inline"`
	OfString   param.Opt[string]  `json:",omitzero,inline"`
	OfAnyArray []any              `json:",omitzero,inline"`
	paramUnion
}

func (u QueryConditionValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfAnyArray)
}
func (u *QueryConditionValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryConditionValueUnionParam) asAny() any {
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

// Response containing a list of spans.
type QuerySpansResponse struct {
	// List of spans matching the query criteria
	Data []QuerySpansResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySpansResponse) RawJSON() string { return r.JSON.raw }
func (r *QuerySpansResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A span representing a single operation within a trace.
type QuerySpansResponseData struct {
	// Human-readable name describing the operation this span represents
	Name string `json:"name,required"`
	// Unique identifier for the span
	SpanID string `json:"span_id,required"`
	// Timestamp when the operation began
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Unique identifier for the trace this span belongs to
	TraceID string `json:"trace_id,required"`
	// (Optional) Key-value pairs containing additional metadata about the span
	Attributes map[string]QuerySpansResponseDataAttributeUnion `json:"attributes"`
	// (Optional) Timestamp when the operation finished, if completed
	EndTime time.Time `json:"end_time" format:"date-time"`
	// (Optional) Unique identifier for the parent span, if this is a child span
	ParentSpanID string `json:"parent_span_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		SpanID       respjson.Field
		StartTime    respjson.Field
		TraceID      respjson.Field
		Attributes   respjson.Field
		EndTime      respjson.Field
		ParentSpanID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QuerySpansResponseData) RawJSON() string { return r.JSON.raw }
func (r *QuerySpansResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QuerySpansResponseDataAttributeUnion contains all possible properties and values
// from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type QuerySpansResponseDataAttributeUnion struct {
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

func (u QuerySpansResponseDataAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuerySpansResponseDataAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuerySpansResponseDataAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuerySpansResponseDataAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QuerySpansResponseDataAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *QuerySpansResponseDataAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A span that includes status information.
type SpanWithStatus struct {
	// Human-readable name describing the operation this span represents
	Name string `json:"name,required"`
	// Unique identifier for the span
	SpanID string `json:"span_id,required"`
	// Timestamp when the operation began
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Unique identifier for the trace this span belongs to
	TraceID string `json:"trace_id,required"`
	// (Optional) Key-value pairs containing additional metadata about the span
	Attributes map[string]SpanWithStatusAttributeUnion `json:"attributes"`
	// (Optional) Timestamp when the operation finished, if completed
	EndTime time.Time `json:"end_time" format:"date-time"`
	// (Optional) Unique identifier for the parent span, if this is a child span
	ParentSpanID string `json:"parent_span_id"`
	// (Optional) The current status of the span
	//
	// Any of "ok", "error".
	Status SpanWithStatusStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		SpanID       respjson.Field
		StartTime    respjson.Field
		TraceID      respjson.Field
		Attributes   respjson.Field
		EndTime      respjson.Field
		ParentSpanID respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpanWithStatus) RawJSON() string { return r.JSON.raw }
func (r *SpanWithStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SpanWithStatusAttributeUnion contains all possible properties and values from
// [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type SpanWithStatusAttributeUnion struct {
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

func (u SpanWithStatusAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SpanWithStatusAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SpanWithStatusAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SpanWithStatusAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SpanWithStatusAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *SpanWithStatusAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// (Optional) The current status of the span
type SpanWithStatusStatus string

const (
	SpanWithStatusStatusOk    SpanWithStatusStatus = "ok"
	SpanWithStatusStatusError SpanWithStatusStatus = "error"
)

// A trace representing the complete execution path of a request across multiple
// operations.
type Trace struct {
	// Unique identifier for the root span that started this trace
	RootSpanID string `json:"root_span_id,required"`
	// Timestamp when the trace began
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Unique identifier for the trace
	TraceID string `json:"trace_id,required"`
	// (Optional) Timestamp when the trace finished, if completed
	EndTime time.Time `json:"end_time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RootSpanID  respjson.Field
		StartTime   respjson.Field
		TraceID     respjson.Field
		EndTime     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Trace) RawJSON() string { return r.JSON.raw }
func (r *Trace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A span representing a single operation within a trace.
type TelemetryGetSpanResponse struct {
	// Human-readable name describing the operation this span represents
	Name string `json:"name,required"`
	// Unique identifier for the span
	SpanID string `json:"span_id,required"`
	// Timestamp when the operation began
	StartTime time.Time `json:"start_time,required" format:"date-time"`
	// Unique identifier for the trace this span belongs to
	TraceID string `json:"trace_id,required"`
	// (Optional) Key-value pairs containing additional metadata about the span
	Attributes map[string]TelemetryGetSpanResponseAttributeUnion `json:"attributes"`
	// (Optional) Timestamp when the operation finished, if completed
	EndTime time.Time `json:"end_time" format:"date-time"`
	// (Optional) Unique identifier for the parent span, if this is a child span
	ParentSpanID string `json:"parent_span_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		SpanID       respjson.Field
		StartTime    respjson.Field
		TraceID      respjson.Field
		Attributes   respjson.Field
		EndTime      respjson.Field
		ParentSpanID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryGetSpanResponse) RawJSON() string { return r.JSON.raw }
func (r *TelemetryGetSpanResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TelemetryGetSpanResponseAttributeUnion contains all possible properties and
// values from [bool], [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type TelemetryGetSpanResponseAttributeUnion struct {
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

func (u TelemetryGetSpanResponseAttributeUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TelemetryGetSpanResponseAttributeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TelemetryGetSpanResponseAttributeUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TelemetryGetSpanResponseAttributeUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TelemetryGetSpanResponseAttributeUnion) RawJSON() string { return u.JSON.raw }

func (r *TelemetryGetSpanResponseAttributeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryGetSpanTreeResponse map[string]SpanWithStatus

// A time series of metric data points.
type TelemetryQueryMetricsResponse struct {
	// List of labels associated with this metric series
	Labels []TelemetryQueryMetricsResponseLabel `json:"labels,required"`
	// The name of the metric
	Metric string `json:"metric,required"`
	// List of data points in chronological order
	Values []TelemetryQueryMetricsResponseValue `json:"values,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Labels      respjson.Field
		Metric      respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A label associated with a metric.
type TelemetryQueryMetricsResponseLabel struct {
	// The name of the label
	Name string `json:"name,required"`
	// The value of the label
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryMetricsResponseLabel) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryMetricsResponseLabel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single data point in a metric time series.
type TelemetryQueryMetricsResponseValue struct {
	// Unix timestamp when the metric value was recorded
	Timestamp int64  `json:"timestamp,required"`
	Unit      string `json:"unit,required"`
	// The numeric value of the metric at this timestamp
	Value float64 `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Timestamp   respjson.Field
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryMetricsResponseValue) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryMetricsResponseValue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryGetSpanParams struct {
	TraceID string `path:"trace_id,required" json:"-"`
	paramObj
}

type TelemetryGetSpanTreeParams struct {
	// The maximum depth of the tree.
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	// The attributes to return in the tree.
	AttributesToReturn []string `json:"attributes_to_return,omitzero"`
	paramObj
}

func (r TelemetryGetSpanTreeParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryGetSpanTreeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryGetSpanTreeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a tree structure of spans.
type TelemetryGetSpanTreeResponseEnvelope struct {
	// Dictionary mapping span IDs to spans with status information
	Data TelemetryGetSpanTreeResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryGetSpanTreeResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TelemetryGetSpanTreeResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type TelemetryQueryMetricsParams struct {
	// The type of query to perform.
	//
	// Any of "range", "instant".
	QueryType TelemetryQueryMetricsParamsQueryType `json:"query_type,omitzero,required"`
	// The start time of the metric to query.
	StartTime int64 `json:"start_time,required"`
	// The end time of the metric to query.
	EndTime param.Opt[int64] `json:"end_time,omitzero"`
	// The granularity of the metric to query.
	Granularity param.Opt[string] `json:"granularity,omitzero"`
	// The label matchers to apply to the metric.
	LabelMatchers []TelemetryQueryMetricsParamsLabelMatcher `json:"label_matchers,omitzero"`
	paramObj
}

func (r TelemetryQueryMetricsParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryQueryMetricsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryQueryMetricsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of query to perform.
type TelemetryQueryMetricsParamsQueryType string

const (
	TelemetryQueryMetricsParamsQueryTypeRange   TelemetryQueryMetricsParamsQueryType = "range"
	TelemetryQueryMetricsParamsQueryTypeInstant TelemetryQueryMetricsParamsQueryType = "instant"
)

// A matcher for filtering metrics by label values.
//
// The properties Name, Operator, Value are required.
type TelemetryQueryMetricsParamsLabelMatcher struct {
	// The name of the label to match
	Name string `json:"name,required"`
	// The comparison operator to use for matching
	//
	// Any of "=", "!=", "=~", "!~".
	Operator string `json:"operator,omitzero,required"`
	// The value to match against
	Value string `json:"value,required"`
	paramObj
}

func (r TelemetryQueryMetricsParamsLabelMatcher) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryQueryMetricsParamsLabelMatcher
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryQueryMetricsParamsLabelMatcher) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TelemetryQueryMetricsParamsLabelMatcher](
		"operator", "=", "!=", "=~", "!~",
	)
}

// Response containing metric time series data.
type TelemetryQueryMetricsResponseEnvelope struct {
	// List of metric series matching the query criteria
	Data []TelemetryQueryMetricsResponse `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryMetricsResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryMetricsResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryQuerySpansParams struct {
	// The attribute filters to apply to the spans.
	AttributeFilters []QueryConditionParam `json:"attribute_filters,omitzero,required"`
	// The attributes to return in the spans.
	AttributesToReturn []string `json:"attributes_to_return,omitzero,required"`
	// The maximum depth of the tree.
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	paramObj
}

func (r TelemetryQuerySpansParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryQuerySpansParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryQuerySpansParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetryQueryTracesParams struct {
	// The limit of traces to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// The offset of the traces to return.
	Offset param.Opt[int64] `json:"offset,omitzero"`
	// The attribute filters to apply to the traces.
	AttributeFilters []QueryConditionParam `json:"attribute_filters,omitzero"`
	// The order by of the traces to return.
	OrderBy []string `json:"order_by,omitzero"`
	paramObj
}

func (r TelemetryQueryTracesParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetryQueryTracesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetryQueryTracesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a list of traces.
type TelemetryQueryTracesResponseEnvelope struct {
	// List of traces matching the query criteria
	Data []Trace `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TelemetryQueryTracesResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TelemetryQueryTracesResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelemetrySaveSpansToDatasetParams struct {
	// The attribute filters to apply to the spans.
	AttributeFilters []QueryConditionParam `json:"attribute_filters,omitzero,required"`
	// The attributes to save to the dataset.
	AttributesToSave []string `json:"attributes_to_save,omitzero,required"`
	// The ID of the dataset to save the spans to.
	DatasetID string `json:"dataset_id,required"`
	// The maximum depth of the tree.
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	paramObj
}

func (r TelemetrySaveSpansToDatasetParams) MarshalJSON() (data []byte, err error) {
	type shadow TelemetrySaveSpansToDatasetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TelemetrySaveSpansToDatasetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
