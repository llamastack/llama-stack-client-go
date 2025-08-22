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

func TestSafetyNew(t *testing.T) {
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
	_, err := client.Safety.New(context.TODO(), llamastackclient.SafetyNewParams{
		Input: llamastackclient.SafetyNewParamsInputUnion{
			OfString: llamastackclient.String("string"),
		},
		Model: "model",
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSafetyRunShield(t *testing.T) {
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
	_, err := client.Safety.RunShield(context.TODO(), llamastackclient.SafetyRunShieldParams{
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
		Params: map[string]llamastackclient.SafetyRunShieldParamsParamUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		ShieldID: "shield_id",
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
