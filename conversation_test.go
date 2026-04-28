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

func TestConversationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Conversations.New(context.TODO(), ogxclient.ConversationNewParams{
		Items: []ogxclient.ConversationNewParamsItemUnion{{
			OfMessage: &ogxclient.ConversationNewParamsItemMessage{
				Content: ogxclient.ConversationNewParamsItemMessageContentUnion{
					OfString: ogxclient.String("string"),
				},
				Role:   ogxclient.ConversationNewParamsItemMessageRoleSystem,
				ID:     ogxclient.String("id"),
				Status: ogxclient.String("status"),
				Type:   "message",
			},
		}},
		Metadata: map[string]string{
			"foo": "string",
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

func TestConversationGet(t *testing.T) {
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
	_, err := client.Conversations.Get(context.TODO(), "conversation_id")
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConversationUpdate(t *testing.T) {
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
	_, err := client.Conversations.Update(
		context.TODO(),
		"conversation_id",
		ogxclient.ConversationUpdateParams{
			Metadata: map[string]string{
				"foo": "string",
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

func TestConversationDelete(t *testing.T) {
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
	_, err := client.Conversations.Delete(context.TODO(), "conversation_id")
	if err != nil {
		var apierr *ogxclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
