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
	"os"
	"testing"

	"github.com/ogx-ai/ogx-client-go"
	"github.com/ogx-ai/ogx-client-go/internal/testutil"
	"github.com/ogx-ai/ogx-client-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Responses.ListAutoPaging(context.TODO(), ogxclient.ResponseListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		response := iter.Current()
		t.Logf("%+v\n", response.ID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
