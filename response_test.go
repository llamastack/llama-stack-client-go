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
		Model:              "model",
		Include:            []string{"string"},
		Instructions:       llamastackclient.String("instructions"),
		MaxInferIters:      llamastackclient.Int(0),
		PreviousResponseID: llamastackclient.String("previous_response_id"),
		Store:              llamastackclient.Bool(true),
		Temperature:        llamastackclient.Float(0),
		Text: llamastackclient.ResponseNewParamsText{
			Format: llamastackclient.ResponseNewParamsTextFormat{
				Type:        llamastackclient.ResponseNewParamsTextFormatTypeText,
				Description: llamastackclient.String("description"),
				Name:        llamastackclient.String("name"),
				Schema: map[string]llamastackclient.ResponseNewParamsTextFormatSchemaUnion{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
				Strict: llamastackclient.Bool(true),
			},
		},
		Tools: []llamastackclient.ResponseNewParamsToolUnion{{
			OfOpenAIResponseInputToolWebSearch: &llamastackclient.ResponseNewParamsToolOpenAIResponseInputToolWebSearch{
				Type:              llamastackclient.ResponseNewParamsToolOpenAIResponseInputToolWebSearchTypeWebSearch,
				SearchContextSize: llamastackclient.String("search_context_size"),
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
