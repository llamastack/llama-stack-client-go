// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type DatasetsIterrows[T any] struct {
	Data      []T   `json:"data"`
	NextIndex int64 `json:"next_index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		NextIndex   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r DatasetsIterrows[T]) RawJSON() string { return r.JSON.raw }
func (r *DatasetsIterrows[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *DatasetsIterrows[T]) GetNextPage() (res *DatasetsIterrows[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	next := r.NextIndex
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("start_index", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *DatasetsIterrows[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &DatasetsIterrows[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type DatasetsIterrowsAutoPager[T any] struct {
	page *DatasetsIterrows[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewDatasetsIterrowsAutoPager[T any](page *DatasetsIterrows[T], err error) *DatasetsIterrowsAutoPager[T] {
	return &DatasetsIterrowsAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *DatasetsIterrowsAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *DatasetsIterrowsAutoPager[T]) Current() T {
	return r.cur
}

func (r *DatasetsIterrowsAutoPager[T]) Err() error {
	return r.err
}

func (r *DatasetsIterrowsAutoPager[T]) Index() int {
	return r.run
}

type OpenAICursorPagination[T any] struct {
	Data    []T    `json:"data"`
	HasMore bool   `json:"has_more"`
	LastID  string `json:"last_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HasMore     respjson.Field
		LastID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r OpenAICursorPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *OpenAICursorPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *OpenAICursorPagination[T]) GetNextPage() (res *OpenAICursorPagination[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}

	if r.JSON.HasMore.Valid() && r.HasMore == false {
		return nil, nil
	}
	next := r.LastID
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("after", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *OpenAICursorPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &OpenAICursorPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type OpenAICursorPaginationAutoPager[T any] struct {
	page *OpenAICursorPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewOpenAICursorPaginationAutoPager[T any](page *OpenAICursorPagination[T], err error) *OpenAICursorPaginationAutoPager[T] {
	return &OpenAICursorPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *OpenAICursorPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *OpenAICursorPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *OpenAICursorPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *OpenAICursorPaginationAutoPager[T]) Index() int {
	return r.run
}
