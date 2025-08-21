// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/llamastack/llama-stack-client-go"
	"github.com/llamastack/llama-stack-client-go/internal/testutil"
	"github.com/llamastack/llama-stack-client-go/option"
)

func TestToolRuntimeInvokeTool(t *testing.T) {
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
	_, err := client.ToolRuntime.InvokeTool(context.TODO(), llamastackclient.ToolRuntimeInvokeToolParams{
		Kwargs: map[string]llamastackclient.ToolRuntimeInvokeToolParamsKwargUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		ToolName: "tool_name",
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolRuntimeListToolsWithOptionalParams(t *testing.T) {
	t.Skip("invalid query params unsupported by prism")
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
	_, err := client.ToolRuntime.ListTools(context.TODO(), llamastackclient.ToolRuntimeListToolsParams{
		McpEndpoint: llamastackclient.ToolRuntimeListToolsParamsMcpEndpoint{
			Uri: "uri",
		},
		ToolGroupID: llamastackclient.String("tool_group_id"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
