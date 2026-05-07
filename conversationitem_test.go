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

func TestConversationItemNew(t *testing.T) {
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
	_, err := client.Conversations.Items.New(
		context.TODO(),
		"conversation_id",
		ogxclient.ConversationItemNewParams{
			Items: []ogxclient.ConversationItemNewParamsItemUnion{{
				OfMessage: &ogxclient.ConversationItemNewParamsItemMessage{
					Content: ogxclient.ConversationItemNewParamsItemMessageContentUnion{
						OfString: ogxclient.String("string"),
					},
					Role:   ogxclient.ConversationItemNewParamsItemMessageRoleSystem,
					ID:     ogxclient.String("id"),
					Status: ogxclient.String("status"),
					Type:   "message",
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

func TestConversationItemListWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.Items.List(
		context.TODO(),
		"conversation_id",
		ogxclient.ConversationItemListParams{
			After:   ogxclient.String("after"),
			Include: []string{"web_search_call.action.sources"},
			Limit:   ogxclient.Int(0),
			Order:   ogxclient.ConversationItemListParamsOrderAsc,
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

func TestConversationItemDelete(t *testing.T) {
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
	_, err := client.Conversations.Items.Delete(
		context.TODO(),
		"item_id",
		ogxclient.ConversationItemDeleteParams{
			ConversationID: "conversation_id",
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

func TestConversationItemGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.Items.Get(
		context.TODO(),
		"item_id",
		ogxclient.ConversationItemGetParams{
			ConversationID: "conversation_id",
			Include:        []string{"web_search_call.action.sources"},
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
