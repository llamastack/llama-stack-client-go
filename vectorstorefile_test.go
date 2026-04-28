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

func TestVectorStoreFileNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.Files.New(
		context.TODO(),
		"vector_store_id",
		ogxclient.VectorStoreFileNewParams{
			FileID: "file_id",
			Attributes: map[string]ogxclient.VectorStoreFileNewParamsAttributeUnion{
				"foo": {
					OfString: ogxclient.String("string"),
				},
			},
			ChunkingStrategy: ogxclient.VectorStoreFileNewParamsChunkingStrategyUnion{
				OfAuto: &ogxclient.VectorStoreFileNewParamsChunkingStrategyAuto{
					Type: "auto",
				},
			},
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

func TestVectorStoreFileGet(t *testing.T) {
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
	_, err := client.VectorStores.Files.Get(
		context.TODO(),
		"file_id",
		ogxclient.VectorStoreFileGetParams{
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

func TestVectorStoreFileUpdate(t *testing.T) {
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
	_, err := client.VectorStores.Files.Update(
		context.TODO(),
		"file_id",
		ogxclient.VectorStoreFileUpdateParams{
			VectorStoreID: "vector_store_id",
			Attributes: map[string]any{
				"foo": "bar",
			},
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

func TestVectorStoreFileListWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.Files.List(
		context.TODO(),
		"vector_store_id",
		ogxclient.VectorStoreFileListParams{
			After:  ogxclient.String("after"),
			Before: ogxclient.String("before"),
			Filter: ogxclient.VectorStoreFileListParamsFilterInProgress,
			Limit:  ogxclient.Int(1),
			Order:  ogxclient.String("order"),
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

func TestVectorStoreFileDelete(t *testing.T) {
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
	_, err := client.VectorStores.Files.Delete(
		context.TODO(),
		"file_id",
		ogxclient.VectorStoreFileDeleteParams{
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

func TestVectorStoreFileContentWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.Files.Content(
		context.TODO(),
		"file_id",
		ogxclient.VectorStoreFileContentParams{
			VectorStoreID:     "vector_store_id",
			IncludeEmbeddings: ogxclient.Bool(true),
			IncludeMetadata:   ogxclient.Bool(true),
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
