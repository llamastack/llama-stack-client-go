// Copyright (c) The OGX Contributors.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ogxclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ogx-ai/ogx-client-go"
	"github.com/ogx-ai/ogx-client-go/internal/testutil"
	"github.com/ogx-ai/ogx-client-go/option"
)

func TestResponseNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ogxclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Responses.New(context.TODO(), ogxclient.ResponseNewParams{
		Input: ogxclient.ResponseNewParamsInputUnion{
			OfString: ogxclient.String("string"),
		},
		Model:      "model",
		Background: ogxclient.Bool(true),
		ContextManagement: []ogxclient.ResponseNewParamsContextManagement{{
			Type:             "compaction",
			CompactThreshold: ogxclient.Int(0),
		}},
		Conversation:     ogxclient.String("conversation"),
		FrequencyPenalty: ogxclient.Float(-2),
		Include:          []string{"web_search_call.action.sources"},
		Instructions:     ogxclient.String("instructions"),
		MaxInferIters:    ogxclient.Int(1),
		MaxOutputTokens:  ogxclient.Int(16),
		MaxToolCalls:     ogxclient.Int(1),
		Metadata: map[string]string{
			"foo": "string",
		},
		ParallelToolCalls:  ogxclient.Bool(true),
		PresencePenalty:    ogxclient.Float(-2),
		PreviousResponseID: ogxclient.String("previous_response_id"),
		Prompt: ogxclient.ResponseNewParamsPrompt{
			ID: "id",
			Variables: map[string]ogxclient.ResponseNewParamsPromptVariableUnion{
				"foo": {
					OfInputText: &ogxclient.ResponseNewParamsPromptVariableInputText{
						Text: "text",
						Type: "input_text",
					},
				},
			},
			Version: ogxclient.String("version"),
		},
		PromptCacheKey: ogxclient.String("prompt_cache_key"),
		Reasoning: ogxclient.ResponseNewParamsReasoning{
			Effort:          "none",
			GenerateSummary: "auto",
			Summary:         "auto",
		},
		ServiceTier: ogxclient.ResponseNewParamsServiceTierAuto,
		Store:       ogxclient.Bool(true),
		StreamOptions: ogxclient.ResponseNewParamsStreamOptions{
			IncludeObfuscation: ogxclient.Bool(true),
		},
		Temperature: ogxclient.Float(0),
		Text: ogxclient.ResponseNewParamsText{
			Format: ogxclient.ResponseNewParamsTextFormat{
				Description: ogxclient.String("description"),
				Name:        ogxclient.String("name"),
				Schema: map[string]any{
					"foo": "bar",
				},
				Strict: ogxclient.Bool(true),
				Type:   ogxclient.ResponseNewParamsTextFormatTypeText,
			},
			Verbosity: "low",
		},
		ToolChoice: ogxclient.ResponseNewParamsToolChoiceUnion{
			OfOpenAIResponseInputToolChoiceMode: ogxclient.String("auto"),
		},
		Tools: []ogxclient.ResponseNewParamsToolUnion{{
			OfOpenAIResponseInputToolWebSearch: &ogxclient.ResponseNewParamsToolOpenAIResponseInputToolWebSearch{
				SearchContextSize: ogxclient.String("S?oC\"high"),
				Type:              ogxclient.ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearch,
			},
		}},
		TopLogprobs: ogxclient.Int(0),
		TopP:        ogxclient.Float(0),
		Truncation:  ogxclient.ResponseNewParamsTruncationAuto,
	})
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponseGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ogxclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Responses.Get(context.TODO(), "response_id")
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponseListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ogxclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Responses.List(context.TODO(), ogxclient.ResponseListParams{
		After: ogxclient.String("after"),
		Limit: ogxclient.Int(0),
		Model: ogxclient.String("model"),
		Order: ogxclient.ResponseListParamsOrderAsc,
	})
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponseDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ogxclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Responses.Delete(context.TODO(), "response_id")
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestResponseCompactWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ogxclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Responses.Compact(context.TODO(), ogxclient.ResponseCompactParams{
		Model: "model",
		Input: ogxclient.ResponseCompactParamsInputUnion{
			OfString: ogxclient.String("string"),
		},
		Instructions:       ogxclient.String("instructions"),
		ParallelToolCalls:  ogxclient.Bool(true),
		PreviousResponseID: ogxclient.String("previous_response_id"),
		PromptCacheKey:     ogxclient.String("prompt_cache_key"),
		Reasoning: ogxclient.ResponseCompactParamsReasoning{
			Effort:          "none",
			GenerateSummary: "auto",
			Summary:         "auto",
		},
		Text: ogxclient.ResponseCompactParamsText{
			Format: ogxclient.ResponseCompactParamsTextFormat{
				Description: ogxclient.String("description"),
				Name:        ogxclient.String("name"),
				Schema: map[string]any{
					"foo": "bar",
				},
				Strict: ogxclient.Bool(true),
				Type:   ogxclient.ResponseCompactParamsTextFormatTypeText,
			},
			Verbosity: "low",
		},
		Tools: []ogxclient.ResponseCompactParamsToolUnion{{
			OfOpenAIResponseInputToolWebSearch: &ogxclient.ResponseCompactParamsToolOpenAIResponseInputToolWebSearch{
				SearchContextSize: ogxclient.String("S?oC\"high"),
				Type:              ogxclient.ResponseCompactParamsToolOpenAIResponseInputToolWebSearchTypeWebSearch,
			},
		}},
	})
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
