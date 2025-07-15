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

func TestChatCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Chat.Completions.New(context.TODO(), llamastackclient.ChatCompletionNewParams{
		Messages: []llamastackclient.ChatCompletionNewParamsMessageUnion{{
			OfUser: &llamastackclient.ChatCompletionNewParamsMessageUser{
				Content: llamastackclient.ChatCompletionNewParamsMessageUserContentUnion{
					OfString: llamastackclient.String("string"),
				},
				Name: llamastackclient.String("name"),
			},
		}},
		Model:            "model",
		FrequencyPenalty: llamastackclient.Float(0),
		FunctionCall: llamastackclient.ChatCompletionNewParamsFunctionCallUnion{
			OfString: llamastackclient.String("string"),
		},
		Functions: []map[string]llamastackclient.ChatCompletionNewParamsFunctionUnion{{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		}},
		LogitBias: map[string]float64{
			"foo": 0,
		},
		Logprobs:            llamastackclient.Bool(true),
		MaxCompletionTokens: llamastackclient.Int(0),
		MaxTokens:           llamastackclient.Int(0),
		N:                   llamastackclient.Int(0),
		ParallelToolCalls:   llamastackclient.Bool(true),
		PresencePenalty:     llamastackclient.Float(0),
		ResponseFormat: llamastackclient.ChatCompletionNewParamsResponseFormatUnion{
			OfText: &llamastackclient.ChatCompletionNewParamsResponseFormatText{},
		},
		Seed: llamastackclient.Int(0),
		Stop: llamastackclient.ChatCompletionNewParamsStopUnion{
			OfString: llamastackclient.String("string"),
		},
		StreamOptions: map[string]llamastackclient.ChatCompletionNewParamsStreamOptionUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		Temperature: llamastackclient.Float(0),
		ToolChoice: llamastackclient.ChatCompletionNewParamsToolChoiceUnion{
			OfString: llamastackclient.String("string"),
		},
		Tools: []map[string]llamastackclient.ChatCompletionNewParamsToolUnion{{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		}},
		TopLogprobs: llamastackclient.Int(0),
		TopP:        llamastackclient.Float(0),
		User:        llamastackclient.String("user"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Chat.Completions.Get(context.TODO(), "completion_id")
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.Chat.Completions.List(context.TODO(), llamastackclient.ChatCompletionListParams{
		After: llamastackclient.String("after"),
		Limit: llamastackclient.Int(0),
		Model: llamastackclient.String("model"),
		Order: llamastackclient.ChatCompletionListParamsOrderAsc,
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
