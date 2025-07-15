// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/llama-stack-client-go"
	"github.com/stainless-sdks/llama-stack-client-go/internal/testutil"
	"github.com/stainless-sdks/llama-stack-client-go/option"
)

func TestResponseInputItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.Responses.InputItems.List(
		context.TODO(),
		"response_id",
		llamastackclient.ResponseInputItemListParams{
			After:   llamastackclient.String("after"),
			Before:  llamastackclient.String("before"),
			Include: []string{"string"},
			Limit:   llamastackclient.Int(0),
			Order:   llamastackclient.ResponseInputItemListParamsOrderAsc,
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
