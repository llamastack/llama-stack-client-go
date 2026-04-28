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

func TestVectorIoInsertWithOptionalParams(t *testing.T) {
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
	err := client.VectorIo.Insert(context.TODO(), ogxclient.VectorIoInsertParams{
		Chunks: []ogxclient.VectorIoInsertParamsChunk{{
			ChunkID: "chunk_id",
			ChunkMetadata: ogxclient.VectorIoInsertParamsChunkChunkMetadata{
				ChunkID:            ogxclient.String("chunk_id"),
				ChunkTokenizer:     ogxclient.String("chunk_tokenizer"),
				ChunkWindow:        ogxclient.String("chunk_window"),
				ContentTokenCount:  ogxclient.Int(0),
				CreatedTimestamp:   ogxclient.Int(0),
				DocumentID:         ogxclient.String("document_id"),
				MetadataTokenCount: ogxclient.Int(0),
				Source:             ogxclient.String("source"),
				UpdatedTimestamp:   ogxclient.Int(0),
			},
			Content: ogxclient.VectorIoInsertParamsChunkContentUnion{
				OfString: ogxclient.String("string"),
			},
			Embedding:          []float64{0},
			EmbeddingDimension: 0,
			EmbeddingModel:     "embedding_model",
			Metadata: map[string]any{
				"foo": "bar",
			},
		}},
		VectorStoreID: "vector_store_id",
		TtlSeconds:    ogxclient.Int(0),
	})
	if err != nil {
		var apierr *ogxclient.Error
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
	client := ogxclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.VectorIo.Query(context.TODO(), ogxclient.VectorIoQueryParams{
		Query: ogxclient.VectorIoQueryParamsQueryUnion{
			OfString: ogxclient.String("string"),
		},
		VectorStoreID: "vector_store_id",
		Params: map[string]any{
			"foo": "bar",
		},
	})
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
