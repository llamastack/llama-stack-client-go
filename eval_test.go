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

func TestEvalEvaluateRowsWithOptionalParams(t *testing.T) {
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
	_, err := client.Eval.EvaluateRows(
		context.TODO(),
		"benchmark_id",
		llamastackclient.EvalEvaluateRowsParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateUnionParam{
					OfModel: &llamastackclient.BenchmarkConfigEvalCandidateModelParam{
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
			InputRows: []map[string]llamastackclient.EvalEvaluateRowsParamsInputRowUnion{{
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

func TestEvalEvaluateRowsAlphaWithOptionalParams(t *testing.T) {
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
	_, err := client.Eval.EvaluateRowsAlpha(
		context.TODO(),
		"benchmark_id",
		llamastackclient.EvalEvaluateRowsAlphaParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateUnionParam{
					OfModel: &llamastackclient.BenchmarkConfigEvalCandidateModelParam{
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
			InputRows: []map[string]llamastackclient.EvalEvaluateRowsAlphaParamsInputRowUnion{{
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

func TestEvalRunEvalWithOptionalParams(t *testing.T) {
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
	_, err := client.Eval.RunEval(
		context.TODO(),
		"benchmark_id",
		llamastackclient.EvalRunEvalParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateUnionParam{
					OfModel: &llamastackclient.BenchmarkConfigEvalCandidateModelParam{
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

func TestEvalRunEvalAlphaWithOptionalParams(t *testing.T) {
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
	_, err := client.Eval.RunEvalAlpha(
		context.TODO(),
		"benchmark_id",
		llamastackclient.EvalRunEvalAlphaParams{
			BenchmarkConfig: llamastackclient.BenchmarkConfigParam{
				EvalCandidate: llamastackclient.BenchmarkConfigEvalCandidateUnionParam{
					OfModel: &llamastackclient.BenchmarkConfigEvalCandidateModelParam{
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
