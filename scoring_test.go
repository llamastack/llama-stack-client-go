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

func TestScoringScore(t *testing.T) {
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
	_, err := client.Scoring.Score(context.TODO(), llamastackclient.ScoringScoreParams{
		InputRows: []map[string]any{{
			"foo": "bar",
		}},
		ScoringFunctions: map[string]llamastackclient.ScoringScoreParamsScoringFunctionUnion{
			"foo": {
				OfLlmAsJudge: &llamastackclient.ScoringScoreParamsScoringFunctionLlmAsJudge{
					JudgeModel:           "judge_model",
					AggregationFunctions: []string{"average"},
					JudgeScoreRegexes:    []string{"string"},
					PromptTemplate:       llamastackclient.String("prompt_template"),
					Type:                 "llm_as_judge",
				},
			},
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

func TestScoringScoreBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Scoring.ScoreBatch(context.TODO(), llamastackclient.ScoringScoreBatchParams{
		DatasetID: "dataset_id",
		ScoringFunctions: map[string]llamastackclient.ScoringScoreBatchParamsScoringFunctionUnion{
			"foo": {
				OfLlmAsJudge: &llamastackclient.ScoringScoreBatchParamsScoringFunctionLlmAsJudge{
					JudgeModel:           "judge_model",
					AggregationFunctions: []string{"average"},
					JudgeScoreRegexes:    []string{"string"},
					PromptTemplate:       llamastackclient.String("prompt_template"),
					Type:                 "llm_as_judge",
				},
			},
		},
		SaveResultsDataset: llamastackclient.Bool(true),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
