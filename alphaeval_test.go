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

func TestAlphaEvalEvaluateRowsWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Eval.EvaluateRows(
		context.TODO(),
		"benchmark_id",
		llamastackclient.AlphaEvalEvaluateRowsParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateParam{
					Model: "model",
					SamplingParams: llamastackclient.SamplingParams{
						Strategy: llamastackclient.SamplingParamsStrategyUnion{
							OfGreedy: &llamastackclient.SamplingParamsStrategyGreedy{},
						},
						MaxTokens:         llamastackclient.Int(0),
						RepetitionPenalty: llamastackclient.Float(0),
						Stop:              []string{"string"},
					},
					SystemMessage: llamastackclient.SystemMessageParam{
						Content: llamastackclient.InterleavedContentUnionParam{
							OfString: llamastackclient.String("string"),
						},
					},
				},
				ScoringParams: map[string]llamastackclient.ScoringFnParamsUnion{
					"foo": {
						OfLlmAsJudge: &llamastackclient.ScoringFnParamsLlmAsJudge{
							AggregationFunctions: []string{"average"},
							JudgeModel:           "judge_model",
							JudgeScoreRegexes:    []string{"string"},
							PromptTemplate:       llamastackclient.String("prompt_template"),
						},
					},
				},
				NumExamples: llamastackclient.Int(0),
			},
			InputRows: []map[string]llamastackclient.AlphaEvalEvaluateRowsParamsInputRowUnion{{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			}},
			ScoringFunctions: []string{"string"},
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlphaEvalEvaluateRowsAlphaWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Eval.EvaluateRowsAlpha(
		context.TODO(),
		"benchmark_id",
		llamastackclient.AlphaEvalEvaluateRowsAlphaParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateParam{
					Model: "model",
					SamplingParams: llamastackclient.SamplingParams{
						Strategy: llamastackclient.SamplingParamsStrategyUnion{
							OfGreedy: &llamastackclient.SamplingParamsStrategyGreedy{},
						},
						MaxTokens:         llamastackclient.Int(0),
						RepetitionPenalty: llamastackclient.Float(0),
						Stop:              []string{"string"},
					},
					SystemMessage: llamastackclient.SystemMessageParam{
						Content: llamastackclient.InterleavedContentUnionParam{
							OfString: llamastackclient.String("string"),
						},
					},
				},
				ScoringParams: map[string]llamastackclient.ScoringFnParamsUnion{
					"foo": {
						OfLlmAsJudge: &llamastackclient.ScoringFnParamsLlmAsJudge{
							AggregationFunctions: []string{"average"},
							JudgeModel:           "judge_model",
							JudgeScoreRegexes:    []string{"string"},
							PromptTemplate:       llamastackclient.String("prompt_template"),
						},
					},
				},
				NumExamples: llamastackclient.Int(0),
			},
			InputRows: []map[string]llamastackclient.AlphaEvalEvaluateRowsAlphaParamsInputRowUnion{{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			}},
			ScoringFunctions: []string{"string"},
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlphaEvalRunEvalWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Eval.RunEval(
		context.TODO(),
		"benchmark_id",
		llamastackclient.AlphaEvalRunEvalParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateParam{
					Model: "model",
					SamplingParams: llamastackclient.SamplingParams{
						Strategy: llamastackclient.SamplingParamsStrategyUnion{
							OfGreedy: &llamastackclient.SamplingParamsStrategyGreedy{},
						},
						MaxTokens:         llamastackclient.Int(0),
						RepetitionPenalty: llamastackclient.Float(0),
						Stop:              []string{"string"},
					},
					SystemMessage: llamastackclient.SystemMessageParam{
						Content: llamastackclient.InterleavedContentUnionParam{
							OfString: llamastackclient.String("string"),
						},
					},
				},
				ScoringParams: map[string]llamastackclient.ScoringFnParamsUnion{
					"foo": {
						OfLlmAsJudge: &llamastackclient.ScoringFnParamsLlmAsJudge{
							AggregationFunctions: []string{"average"},
							JudgeModel:           "judge_model",
							JudgeScoreRegexes:    []string{"string"},
							PromptTemplate:       llamastackclient.String("prompt_template"),
						},
					},
				},
				NumExamples: llamastackclient.Int(0),
			},
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAlphaEvalRunEvalAlphaWithOptionalParams(t *testing.T) {
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
	_, err := client.Alpha.Eval.RunEvalAlpha(
		context.TODO(),
		"benchmark_id",
		llamastackclient.AlphaEvalRunEvalAlphaParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateParam{
					Model: "model",
					SamplingParams: llamastackclient.SamplingParams{
						Strategy: llamastackclient.SamplingParamsStrategyUnion{
							OfGreedy: &llamastackclient.SamplingParamsStrategyGreedy{},
						},
						MaxTokens:         llamastackclient.Int(0),
						RepetitionPenalty: llamastackclient.Float(0),
						Stop:              []string{"string"},
					},
					SystemMessage: llamastackclient.SystemMessageParam{
						Content: llamastackclient.InterleavedContentUnionParam{
							OfString: llamastackclient.String("string"),
						},
					},
				},
				ScoringParams: map[string]llamastackclient.ScoringFnParamsUnion{
					"foo": {
						OfLlmAsJudge: &llamastackclient.ScoringFnParamsLlmAsJudge{
							AggregationFunctions: []string{"average"},
							JudgeModel:           "judge_model",
							JudgeScoreRegexes:    []string{"string"},
							PromptTemplate:       llamastackclient.String("prompt_template"),
						},
					},
				},
				NumExamples: llamastackclient.Int(0),
			},
		},
	)
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
