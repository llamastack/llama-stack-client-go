// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/llamastack/llama-stack-client-go/internal/apijson"
	"github.com/llamastack/llama-stack-client-go/internal/requestconfig"
	"github.com/llamastack/llama-stack-client-go/option"
	"github.com/llamastack/llama-stack-client-go/packages/param"
	"github.com/llamastack/llama-stack-client-go/packages/respjson"
	"github.com/llamastack/llama-stack-client-go/shared/constant"
)

// ScoringFunctionService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScoringFunctionService] method instead.
type ScoringFunctionService struct {
	Options []option.RequestOption
}

// NewScoringFunctionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScoringFunctionService(opts ...option.RequestOption) (r ScoringFunctionService) {
	r = ScoringFunctionService{}
	r.Options = opts
	return
}

// Get a scoring function by its ID.
func (r *ScoringFunctionService) Get(ctx context.Context, scoringFnID string, opts ...option.RequestOption) (res *ScoringFn, err error) {
	opts = append(r.Options[:], opts...)
	if scoringFnID == "" {
		err = errors.New("missing required scoring_fn_id parameter")
		return
	}
	path := fmt.Sprintf("v1/scoring-functions/%s", scoringFnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all scoring functions.
func (r *ScoringFunctionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ScoringFn, err error) {
	var env ListScoringFunctionsResponse
	opts = append(r.Options[:], opts...)
	path := "v1/scoring-functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Register a scoring function.
func (r *ScoringFunctionService) Register(ctx context.Context, body ScoringFunctionRegisterParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "v1/scoring-functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type ListScoringFunctionsResponse struct {
	Data []ScoringFn `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListScoringFunctionsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListScoringFunctionsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A scoring function resource for evaluating model outputs.
type ScoringFn struct {
	Identifier string                            `json:"identifier,required"`
	Metadata   map[string]ScoringFnMetadataUnion `json:"metadata,required"`
	ProviderID string                            `json:"provider_id,required"`
	ReturnType ScoringFnReturnType               `json:"return_type,required"`
	// The resource type, always scoring_function
	Type        constant.ScoringFunction `json:"type,required"`
	Description string                   `json:"description"`
	// Parameters for LLM-as-judge scoring function configuration.
	Params             ScoringFnParamsUnionResp `json:"params"`
	ProviderResourceID string                   `json:"provider_resource_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Identifier         respjson.Field
		Metadata           respjson.Field
		ProviderID         respjson.Field
		ReturnType         respjson.Field
		Type               respjson.Field
		Description        respjson.Field
		Params             respjson.Field
		ProviderResourceID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFn) RawJSON() string { return r.JSON.raw }
func (r *ScoringFn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringFnMetadataUnion contains all possible properties and values from [bool],
// [float64], [string], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfBool OfFloat OfString OfAnyArray]
type ScoringFnMetadataUnion struct {
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfBool     respjson.Field
		OfFloat    respjson.Field
		OfString   respjson.Field
		OfAnyArray respjson.Field
		raw        string
	} `json:"-"`
}

func (u ScoringFnMetadataUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnMetadataUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnMetadataUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnMetadataUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringFnMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoringFnMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFnReturnType struct {
	// Any of "string", "number", "boolean", "array", "object", "json", "union",
	// "chat_completion_input", "completion_input", "agent_turn_input".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnReturnType) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnReturnType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoringFnParamsUnionResp contains all possible properties and values from
// [ScoringFnParamsLlmAsJudgeResp], [ScoringFnParamsRegexParserResp],
// [ScoringFnParamsBasicResp].
//
// Use the [ScoringFnParamsUnionResp.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ScoringFnParamsUnionResp struct {
	AggregationFunctions []string `json:"aggregation_functions"`
	// This field is from variant [ScoringFnParamsLlmAsJudgeResp].
	JudgeModel string `json:"judge_model"`
	// This field is from variant [ScoringFnParamsLlmAsJudgeResp].
	JudgeScoreRegexes []string `json:"judge_score_regexes"`
	// Any of "llm_as_judge", "regex_parser", "basic".
	Type string `json:"type"`
	// This field is from variant [ScoringFnParamsLlmAsJudgeResp].
	PromptTemplate string `json:"prompt_template"`
	// This field is from variant [ScoringFnParamsRegexParserResp].
	ParsingRegexes []string `json:"parsing_regexes"`
	JSON           struct {
		AggregationFunctions respjson.Field
		JudgeModel           respjson.Field
		JudgeScoreRegexes    respjson.Field
		Type                 respjson.Field
		PromptTemplate       respjson.Field
		ParsingRegexes       respjson.Field
		raw                  string
	} `json:"-"`
}

// anyScoringFnParamsResp is implemented by each variant of
// [ScoringFnParamsUnionResp] to add type safety for the return type of
// [ScoringFnParamsUnionResp.AsAny]
type anyScoringFnParamsResp interface {
	implScoringFnParamsUnionResp()
}

func (ScoringFnParamsLlmAsJudgeResp) implScoringFnParamsUnionResp()  {}
func (ScoringFnParamsRegexParserResp) implScoringFnParamsUnionResp() {}
func (ScoringFnParamsBasicResp) implScoringFnParamsUnionResp()       {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ScoringFnParamsUnionResp.AsAny().(type) {
//	case llamastackclient.ScoringFnParamsLlmAsJudgeResp:
//	case llamastackclient.ScoringFnParamsRegexParserResp:
//	case llamastackclient.ScoringFnParamsBasicResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ScoringFnParamsUnionResp) AsAny() anyScoringFnParamsResp {
	switch u.Type {
	case "llm_as_judge":
		return u.AsLlmAsJudge()
	case "regex_parser":
		return u.AsRegexParser()
	case "basic":
		return u.AsBasic()
	}
	return nil
}

func (u ScoringFnParamsUnionResp) AsLlmAsJudge() (v ScoringFnParamsLlmAsJudgeResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnParamsUnionResp) AsRegexParser() (v ScoringFnParamsRegexParserResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoringFnParamsUnionResp) AsBasic() (v ScoringFnParamsBasicResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoringFnParamsUnionResp) RawJSON() string { return u.JSON.raw }

func (r *ScoringFnParamsUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ScoringFnParamsUnionResp to a ScoringFnParamsUnion.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ScoringFnParamsUnion.Overrides()
func (r ScoringFnParamsUnionResp) ToParam() ScoringFnParamsUnion {
	return param.Override[ScoringFnParamsUnion](json.RawMessage(r.RawJSON()))
}

// Parameters for LLM-as-judge scoring function configuration.
type ScoringFnParamsLlmAsJudgeResp struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,required"`
	// Identifier of the LLM model to use as a judge for scoring
	JudgeModel string `json:"judge_model,required"`
	// Regexes to extract the answer from generated response
	JudgeScoreRegexes []string `json:"judge_score_regexes,required"`
	// The type of scoring function parameters, always llm_as_judge
	Type constant.LlmAsJudge `json:"type,required"`
	// (Optional) Custom prompt template for the judge model
	PromptTemplate string `json:"prompt_template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationFunctions respjson.Field
		JudgeModel           respjson.Field
		JudgeScoreRegexes    respjson.Field
		Type                 respjson.Field
		PromptTemplate       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsLlmAsJudgeResp) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsLlmAsJudgeResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for regex parser scoring function configuration.
type ScoringFnParamsRegexParserResp struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,required"`
	// Regex to extract the answer from generated response
	ParsingRegexes []string `json:"parsing_regexes,required"`
	// The type of scoring function parameters, always regex_parser
	Type constant.RegexParser `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationFunctions respjson.Field
		ParsingRegexes       respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsRegexParserResp) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsRegexParserResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for basic scoring function configuration.
type ScoringFnParamsBasicResp struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,required"`
	// The type of scoring function parameters, always basic
	Type constant.Basic `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AggregationFunctions respjson.Field
		Type                 respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoringFnParamsBasicResp) RawJSON() string { return r.JSON.raw }
func (r *ScoringFnParamsBasicResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func ScoringFnParamsOfLlmAsJudge(aggregationFunctions []string, judgeModel string, judgeScoreRegexes []string) ScoringFnParamsUnion {
	var llmAsJudge ScoringFnParamsLlmAsJudge
	llmAsJudge.AggregationFunctions = aggregationFunctions
	llmAsJudge.JudgeModel = judgeModel
	llmAsJudge.JudgeScoreRegexes = judgeScoreRegexes
	return ScoringFnParamsUnion{OfLlmAsJudge: &llmAsJudge}
}

func ScoringFnParamsOfRegexParser(aggregationFunctions []string, parsingRegexes []string) ScoringFnParamsUnion {
	var regexParser ScoringFnParamsRegexParser
	regexParser.AggregationFunctions = aggregationFunctions
	regexParser.ParsingRegexes = parsingRegexes
	return ScoringFnParamsUnion{OfRegexParser: &regexParser}
}

func ScoringFnParamsOfBasic(aggregationFunctions []string) ScoringFnParamsUnion {
	var basic ScoringFnParamsBasic
	basic.AggregationFunctions = aggregationFunctions
	return ScoringFnParamsUnion{OfBasic: &basic}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ScoringFnParamsUnion struct {
	OfLlmAsJudge  *ScoringFnParamsLlmAsJudge  `json:",omitzero,inline"`
	OfRegexParser *ScoringFnParamsRegexParser `json:",omitzero,inline"`
	OfBasic       *ScoringFnParamsBasic       `json:",omitzero,inline"`
	paramUnion
}

func (u ScoringFnParamsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlmAsJudge, u.OfRegexParser, u.OfBasic)
}
func (u *ScoringFnParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ScoringFnParamsUnion) asAny() any {
	if !param.IsOmitted(u.OfLlmAsJudge) {
		return u.OfLlmAsJudge
	} else if !param.IsOmitted(u.OfRegexParser) {
		return u.OfRegexParser
	} else if !param.IsOmitted(u.OfBasic) {
		return u.OfBasic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetJudgeModel() *string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return &vt.JudgeModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetJudgeScoreRegexes() []string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return vt.JudgeScoreRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetPromptTemplate() *string {
	if vt := u.OfLlmAsJudge; vt != nil && vt.PromptTemplate.Valid() {
		return &vt.PromptTemplate.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetParsingRegexes() []string {
	if vt := u.OfRegexParser; vt != nil {
		return vt.ParsingRegexes
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ScoringFnParamsUnion) GetType() *string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRegexParser; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfBasic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's AggregationFunctions property, if
// present.
func (u ScoringFnParamsUnion) GetAggregationFunctions() []string {
	if vt := u.OfLlmAsJudge; vt != nil {
		return vt.AggregationFunctions
	} else if vt := u.OfRegexParser; vt != nil {
		return vt.AggregationFunctions
	} else if vt := u.OfBasic; vt != nil {
		return vt.AggregationFunctions
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ScoringFnParamsUnion](
		"type",
		apijson.Discriminator[ScoringFnParamsLlmAsJudge]("llm_as_judge"),
		apijson.Discriminator[ScoringFnParamsRegexParser]("regex_parser"),
		apijson.Discriminator[ScoringFnParamsBasic]("basic"),
	)
}

// Parameters for LLM-as-judge scoring function configuration.
//
// The properties AggregationFunctions, JudgeModel, JudgeScoreRegexes, Type are
// required.
type ScoringFnParamsLlmAsJudge struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero,required"`
	// Identifier of the LLM model to use as a judge for scoring
	JudgeModel string `json:"judge_model,required"`
	// Regexes to extract the answer from generated response
	JudgeScoreRegexes []string `json:"judge_score_regexes,omitzero,required"`
	// (Optional) Custom prompt template for the judge model
	PromptTemplate param.Opt[string] `json:"prompt_template,omitzero"`
	// The type of scoring function parameters, always llm_as_judge
	//
	// This field can be elided, and will marshal its zero value as "llm_as_judge".
	Type constant.LlmAsJudge `json:"type,required"`
	paramObj
}

func (r ScoringFnParamsLlmAsJudge) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFnParamsLlmAsJudge
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFnParamsLlmAsJudge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for regex parser scoring function configuration.
//
// The properties AggregationFunctions, ParsingRegexes, Type are required.
type ScoringFnParamsRegexParser struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero,required"`
	// Regex to extract the answer from generated response
	ParsingRegexes []string `json:"parsing_regexes,omitzero,required"`
	// The type of scoring function parameters, always regex_parser
	//
	// This field can be elided, and will marshal its zero value as "regex_parser".
	Type constant.RegexParser `json:"type,required"`
	paramObj
}

func (r ScoringFnParamsRegexParser) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFnParamsRegexParser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFnParamsRegexParser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Parameters for basic scoring function configuration.
//
// The properties AggregationFunctions, Type are required.
type ScoringFnParamsBasic struct {
	// Aggregation functions to apply to the scores of each row
	//
	// Any of "average", "weighted_average", "median", "categorical_count", "accuracy".
	AggregationFunctions []string `json:"aggregation_functions,omitzero,required"`
	// The type of scoring function parameters, always basic
	//
	// This field can be elided, and will marshal its zero value as "basic".
	Type constant.Basic `json:"type,required"`
	paramObj
}

func (r ScoringFnParamsBasic) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFnParamsBasic
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFnParamsBasic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoringFunctionRegisterParams struct {
	// The description of the scoring function.
	Description string                                  `json:"description,required"`
	ReturnType  ScoringFunctionRegisterParamsReturnType `json:"return_type,omitzero,required"`
	// The ID of the scoring function to register.
	ScoringFnID string `json:"scoring_fn_id,required"`
	// The ID of the provider to use for the scoring function.
	ProviderID param.Opt[string] `json:"provider_id,omitzero"`
	// The ID of the provider scoring function to use for the scoring function.
	ProviderScoringFnID param.Opt[string] `json:"provider_scoring_fn_id,omitzero"`
	// The parameters for the scoring function for benchmark eval, these can be
	// overridden for app eval.
	Params ScoringFnParamsUnion `json:"params,omitzero"`
	paramObj
}

func (r ScoringFunctionRegisterParams) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFunctionRegisterParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFunctionRegisterParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ScoringFunctionRegisterParamsReturnType struct {
	// Any of "string", "number", "boolean", "array", "object", "json", "union",
	// "chat_completion_input", "completion_input", "agent_turn_input".
	Type string `json:"type,omitzero,required"`
	paramObj
}

func (r ScoringFunctionRegisterParamsReturnType) MarshalJSON() (data []byte, err error) {
	type shadow ScoringFunctionRegisterParamsReturnType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ScoringFunctionRegisterParamsReturnType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ScoringFunctionRegisterParamsReturnType](
		"type", "string", "number", "boolean", "array", "object", "json", "union", "chat_completion_input", "completion_input", "agent_turn_input",
	)
}
