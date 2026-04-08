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

func TestVectorStoreFileBatchNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.FileBatches.New(
		context.TODO(),
		"vector_store_id",
		llamastackclient.VectorStoreFileBatchNewParams{
			Attributes: map[string]llamastackclient.VectorStoreFileBatchNewParamsAttributeUnion{
				"foo": {
					OfString: llamastackclient.String("string"),
				},
			},
			ChunkingStrategy: llamastackclient.VectorStoreFileBatchNewParamsChunkingStrategyUnion{
				OfAuto: &llamastackclient.VectorStoreFileBatchNewParamsChunkingStrategyAuto{
					Type: "auto",
				},
			},
			FileIDs: []string{"string"},
			Files: []llamastackclient.VectorStoreFileBatchNewParamsFile{{
				FileID: "file_id",
				Attributes: map[string]llamastackclient.VectorStoreFileBatchNewParamsFileAttributeUnion{
					"foo": {
						OfString: llamastackclient.String("string"),
					},
				},
				ChunkingStrategy: llamastackclient.VectorStoreFileBatchNewParamsFileChunkingStrategyUnion{
					OfAuto: &llamastackclient.VectorStoreFileBatchNewParamsFileChunkingStrategyAuto{
						Type: "auto",
					},
				},
			}},
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

func TestVectorStoreFileBatchGet(t *testing.T) {
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
	_, err := client.VectorStores.FileBatches.Get(
		context.TODO(),
		"batch_id",
		llamastackclient.VectorStoreFileBatchGetParams{
			VectorStoreID: "vector_store_id",
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

func TestVectorStoreFileBatchCancel(t *testing.T) {
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
	_, err := client.VectorStores.FileBatches.Cancel(
		context.TODO(),
		"batch_id",
		llamastackclient.VectorStoreFileBatchCancelParams{
			VectorStoreID: "vector_store_id",
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

func TestVectorStoreFileBatchListFilesWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.FileBatches.ListFiles(
		context.TODO(),
		"batch_id",
		llamastackclient.VectorStoreFileBatchListFilesParams{
			VectorStoreID: "vector_store_id",
			After:         llamastackclient.String("after"),
			Before:        llamastackclient.String("before"),
			Filter:        llamastackclient.String("filter"),
			Limit:         llamastackclient.Int(1),
			Order:         llamastackclient.String("order"),
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
