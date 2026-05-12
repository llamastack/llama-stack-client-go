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

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Chat.Completions.New(context.TODO(), ogxclient.ChatCompletionNewParams{
		Messages: []ogxclient.ChatCompletionNewParamsMessageUnion{{
			OfUser: &ogxclient.ChatCompletionNewParamsMessageUser{
				Content: ogxclient.ChatCompletionNewParamsMessageUserContentUnion{
					OfString: ogxclient.String("string"),
				},
				Name: ogxclient.String("name"),
				Role: "user",
			},
		}},
		Model:            "model",
		FrequencyPenalty: ogxclient.Float(-2),
		FunctionCall: ogxclient.ChatCompletionNewParamsFunctionCallUnion{
			OfString: ogxclient.String("string"),
		},
		Functions: []map[string]any{{
			"foo": "bar",
		}},
		LogitBias: map[string]float64{
			"foo": 0,
		},
		Logprobs:            ogxclient.Bool(true),
		MaxCompletionTokens: ogxclient.Int(1),
		MaxTokens:           ogxclient.Int(1),
		N:                   ogxclient.Int(1),
		ParallelToolCalls:   ogxclient.Bool(true),
		PresencePenalty:     ogxclient.Float(-2),
		PromptCacheKey:      ogxclient.String("prompt_cache_key"),
		ReasoningEffort:     ogxclient.ChatCompletionNewParamsReasoningEffortNone,
		ResponseFormat: ogxclient.ChatCompletionNewParamsResponseFormatUnion{
			OfText: &ogxclient.ChatCompletionNewParamsResponseFormatText{
				Type: "text",
			},
		},
		Seed:        ogxclient.Int(0),
		ServiceTier: ogxclient.ChatCompletionNewParamsServiceTierAuto,
		Stop: ogxclient.ChatCompletionNewParamsStopUnion{
			OfString: ogxclient.String("string"),
		},
		StreamOptions: map[string]any{
			"foo": "bar",
		},
		Temperature: ogxclient.Float(0),
		ToolChoice: ogxclient.ChatCompletionNewParamsToolChoiceUnion{
			OfString: ogxclient.String("string"),
		},
		Tools: []map[string]any{{
			"foo": "bar",
		}},
		TopLogprobs: ogxclient.Int(0),
		TopP:        ogxclient.Float(0),
		User:        ogxclient.String("user"),
	})
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatCompletionGet(t *testing.T) {
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
	_, err := client.Chat.Completions.Get(context.TODO(), "completion_id")
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestChatCompletionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Chat.Completions.List(context.TODO(), ogxclient.ChatCompletionListParams{
		After: ogxclient.String("after"),
		Limit: ogxclient.Int(0),
		Model: ogxclient.String("model"),
		Order: ogxclient.ChatCompletionListParamsOrderAsc,
	})
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
