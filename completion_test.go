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

func TestCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Completions.New(context.TODO(), ogxclient.CompletionNewParams{
		Model: "model",
		Prompt: ogxclient.CompletionNewParamsPromptUnion{
			OfString: ogxclient.String("string"),
		},
		BestOf:           ogxclient.Int(1),
		Echo:             ogxclient.Bool(true),
		FrequencyPenalty: ogxclient.Float(-2),
		LogitBias: map[string]float64{
			"foo": 0,
		},
		Logprobs:        ogxclient.Int(0),
		MaxTokens:       ogxclient.Int(1),
		N:               ogxclient.Int(1),
		PresencePenalty: ogxclient.Float(-2),
		Seed:            ogxclient.Int(0),
		Stop: ogxclient.CompletionNewParamsStopUnion{
			OfString: ogxclient.String("string"),
		},
		StreamOptions: map[string]any{
			"foo": "bar",
		},
		Suffix:      ogxclient.String("suffix"),
		Temperature: ogxclient.Float(0),
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
