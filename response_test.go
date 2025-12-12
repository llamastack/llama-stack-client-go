// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
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

func TestResponseNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Responses.New(context.TODO(), llamastackclient.ResponseNewParams{
		Input: llamastackclient.ResponseNewParamsInputUnion{
			OfString: llamastackclient.String("string"),
		},
		Model:         "model",
		Conversation:  llamastackclient.String("conversation"),
		Include:       []string{"web_search_call.action.sources"},
		Instructions:  llamastackclient.String("instructions"),
		MaxInferIters: llamastackclient.Int(0),
		MaxToolCalls:  llamastackclient.Int(0),
		Metadata: map[string]string{
			"foo": "string",
		},
		ParallelToolCalls:  llamastackclient.Bool(true),
		PreviousResponseID: llamastackclient.String("previous_response_id"),
		Prompt: llamastackclient.ResponseNewParamsPrompt{
			ID: "id",
			Variables: map[string]llamastackclient.ResponseNewParamsPromptVariableUnion{
				"foo": {
					OfInputText: &llamastackclient.ResponseNewParamsPromptVariableInputText{
						Text: "text",
						Type: "input_text",
					},
				},
			},
			Version: llamastackclient.String("version"),
		},
		Store:       llamastackclient.Bool(true),
		Temperature: llamastackclient.Float(0),
		Text: llamastackclient.ResponseNewParamsText{
			Format: llamastackclient.ResponseNewParamsTextFormat{
				Description: llamastackclient.String("description"),
				Name:        llamastackclient.String("name"),
				Schema: map[string]any{
					"foo": "bar",
				},
				Strict: llamastackclient.Bool(true),
				Type:   "text",
			},
		},
		Tools: []llamastackclient.ResponseNewParamsToolUnion{{
			OfOpenAIResponseInputToolWebSearch: &llamastackclient.ResponseNewParamsToolOpenAIResponseInputToolWebSearch{
				SearchContextSize: llamastackclient.String("S?oC\"high"),
				Type:              "web_search",
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

func TestResponseGet(t *testing.T) {
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
	_, err := client.Responses.Get(context.TODO(), "response_id")
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Responses.List(context.TODO(), llamastackclient.ResponseListParams{
		After: llamastackclient.String("after"),
		Limit: llamastackclient.Int(0),
		Model: llamastackclient.String("model"),
		Order: llamastackclient.ResponseListParamsOrderAsc,
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Responses.Delete(context.TODO(), "response_id")
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
