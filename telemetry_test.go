// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/llamastack/llama-stack-client-go"
	"github.com/llamastack/llama-stack-client-go/internal/testutil"
	"github.com/llamastack/llama-stack-client-go/option"
)

func TestTelemetryGetSpan(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Telemetry.GetSpan(
		context.TODO(),
		"span_id",
		llamastackclient.TelemetryGetSpanParams{
			TraceID: "trace_id",
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetryGetSpanTreeWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Telemetry.GetSpanTree(
		context.TODO(),
		"span_id",
		llamastackclient.TelemetryGetSpanTreeParams{
			AttributesToReturn: []string{"string"},
			MaxDepth:           llamastackclient.Int(0),
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetryGetTrace(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Telemetry.GetTrace(context.TODO(), "trace_id")
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetryLogEventWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Telemetry.LogEvent(context.TODO(), llamastackclient.TelemetryLogEventParams{
		Event: llamastackclient.EventUnionParam{
			OfUnstructuredLog: &llamastackclient.EventUnstructuredLogParam{
				Message:   "message",
				Severity:  "verbose",
				SpanID:    "span_id",
				Timestamp: time.Now(),
				TraceID:   "trace_id",
				Attributes: map[string]llamastackclient.EventUnstructuredLogAttributeUnionParam{
					"foo": {
						OfString: llamastackclient.String("string"),
					},
				},
			},
		},
		TtlSeconds: 0,
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetryQuerySpansWithOptionalParams(t *testing.T) {
	t.Skip("unsupported query params in java / kotlin")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Telemetry.QuerySpans(context.TODO(), llamastackclient.TelemetryQuerySpansParams{
		AttributeFilters: []llamastackclient.QueryConditionParam{{
			Key: "key",
			Op:  llamastackclient.QueryConditionOpEq,
			Value: llamastackclient.QueryConditionValueUnionParam{
				OfBool: llamastackclient.Bool(true),
			},
		}},
		AttributesToReturn: []string{"string"},
		MaxDepth:           llamastackclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetryQueryTracesWithOptionalParams(t *testing.T) {
	t.Skip("unsupported query params in java / kotlin")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Telemetry.QueryTraces(context.TODO(), llamastackclient.TelemetryQueryTracesParams{
		AttributeFilters: []llamastackclient.QueryConditionParam{{
			Key: "key",
			Op:  llamastackclient.QueryConditionOpEq,
			Value: llamastackclient.QueryConditionValueUnionParam{
				OfBool: llamastackclient.Bool(true),
			},
		}},
		Limit:   llamastackclient.Int(0),
		Offset:  llamastackclient.Int(0),
		OrderBy: []string{"string"},
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTelemetrySaveSpansToDatasetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	err := client.Telemetry.SaveSpansToDataset(context.TODO(), llamastackclient.TelemetrySaveSpansToDatasetParams{
		AttributeFilters: []llamastackclient.QueryConditionParam{{
			Key: "key",
			Op:  llamastackclient.QueryConditionOpEq,
			Value: llamastackclient.QueryConditionValueUnionParam{
				OfBool: llamastackclient.Bool(true),
			},
		}},
		AttributesToSave: []string{"string"},
		DatasetID:        "dataset_id",
		MaxDepth:         llamastackclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
