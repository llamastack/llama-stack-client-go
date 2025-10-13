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

func TestCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Completions.New(context.TODO(), llamastackclient.CompletionNewParams{
		Model: "model",
		Prompt: llamastackclient.CompletionNewParamsPromptUnion{
			OfString: llamastackclient.String("string"),
		},
		BestOf:           llamastackclient.Int(0),
		Echo:             llamastackclient.Bool(true),
		FrequencyPenalty: llamastackclient.Float(0),
		LogitBias: map[string]float64{
			"foo": 0,
		},
		Logprobs:        llamastackclient.Bool(true),
		MaxTokens:       llamastackclient.Int(0),
		N:               llamastackclient.Int(0),
		PresencePenalty: llamastackclient.Float(0),
		Seed:            llamastackclient.Int(0),
		Stop: llamastackclient.CompletionNewParamsStopUnion{
			OfString: llamastackclient.String("string"),
		},
		StreamOptions: map[string]llamastackclient.CompletionNewParamsStreamOptionUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		Suffix:      llamastackclient.String("suffix"),
		Temperature: llamastackclient.Float(0),
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
