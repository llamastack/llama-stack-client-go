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

func TestVectorStoreNewWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.New(context.TODO(), llamastackclient.VectorStoreNewParams{
		Name: "name",
		ChunkingStrategy: map[string]llamastackclient.VectorStoreNewParamsChunkingStrategyUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		EmbeddingDimension: llamastackclient.Int(0),
		EmbeddingModel:     llamastackclient.String("embedding_model"),
		ExpiresAfter: map[string]llamastackclient.VectorStoreNewParamsExpiresAfterUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		FileIDs: []string{"string"},
		Metadata: map[string]llamastackclient.VectorStoreNewParamsMetadataUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		ProviderID:         llamastackclient.String("provider_id"),
		ProviderVectorDBID: llamastackclient.String("provider_vector_db_id"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.VectorStores.Get(context.TODO(), "vector_store_id")
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.VectorStores.Update(
		context.TODO(),
		"vector_store_id",
		llamastackclient.VectorStoreUpdateParams{
			ExpiresAfter: map[string]llamastackclient.VectorStoreUpdateParamsExpiresAfterUnion{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			Metadata: map[string]llamastackclient.VectorStoreUpdateParamsMetadataUnion{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			Name: llamastackclient.String("name"),
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

func TestVectorStoreListWithOptionalParams(t *testing.T) {
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
	_, err := client.VectorStores.List(context.TODO(), llamastackclient.VectorStoreListParams{
		After:  llamastackclient.String("after"),
		Before: llamastackclient.String("before"),
		Limit:  llamastackclient.Int(0),
		Order:  llamastackclient.String("order"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.VectorStores.Delete(context.TODO(), "vector_store_id")
	if err != nil {
		var apierr *llamastackclient.Error
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
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
	)
	_, err := client.VectorStores.Search(
		context.TODO(),
		"vector_store_id",
		llamastackclient.VectorStoreSearchParams{
			Query: llamastackclient.VectorStoreSearchParamsQueryUnion{
				OfString: llamastackclient.String("string"),
			},
			Filters: map[string]llamastackclient.VectorStoreSearchParamsFilterUnion{
				"foo": {
					OfBool: llamastackclient.Bool(true),
				},
			},
			MaxNumResults: llamastackclient.Int(0),
			RankingOptions: llamastackclient.VectorStoreSearchParamsRankingOptions{
				Ranker:         llamastackclient.String("ranker"),
				ScoreThreshold: llamastackclient.Float(0),
			},
			RewriteQuery: llamastackclient.Bool(true),
			SearchMode:   llamastackclient.String("search_mode"),
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
