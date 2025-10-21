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

func TestVectorIoInsertWithOptionalParams(t *testing.T) {
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
	err := client.VectorIo.Insert(context.TODO(), llamastackclient.VectorIoInsertParams{
		Chunks: []llamastackclient.VectorIoInsertParamsChunk{{
			Content: llamastackclient.InterleavedContentUnionParam{
				OfString: llamastackclient.String("string"),
			},
			Metadata: map[string]llamastackclient.VectorIoInsertParamsChunkMetadataUnion{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			ChunkMetadata: llamastackclient.VectorIoInsertParamsChunkChunkMetadata{
				ChunkEmbeddingDimension: llamastackclient.Int(0),
				ChunkEmbeddingModel:     llamastackclient.String("chunk_embedding_model"),
				ChunkID:                 llamastackclient.String("chunk_id"),
				ChunkTokenizer:          llamastackclient.String("chunk_tokenizer"),
				ChunkWindow:             llamastackclient.String("chunk_window"),
				ContentTokenCount:       llamastackclient.Int(0),
				CreatedTimestamp:        llamastackclient.Int(0),
				DocumentID:              llamastackclient.String("document_id"),
				MetadataTokenCount:      llamastackclient.Int(0),
				Source:                  llamastackclient.String("source"),
				UpdatedTimestamp:        llamastackclient.Int(0),
			},
			Embedding:     []float64{0},
			StoredChunkID: llamastackclient.String("stored_chunk_id"),
		}},
		VectorDBID: "vector_db_id",
		TtlSeconds: llamastackclient.Int(0),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorIoQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorIo.Query(context.TODO(), llamastackclient.VectorIoQueryParams{
		Query: llamastackclient.InterleavedContentUnionParam{
			OfString: llamastackclient.String("string"),
		},
		VectorDBID: "vector_db_id",
		Params: map[string]llamastackclient.VectorIoQueryParamsParamUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
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
