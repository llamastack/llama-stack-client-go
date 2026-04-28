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

func TestVectorStoreFileBatchNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.FileBatches.New(
		context.TODO(),
		"vector_store_id",
		ogxclient.VectorStoreFileBatchNewParams{
			Attributes: map[string]ogxclient.VectorStoreFileBatchNewParamsAttributeUnion{
				"foo": {
					OfString: ogxclient.String("string"),
				},
			},
			ChunkingStrategy: ogxclient.VectorStoreFileBatchNewParamsChunkingStrategyUnion{
				OfAuto: &ogxclient.VectorStoreFileBatchNewParamsChunkingStrategyAuto{
					Type: "auto",
				},
			},
			FileIDs: []string{"string"},
			Files: []ogxclient.VectorStoreFileBatchNewParamsFile{{
				FileID: "file_id",
				Attributes: map[string]ogxclient.VectorStoreFileBatchNewParamsFileAttributeUnion{
					"foo": {
						OfString: ogxclient.String("string"),
					},
				},
				ChunkingStrategy: ogxclient.VectorStoreFileBatchNewParamsFileChunkingStrategyUnion{
					OfAuto: &ogxclient.VectorStoreFileBatchNewParamsFileChunkingStrategyAuto{
						Type: "auto",
					},
				},
			}},
		},
	)
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorStoreFileBatchGet(t *testing.T) {
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
	_, err := client.VectorStores.FileBatches.Get(
		context.TODO(),
		"batch_id",
		ogxclient.VectorStoreFileBatchGetParams{
			VectorStoreID: "vector_store_id",
		},
	)
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorStoreFileBatchCancel(t *testing.T) {
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
	_, err := client.VectorStores.FileBatches.Cancel(
		context.TODO(),
		"batch_id",
		ogxclient.VectorStoreFileBatchCancelParams{
			VectorStoreID: "vector_store_id",
		},
	)
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorStoreFileBatchListFilesWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.FileBatches.ListFiles(
		context.TODO(),
		"batch_id",
		ogxclient.VectorStoreFileBatchListFilesParams{
			VectorStoreID: "vector_store_id",
			After:         ogxclient.String("after"),
			Before:        ogxclient.String("before"),
			Filter:        ogxclient.String("filter"),
			Limit:         ogxclient.Int(1),
			Order:         ogxclient.String("order"),
		},
	)
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
