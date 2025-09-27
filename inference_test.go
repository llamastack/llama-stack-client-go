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

func TestInferenceChatCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.ChatCompletion(context.TODO(), llamastackclient.InferenceChatCompletionParams{
		Messages: []llamastackclient.MessageUnionParam{{
			OfUser: &llamastackclient.UserMessageParam{
				Content: llamastackclient.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
				Context: llamastackclient.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
			},
		}},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceChatCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: llamastackclient.ResponseFormatUnionParam{
			OfJsonSchema: &llamastackclient.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]llamastackclient.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: llamastackclient.SamplingParams{
			Strategy: llamastackclient.SamplingParamsStrategyUnion{
				OfGreedy: &llamastackclient.SamplingParamsStrategyGreedy{},
			},
			MaxTokens:         llamastackclient.Int(0),
			RepetitionPenalty: llamastackclient.Float(0),
			Stop:              []string{"string"},
		},
		ToolChoice: llamastackclient.InferenceChatCompletionParamsToolChoiceAuto,
		ToolConfig: llamastackclient.InferenceChatCompletionParamsToolConfig{
			SystemMessageBehavior: "append",
			ToolChoice:            "auto",
			ToolPromptFormat:      "json",
		},
		ToolPromptFormat: llamastackclient.InferenceChatCompletionParamsToolPromptFormatJson,
		Tools: []llamastackclient.InferenceChatCompletionParamsTool{{
			ToolName:    "brave_search",
			Description: llamastackclient.String("description"),
			Parameters: map[string]llamastackclient.ToolParamDefinition{
				"foo": {
					ParamType: "param_type",
					Default: llamastackclient.ToolParamDefinitionDefaultUnion{
						OfBool: llamastackclient.Bool(true),
					},
					Description: llamastackclient.String("description"),
					Items: llamastackclient.ToolParamDefinitionItemsUnion{
						OfBool: llamastackclient.Bool(true),
					},
					Required: llamastackclient.Bool(true),
					Title:    llamastackclient.String("title"),
				},
			},
		}},
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceEmbeddingsWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.Embeddings(context.TODO(), llamastackclient.InferenceEmbeddingsParams{
		Contents: llamastackclient.InferenceEmbeddingsParamsContentsUnion{
			OfStringArray: []string{"string"},
		},
		ModelID:         "model_id",
		OutputDimension: llamastackclient.Int(0),
		TaskType:        llamastackclient.InferenceEmbeddingsParamsTaskTypeQuery,
		TextTruncation:  llamastackclient.InferenceEmbeddingsParamsTextTruncationNone,
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestInferenceRerankWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.Rerank(context.TODO(), llamastackclient.InferenceRerankParams{
		Items: []llamastackclient.InferenceRerankParamsItemUnion{{
			OfString: llamastackclient.String("string"),
		}},
		Model: "model",
		Query: llamastackclient.InferenceRerankParamsQueryUnion{
			OfString: llamastackclient.String("string"),
		},
		MaxNumResults: llamastackclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
