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
