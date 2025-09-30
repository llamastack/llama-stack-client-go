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

func TestAlphaAgentTurnNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Agents.Turn.New(
		context.TODO(),
		"session_id",
		llamastackclient.AlphaAgentTurnNewParams{
			AgentID: "agent_id",
			Messages: []llamastackclient.AlphaAgentTurnNewParamsMessageUnion{{
				OfUserMessage: &llamastackclient.UserMessageParam{
					Content: llamastackclient.InterleavedContentUnionParam{
						OfString: llamastackclient.String("string"),
					},
					Context: llamastackclient.InterleavedContentUnionParam{
						OfString: llamastackclient.String("string"),
					},
				},
			}},
			Documents: []llamastackclient.AlphaAgentTurnNewParamsDocument{{
				Content: llamastackclient.AlphaAgentTurnNewParamsDocumentContentUnion{
					OfString: llamastackclient.String("string"),
				},
				MimeType: "mime_type",
			}},
			ToolConfig: llamastackclient.AlphaAgentTurnNewParamsToolConfig{
				SystemMessageBehavior: "append",
				ToolChoice:            "auto",
				ToolPromptFormat:      "json",
			},
			Toolgroups: []llamastackclient.AlphaAgentTurnNewParamsToolgroupUnion{{
				OfString: llamastackclient.String("string"),
			}},
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

func TestAlphaAgentTurnGet(t *testing.T) {
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
	_, err := client.Alpha.Agents.Turn.Get(
		context.TODO(),
		"turn_id",
		llamastackclient.AlphaAgentTurnGetParams{
			AgentID:   "agent_id",
			SessionID: "session_id",
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

func TestAlphaAgentTurnResumeWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Agents.Turn.Resume(
		context.TODO(),
		"turn_id",
		llamastackclient.AlphaAgentTurnResumeParams{
			AgentID:   "agent_id",
			SessionID: "session_id",
			ToolResponses: []llamastackclient.ToolResponseParam{{
				CallID: "call_id",
				Content: llamastackclient.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
				ToolName: llamastackclient.ToolResponseToolNameBraveSearch,
				Metadata: map[string]llamastackclient.ToolResponseMetadataUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
			}},
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
