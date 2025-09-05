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

func TestToolRuntimeRagToolInsert(t *testing.T) {
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
	err := client.ToolRuntime.RagTool.Insert(context.TODO(), llamastackclient.ToolRuntimeRagToolInsertParams{
		ChunkSizeInTokens: 0,
		Documents: []llamastackclient.DocumentParam{{
			Content: llamastackclient.DocumentContentUnionParam{
				OfString: llamastackclient.String("string"),
			},
			DocumentID: "document_id",
			Metadata: map[string]llamastackclient.DocumentMetadataUnionParam{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			MimeType: llamastackclient.String("mime_type"),
		}},
		VectorDBID: "vector_db_id",
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestToolRuntimeRagToolQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.ToolRuntime.RagTool.Query(context.TODO(), llamastackclient.ToolRuntimeRagToolQueryParams{
		Content: llamastackclient.InterleavedContentUnionParam{
			OfString: llamastackclient.String("string"),
		},
		VectorDBIDs: []string{"string"},
		QueryConfig: llamastackclient.QueryConfigParam{
			ChunkTemplate:      "chunk_template",
			MaxChunks:          0,
			MaxTokensInContext: 0,
			QueryGeneratorConfig: llamastackclient.QueryGeneratorConfigUnionParam{
				OfDefault: &llamastackclient.QueryGeneratorConfigDefaultParam{
					Separator: "separator",
				},
			},
			Mode: llamastackclient.QueryConfigModeVector,
			Ranker: llamastackclient.QueryConfigRankerUnionParam{
				OfRrf: &llamastackclient.QueryConfigRankerRrfParam{
					ImpactFactor: 0,
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
