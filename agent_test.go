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
	"github.com/stainless-sdks/llama-stack-client-go/shared"
)

func TestAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.New(context.TODO(), llamastackclient.AgentNewParams{
		AgentConfig: shared.AgentConfigParam{
			Instructions: "instructions",
			Model:        "model",
			ClientTools: []llamastackclient.ToolDefParam{{
				Name:        "name",
				Description: llamastackclient.String("description"),
				Metadata: map[string]llamastackclient.ToolDefMetadataUnionParam{
					"foo": {
						OfBool: llamastackclient.Bool(true),
					},
				},
				Parameters: []llamastackclient.ToolDefParameterParam{{
					Description:   "description",
					Name:          "name",
					ParameterType: "parameter_type",
					Required:      true,
					Default: llamastackclient.ToolDefParameterDefaultUnionParam{
						OfBool: llamastackclient.Bool(true),
					},
				}},
			}},
			EnableSessionPersistence: llamastackclient.Bool(true),
			InputShields:             []string{"string"},
			MaxInferIters:            llamastackclient.Int(0),
			Name:                     llamastackclient.String("name"),
			OutputShields:            []string{"string"},
			ResponseFormat: shared.ResponseFormatUnionParam{
				OfJsonSchema: &shared.ResponseFormatJsonSchemaParam{
					JsonSchema: map[string]shared.ResponseFormatJsonSchemaJsonSchemaUnionParam{
						"foo": {
							OfBool: llamastackclient.Bool(true),
						},
					},
				},
			},
			SamplingParams: shared.SamplingParams{
				Strategy: shared.SamplingParamsStrategyUnion{
					OfGreedy: &shared.SamplingParamsStrategyGreedy{},
				},
				MaxTokens:         llamastackclient.Int(0),
				RepetitionPenalty: llamastackclient.Float(0),
				Stop:              []string{"string"},
			},
			ToolChoice: shared.AgentConfigToolChoiceAuto,
			ToolConfig: shared.AgentConfigToolConfigParam{
				SystemMessageBehavior: "append",
				ToolChoice:            "auto",
				ToolPromptFormat:      "json",
			},
			ToolPromptFormat: shared.AgentConfigToolPromptFormatJson,
			Toolgroups: []shared.AgentConfigToolgroupUnionParam{{
				OfString: llamastackclient.String("string"),
			}},
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

func TestAgentGet(t *testing.T) {
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
	_, err := client.Agents.Get(context.TODO(), "agent_id")
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.List(context.TODO(), llamastackclient.AgentListParams{
		Limit:      llamastackclient.Int(0),
		StartIndex: llamastackclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentDelete(t *testing.T) {
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
	err := client.Agents.Delete(context.TODO(), "agent_id")
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
