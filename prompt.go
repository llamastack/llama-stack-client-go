// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/apiquery"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
)

// PromptService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPromptService] method instead.
type PromptService struct {
	Options  []option.RequestOption
	Versions PromptVersionService
}

// NewPromptService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPromptService(opts ...option.RequestOption) (r PromptService) {
	r = PromptService{}
	r.Options = opts
	r.Versions = NewPromptVersionService(opts...)
	return
}

// Create prompt. Create a new prompt.
func (r *PromptService) New(ctx context.Context, body PromptNewParams, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get prompt. Get a prompt by its identifier and optional version.
func (r *PromptService) Get(ctx context.Context, promptID string, query PromptGetParams, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return
	}
	path := fmt.Sprintf("v1/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update prompt. Update an existing prompt (increments version).
func (r *PromptService) Update(ctx context.Context, promptID string, body PromptUpdateParams, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return
	}
	path := fmt.Sprintf("v1/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all prompts.
func (r *PromptService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Prompt, err error) {
	var env ListPromptsResponse
	opts = slices.Concat(r.Options, opts)
	path := "v1/prompts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Delete prompt. Delete a prompt.
func (r *PromptService) Delete(ctx context.Context, promptID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return
	}
	path := fmt.Sprintf("v1/prompts/%s", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Set prompt version. Set which version of a prompt should be the default in
// get_prompt (latest).
func (r *PromptService) SetDefaultVersion(ctx context.Context, promptID string, body PromptSetDefaultVersionParams, opts ...option.RequestOption) (res *Prompt, err error) {
	opts = slices.Concat(r.Options, opts)
	if promptID == "" {
		err = errors.New("missing required prompt_id parameter")
		return
	}
	path := fmt.Sprintf("v1/prompts/%s/set-default-version", promptID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response model to list prompts.
type ListPromptsResponse struct {
	Data []Prompt `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListPromptsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListPromptsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A prompt resource representing a stored OpenAI Compatible prompt template in
// Llama Stack.
type Prompt struct {
	// Boolean indicating whether this version is the default version for this prompt
	IsDefault bool `json:"is_default,required"`
	// Unique identifier formatted as 'pmpt\_<48-digit-hash>'
	PromptID string `json:"prompt_id,required"`
	// List of prompt variable names that can be used in the prompt template
	Variables []string `json:"variables,required"`
	// Version (integer starting at 1, incremented on save)
	Version int64 `json:"version,required"`
	// The system prompt text with variable placeholders. Variables are only supported
	// when using the Responses API.
	Prompt string `json:"prompt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsDefault   respjson.Field
		PromptID    respjson.Field
		Variables   respjson.Field
		Version     respjson.Field
		Prompt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Prompt) RawJSON() string { return r.JSON.raw }
func (r *Prompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptNewParams struct {
	// The prompt text content with variable placeholders.
	Prompt string `json:"prompt,required"`
	// List of variable names that can be used in the prompt template.
	Variables []string `json:"variables,omitzero"`
	paramObj
}

func (r PromptNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PromptNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptGetParams struct {
	// The version of the prompt to get (defaults to latest).
	Version param.Opt[int64] `query:"version,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PromptGetParams]'s query parameters as `url.Values`.
func (r PromptGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PromptUpdateParams struct {
	// The updated prompt text content.
	Prompt string `json:"prompt,required"`
	// Set the new version as the default (default=True).
	SetAsDefault bool `json:"set_as_default,required"`
	// The current version of the prompt being updated.
	Version int64 `json:"version,required"`
	// Updated list of variable names that can be used in the prompt template.
	Variables []string `json:"variables,omitzero"`
	paramObj
}

func (r PromptUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PromptUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptSetDefaultVersionParams struct {
	// The version to set as default.
	Version int64 `json:"version,required"`
	paramObj
}

func (r PromptSetDefaultVersionParams) MarshalJSON() (data []byte, err error) {
	type shadow PromptSetDefaultVersionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptSetDefaultVersionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
