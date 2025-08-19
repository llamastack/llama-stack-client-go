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
	"github.com/llamastack/llama-stack-client-go/shared"
)

func TestInferenceBatchChatCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.BatchChatCompletion(context.TODO(), llamastackclient.InferenceBatchChatCompletionParams{
		MessagesBatch: [][]shared.MessageUnionParam{{{
			OfUser: &shared.UserMessageParam{
				Content: shared.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
				Context: shared.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
			},
		}}},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceBatchChatCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: shared.ResponseFormatUnionParam{
			OfJsonSchema: &shared.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]shared.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: shared.SamplingParams{
			Strategy: shared.SamplingParamsStrategyUnion{
				OfGreedy: &shared.SamplingParamsStrategyGreedy{},
			},
			MaxTokens:         llamastackclient.Int(0),
			RepetitionPenalty: llamastackclient.Float(0),
			Stop:              []string{"string"},
		},
		ToolConfig: llamastackclient.InferenceBatchChatCompletionParamsToolConfig{
			SystemMessageBehavior: "append",
			ToolChoice:            "auto",
			ToolPromptFormat:      "json",
		},
		Tools: []llamastackclient.InferenceBatchChatCompletionParamsTool{{
			ToolName:    "brave_search",
			Description: llamastackclient.String("description"),
			Parameters: map[string]shared.ToolParamDefinition{
				"foo": {
					ParamType: "param_type",
					Default: shared.ToolParamDefinitionDefaultUnion{
						OfBool: llamastackclient.Bool(true),
					},
					Description: llamastackclient.String("description"),
					Required:    llamastackclient.Bool(true),
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

func TestInferenceBatchCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.BatchCompletion(context.TODO(), llamastackclient.InferenceBatchCompletionParams{
		ContentBatch: []shared.InterleavedContentUnionParam{{
			OfString: llamastackclient.String("string"),
		}},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceBatchCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: shared.ResponseFormatUnionParam{
			OfJsonSchema: &shared.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]shared.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: shared.SamplingParams{
			Strategy: shared.SamplingParamsStrategyUnion{
				OfGreedy: &shared.SamplingParamsStrategyGreedy{},
			},
			MaxTokens:         llamastackclient.Int(0),
			RepetitionPenalty: llamastackclient.Float(0),
			Stop:              []string{"string"},
		},
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

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
		Messages: []shared.MessageUnionParam{{
			OfUser: &shared.UserMessageParam{
				Content: shared.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
				Context: shared.InterleavedContentUnionParam{
					OfString: llamastackclient.String("string"),
				},
			},
		}},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceChatCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: shared.ResponseFormatUnionParam{
			OfJsonSchema: &shared.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]shared.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: shared.SamplingParams{
			Strategy: shared.SamplingParamsStrategyUnion{
				OfGreedy: &shared.SamplingParamsStrategyGreedy{},
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
			Parameters: map[string]shared.ToolParamDefinition{
				"foo": {
					ParamType: "param_type",
					Default: shared.ToolParamDefinitionDefaultUnion{
						OfBool: llamastackclient.Bool(true),
					},
					Description: llamastackclient.String("description"),
					Required:    llamastackclient.Bool(true),
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

func TestInferenceCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Inference.Completion(context.TODO(), llamastackclient.InferenceCompletionParams{
		Content: shared.InterleavedContentUnionParam{
			OfString: llamastackclient.String("string"),
		},
		ModelID: "model_id",
		Logprobs: llamastackclient.InferenceCompletionParamsLogprobs{
			TopK: llamastackclient.Int(0),
		},
		ResponseFormat: shared.ResponseFormatUnionParam{
			OfJsonSchema: &shared.ResponseFormatJsonSchemaParam{
				JsonSchema: map[string]shared.ResponseFormatJsonSchemaJsonSchemaUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
			},
		},
		SamplingParams: shared.SamplingParams{
			Strategy: shared.SamplingParamsStrategyUnion{
				OfGreedy: &shared.SamplingParamsStrategyGreedy{},
			},
			MaxTokens:         llamastackclient.Int(0),
			RepetitionPenalty: llamastackclient.Float(0),
			Stop:              []string{"string"},
		},
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
