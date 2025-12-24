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
			ChunkID: "chunk_id",
			Content: llamastackclient.VectorIoInsertParamsChunkContentUnion{
				OfString: llamastackclient.String("string"),
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
			Embedding: []float64{0},
			Metadata: map[string]any{
				"foo": "bar",
			},
		}},
		VectorStoreID: "vector_store_id",
		TtlSeconds:    llamastackclient.Int(0),
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
		Query: llamastackclient.VectorIoQueryParamsQueryUnion{
			OfString: llamastackclient.String("string"),
		},
		VectorStoreID: "vector_store_id",
		Params: map[string]any{
			"foo": "bar",
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
