// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/llama-stack-client-go/internal/apijson"
	"github.com/stainless-sdks/llama-stack-client-go/internal/requestconfig"
	"github.com/stainless-sdks/llama-stack-client-go/option"
	"github.com/stainless-sdks/llama-stack-client-go/packages/param"
	"github.com/stainless-sdks/llama-stack-client-go/shared"
)

// ToolRuntimeRagToolService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolRuntimeRagToolService] method instead.
type ToolRuntimeRagToolService struct {
	Options []option.RequestOption
}

// NewToolRuntimeRagToolService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewToolRuntimeRagToolService(opts ...option.RequestOption) (r ToolRuntimeRagToolService) {
	r = ToolRuntimeRagToolService{}
	r.Options = opts
	return
}

// Index documents so they can be used by the RAG system
func (r *ToolRuntimeRagToolService) Insert(ctx context.Context, body ToolRuntimeRagToolInsertParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/tool-runtime/rag-tool/insert"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

// Query the RAG system for context; typically invoked by the agent
func (r *ToolRuntimeRagToolService) Query(ctx context.Context, body ToolRuntimeRagToolQueryParams, opts ...option.RequestOption) (res *shared.QueryResult, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tool-runtime/rag-tool/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ToolRuntimeRagToolInsertParams struct {
	ChunkSizeInTokens int64                  `json:"chunk_size_in_tokens,required"`
	Documents         []shared.DocumentParam `json:"documents,omitzero,required"`
	VectorDBID        string                 `json:"vector_db_id,required"`
	paramObj
}

func (r ToolRuntimeRagToolInsertParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolInsertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolInsertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolRuntimeRagToolQueryParams struct {
	// A image content item
	Content     shared.InterleavedContentUnionParam `json:"content,omitzero,required"`
	VectorDBIDs []string                            `json:"vector_db_ids,omitzero,required"`
	// Configuration for the RAG query generation.
	QueryConfig shared.QueryConfigParam `json:"query_config,omitzero"`
	paramObj
}

func (r ToolRuntimeRagToolQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolRuntimeRagToolQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRuntimeRagToolQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
