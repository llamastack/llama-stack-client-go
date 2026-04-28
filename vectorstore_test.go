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

func TestVectorStoreNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.New(context.TODO(), ogxclient.VectorStoreNewParams{
		ChunkingStrategy: ogxclient.VectorStoreNewParamsChunkingStrategyUnion{
			OfAuto: &ogxclient.VectorStoreNewParamsChunkingStrategyAuto{
				Type: "auto",
			},
		},
		Description: ogxclient.String("description"),
		ExpiresAfter: ogxclient.VectorStoreNewParamsExpiresAfter{
			Anchor: "last_active_at",
			Days:   1,
		},
		FileIDs: []string{"string"},
		Metadata: map[string]any{
			"foo": "bar",
		},
		Name: ogxclient.String("name"),
	})
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorStoreGet(t *testing.T) {
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
	_, err := client.VectorStores.Get(context.TODO(), "vector_store_id")
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorStoreUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.Update(
		context.TODO(),
		"vector_store_id",
		ogxclient.VectorStoreUpdateParams{
			ExpiresAfter: ogxclient.VectorStoreUpdateParamsExpiresAfter{
				Anchor: "last_active_at",
				Days:   1,
			},
			Metadata: map[string]any{
				"foo": "bar",
			},
			Name: ogxclient.String("name"),
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

func TestVectorStoreListWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.List(context.TODO(), ogxclient.VectorStoreListParams{
		After:  ogxclient.String("after"),
		Before: ogxclient.String("before"),
		Limit:  ogxclient.Int(1),
		Order:  ogxclient.String("order"),
	})
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorStoreDelete(t *testing.T) {
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
	_, err := client.VectorStores.Delete(context.TODO(), "vector_store_id")
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVectorStoreSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.Search(
		context.TODO(),
		"vector_store_id",
		ogxclient.VectorStoreSearchParams{
			Query: ogxclient.VectorStoreSearchParamsQueryUnion{
				OfString: ogxclient.String("string"),
			},
			Filters: map[string]any{
				"foo": "bar",
			},
			MaxNumResults: ogxclient.Int(1),
			RankingOptions: ogxclient.VectorStoreSearchParamsRankingOptions{
				Alpha:          ogxclient.Float(0),
				ImpactFactor:   ogxclient.Float(0),
				Model:          ogxclient.String("model"),
				Ranker:         ogxclient.String("ranker"),
				ScoreThreshold: ogxclient.Float(0),
				Weights: map[string]float64{
					"foo": 0,
				},
			},
			RewriteQuery: ogxclient.Bool(true),
			SearchMode:   ogxclient.String("search_mode"),
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
