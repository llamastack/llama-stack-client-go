// Copyright (c) Meta Platforms, Inc. and affiliates.
// All rights reserved.
//
// This source code is licensed under the terms described in the LICENSE file in
// the root directory of this source tree.
//
// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/llamastack/llama-stack-client-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Assistant string                          // Always "assistant"
type Auto string                               // Always "auto"
type Basic string                              // Always "basic"
type Benchmark string                          // Always "benchmark"
type ChatCompletion string                     // Always "chat.completion"
type ChatCompletionChunk string                // Always "chat.completion.chunk"
type ContainerFileCitation string              // Always "container_file_citation"
type Conversation string                       // Always "conversation"
type CreatedAt string                          // Always "created_at"
type Dataset string                            // Always "dataset"
type Developer string                          // Always "developer"
type Embedding string                          // Always "embedding"
type File string                               // Always "file"
type FileCitation string                       // Always "file_citation"
type FilePath string                           // Always "file_path"
type FileSearch string                         // Always "file_search"
type FileSearchCall string                     // Always "file_search_call"
type Function string                           // Always "function"
type FunctionCall string                       // Always "function_call"
type FunctionCallOutput string                 // Always "function_call_output"
type Greedy string                             // Always "greedy"
type Image string                              // Always "image"
type ImageURL string                           // Always "image_url"
type InputFile string                          // Always "input_file"
type InputImage string                         // Always "input_image"
type InputText string                          // Always "input_text"
type JsonObject string                         // Always "json_object"
type JsonSchema string                         // Always "json_schema"
type List string                               // Always "list"
type LlmAsJudge string                         // Always "llm_as_judge"
type LoRa string                               // Always "LoRA"
type Mcp string                                // Always "mcp"
type McpApprovalRequest string                 // Always "mcp_approval_request"
type McpApprovalResponse string                // Always "mcp_approval_response"
type McpCall string                            // Always "mcp_call"
type McpListTools string                       // Always "mcp_list_tools"
type Message string                            // Always "message"
type Model string                              // Always "model"
type OutputText string                         // Always "output_text"
type Qat string                                // Always "QAT"
type ReasoningText string                      // Always "reasoning_text"
type Refusal string                            // Always "refusal"
type RegexParser string                        // Always "regex_parser"
type Response string                           // Always "response"
type ResponseCompleted string                  // Always "response.completed"
type ResponseContentPartAdded string           // Always "response.content_part.added"
type ResponseContentPartDone string            // Always "response.content_part.done"
type ResponseCreated string                    // Always "response.created"
type ResponseFailed string                     // Always "response.failed"
type ResponseFileSearchCallCompleted string    // Always "response.file_search_call.completed"
type ResponseFileSearchCallInProgress string   // Always "response.file_search_call.in_progress"
type ResponseFileSearchCallSearching string    // Always "response.file_search_call.searching"
type ResponseFunctionCallArgumentsDelta string // Always "response.function_call_arguments.delta"
type ResponseFunctionCallArgumentsDone string  // Always "response.function_call_arguments.done"
type ResponseInProgress string                 // Always "response.in_progress"
type ResponseIncomplete string                 // Always "response.incomplete"
type ResponseMcpCallArgumentsDelta string      // Always "response.mcp_call.arguments.delta"
type ResponseMcpCallArgumentsDone string       // Always "response.mcp_call.arguments.done"
type ResponseMcpCallCompleted string           // Always "response.mcp_call.completed"
type ResponseMcpCallFailed string              // Always "response.mcp_call.failed"
type ResponseMcpCallInProgress string          // Always "response.mcp_call.in_progress"
type ResponseMcpListToolsCompleted string      // Always "response.mcp_list_tools.completed"
type ResponseMcpListToolsFailed string         // Always "response.mcp_list_tools.failed"
type ResponseMcpListToolsInProgress string     // Always "response.mcp_list_tools.in_progress"
type ResponseOutputItemAdded string            // Always "response.output_item.added"
type ResponseOutputItemDone string             // Always "response.output_item.done"
type ResponseOutputTextAnnotationAdded string  // Always "response.output_text.annotation.added"
type ResponseOutputTextDelta string            // Always "response.output_text.delta"
type ResponseOutputTextDone string             // Always "response.output_text.done"
type ResponseReasoningSummaryPartAdded string  // Always "response.reasoning_summary_part.added"
type ResponseReasoningSummaryPartDone string   // Always "response.reasoning_summary_part.done"
type ResponseReasoningSummaryTextDelta string  // Always "response.reasoning_summary_text.delta"
type ResponseReasoningSummaryTextDone string   // Always "response.reasoning_summary_text.done"
type ResponseReasoningTextDelta string         // Always "response.reasoning_text.delta"
type ResponseReasoningTextDone string          // Always "response.reasoning_text.done"
type ResponseRefusalDelta string               // Always "response.refusal.delta"
type ResponseRefusalDone string                // Always "response.refusal.done"
type ResponseWebSearchCallCompleted string     // Always "response.web_search_call.completed"
type ResponseWebSearchCallInProgress string    // Always "response.web_search_call.in_progress"
type ResponseWebSearchCallSearching string     // Always "response.web_search_call.searching"
type Rows string                               // Always "rows"
type ScoringFunction string                    // Always "scoring_function"
type Shield string                             // Always "shield"
type Static string                             // Always "static"
type SummaryText string                        // Always "summary_text"
type System string                             // Always "system"
type Text string                               // Always "text"
type TextCompletion string                     // Always "text_completion"
type Tool string                               // Always "tool"
type ToolGroup string                          // Always "tool_group"
type TopK string                               // Always "top_k"
type TopP string                               // Always "top_p"
type Uri string                                // Always "uri"
type URLCitation string                        // Always "url_citation"
type User string                               // Always "user"
type VectorStoreFileContentPage string         // Always "vector_store.file_content.page"
type WebSearchCall string                      // Always "web_search_call"

func (c Assistant) Default() Assistant                         { return "assistant" }
func (c Auto) Default() Auto                                   { return "auto" }
func (c Basic) Default() Basic                                 { return "basic" }
func (c Benchmark) Default() Benchmark                         { return "benchmark" }
func (c ChatCompletion) Default() ChatCompletion               { return "chat.completion" }
func (c ChatCompletionChunk) Default() ChatCompletionChunk     { return "chat.completion.chunk" }
func (c ContainerFileCitation) Default() ContainerFileCitation { return "container_file_citation" }
func (c Conversation) Default() Conversation                   { return "conversation" }
func (c CreatedAt) Default() CreatedAt                         { return "created_at" }
func (c Dataset) Default() Dataset                             { return "dataset" }
func (c Developer) Default() Developer                         { return "developer" }
func (c Embedding) Default() Embedding                         { return "embedding" }
func (c File) Default() File                                   { return "file" }
func (c FileCitation) Default() FileCitation                   { return "file_citation" }
func (c FilePath) Default() FilePath                           { return "file_path" }
func (c FileSearch) Default() FileSearch                       { return "file_search" }
func (c FileSearchCall) Default() FileSearchCall               { return "file_search_call" }
func (c Function) Default() Function                           { return "function" }
func (c FunctionCall) Default() FunctionCall                   { return "function_call" }
func (c FunctionCallOutput) Default() FunctionCallOutput       { return "function_call_output" }
func (c Greedy) Default() Greedy                               { return "greedy" }
func (c Image) Default() Image                                 { return "image" }
func (c ImageURL) Default() ImageURL                           { return "image_url" }
func (c InputFile) Default() InputFile                         { return "input_file" }
func (c InputImage) Default() InputImage                       { return "input_image" }
func (c InputText) Default() InputText                         { return "input_text" }
func (c JsonObject) Default() JsonObject                       { return "json_object" }
func (c JsonSchema) Default() JsonSchema                       { return "json_schema" }
func (c List) Default() List                                   { return "list" }
func (c LlmAsJudge) Default() LlmAsJudge                       { return "llm_as_judge" }
func (c LoRa) Default() LoRa                                   { return "LoRA" }
func (c Mcp) Default() Mcp                                     { return "mcp" }
func (c McpApprovalRequest) Default() McpApprovalRequest       { return "mcp_approval_request" }
func (c McpApprovalResponse) Default() McpApprovalResponse     { return "mcp_approval_response" }
func (c McpCall) Default() McpCall                             { return "mcp_call" }
func (c McpListTools) Default() McpListTools                   { return "mcp_list_tools" }
func (c Message) Default() Message                             { return "message" }
func (c Model) Default() Model                                 { return "model" }
func (c OutputText) Default() OutputText                       { return "output_text" }
func (c Qat) Default() Qat                                     { return "QAT" }
func (c ReasoningText) Default() ReasoningText                 { return "reasoning_text" }
func (c Refusal) Default() Refusal                             { return "refusal" }
func (c RegexParser) Default() RegexParser                     { return "regex_parser" }
func (c Response) Default() Response                           { return "response" }
func (c ResponseCompleted) Default() ResponseCompleted         { return "response.completed" }
func (c ResponseContentPartAdded) Default() ResponseContentPartAdded {
	return "response.content_part.added"
}
func (c ResponseContentPartDone) Default() ResponseContentPartDone {
	return "response.content_part.done"
}
func (c ResponseCreated) Default() ResponseCreated { return "response.created" }
func (c ResponseFailed) Default() ResponseFailed   { return "response.failed" }
func (c ResponseFileSearchCallCompleted) Default() ResponseFileSearchCallCompleted {
	return "response.file_search_call.completed"
}
func (c ResponseFileSearchCallInProgress) Default() ResponseFileSearchCallInProgress {
	return "response.file_search_call.in_progress"
}
func (c ResponseFileSearchCallSearching) Default() ResponseFileSearchCallSearching {
	return "response.file_search_call.searching"
}
func (c ResponseFunctionCallArgumentsDelta) Default() ResponseFunctionCallArgumentsDelta {
	return "response.function_call_arguments.delta"
}
func (c ResponseFunctionCallArgumentsDone) Default() ResponseFunctionCallArgumentsDone {
	return "response.function_call_arguments.done"
}
func (c ResponseInProgress) Default() ResponseInProgress { return "response.in_progress" }
func (c ResponseIncomplete) Default() ResponseIncomplete { return "response.incomplete" }
func (c ResponseMcpCallArgumentsDelta) Default() ResponseMcpCallArgumentsDelta {
	return "response.mcp_call.arguments.delta"
}
func (c ResponseMcpCallArgumentsDone) Default() ResponseMcpCallArgumentsDone {
	return "response.mcp_call.arguments.done"
}
func (c ResponseMcpCallCompleted) Default() ResponseMcpCallCompleted {
	return "response.mcp_call.completed"
}
func (c ResponseMcpCallFailed) Default() ResponseMcpCallFailed { return "response.mcp_call.failed" }
func (c ResponseMcpCallInProgress) Default() ResponseMcpCallInProgress {
	return "response.mcp_call.in_progress"
}
func (c ResponseMcpListToolsCompleted) Default() ResponseMcpListToolsCompleted {
	return "response.mcp_list_tools.completed"
}
func (c ResponseMcpListToolsFailed) Default() ResponseMcpListToolsFailed {
	return "response.mcp_list_tools.failed"
}
func (c ResponseMcpListToolsInProgress) Default() ResponseMcpListToolsInProgress {
	return "response.mcp_list_tools.in_progress"
}
func (c ResponseOutputItemAdded) Default() ResponseOutputItemAdded {
	return "response.output_item.added"
}
func (c ResponseOutputItemDone) Default() ResponseOutputItemDone { return "response.output_item.done" }
func (c ResponseOutputTextAnnotationAdded) Default() ResponseOutputTextAnnotationAdded {
	return "response.output_text.annotation.added"
}
func (c ResponseOutputTextDelta) Default() ResponseOutputTextDelta {
	return "response.output_text.delta"
}
func (c ResponseOutputTextDone) Default() ResponseOutputTextDone { return "response.output_text.done" }
func (c ResponseReasoningSummaryPartAdded) Default() ResponseReasoningSummaryPartAdded {
	return "response.reasoning_summary_part.added"
}
func (c ResponseReasoningSummaryPartDone) Default() ResponseReasoningSummaryPartDone {
	return "response.reasoning_summary_part.done"
}
func (c ResponseReasoningSummaryTextDelta) Default() ResponseReasoningSummaryTextDelta {
	return "response.reasoning_summary_text.delta"
}
func (c ResponseReasoningSummaryTextDone) Default() ResponseReasoningSummaryTextDone {
	return "response.reasoning_summary_text.done"
}
func (c ResponseReasoningTextDelta) Default() ResponseReasoningTextDelta {
	return "response.reasoning_text.delta"
}
func (c ResponseReasoningTextDone) Default() ResponseReasoningTextDone {
	return "response.reasoning_text.done"
}
func (c ResponseRefusalDelta) Default() ResponseRefusalDelta { return "response.refusal.delta" }
func (c ResponseRefusalDone) Default() ResponseRefusalDone   { return "response.refusal.done" }
func (c ResponseWebSearchCallCompleted) Default() ResponseWebSearchCallCompleted {
	return "response.web_search_call.completed"
}
func (c ResponseWebSearchCallInProgress) Default() ResponseWebSearchCallInProgress {
	return "response.web_search_call.in_progress"
}
func (c ResponseWebSearchCallSearching) Default() ResponseWebSearchCallSearching {
	return "response.web_search_call.searching"
}
func (c Rows) Default() Rows                       { return "rows" }
func (c ScoringFunction) Default() ScoringFunction { return "scoring_function" }
func (c Shield) Default() Shield                   { return "shield" }
func (c Static) Default() Static                   { return "static" }
func (c SummaryText) Default() SummaryText         { return "summary_text" }
func (c System) Default() System                   { return "system" }
func (c Text) Default() Text                       { return "text" }
func (c TextCompletion) Default() TextCompletion   { return "text_completion" }
func (c Tool) Default() Tool                       { return "tool" }
func (c ToolGroup) Default() ToolGroup             { return "tool_group" }
func (c TopK) Default() TopK                       { return "top_k" }
func (c TopP) Default() TopP                       { return "top_p" }
func (c Uri) Default() Uri                         { return "uri" }
func (c URLCitation) Default() URLCitation         { return "url_citation" }
func (c User) Default() User                       { return "user" }
func (c VectorStoreFileContentPage) Default() VectorStoreFileContentPage {
	return "vector_store.file_content.page"
}
func (c WebSearchCall) Default() WebSearchCall { return "web_search_call" }

func (c Assistant) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Auto) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Basic) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Benchmark) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c ChatCompletion) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c ChatCompletionChunk) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ContainerFileCitation) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Conversation) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c CreatedAt) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Dataset) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c Developer) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Embedding) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c FileCitation) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c FilePath) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c FileSearch) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c FileSearchCall) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Function) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c FunctionCall) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c FunctionCallOutput) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Greedy) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Image) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c ImageURL) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c InputFile) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c InputImage) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c InputText) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c JsonObject) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c JsonSchema) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c LlmAsJudge) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c LoRa) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Mcp) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c McpApprovalRequest) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c McpApprovalResponse) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c McpCall) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c McpListTools) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c Message) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c Model) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c OutputText) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Qat) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c ReasoningText) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c Refusal) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c RegexParser) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Response) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c ResponseCompleted) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c ResponseContentPartAdded) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ResponseContentPartDone) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ResponseCreated) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c ResponseFailed) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c ResponseFileSearchCallCompleted) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ResponseFileSearchCallInProgress) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ResponseFileSearchCallSearching) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ResponseFunctionCallArgumentsDelta) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c ResponseFunctionCallArgumentsDone) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c ResponseInProgress) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c ResponseIncomplete) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c ResponseMcpCallArgumentsDelta) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ResponseMcpCallArgumentsDone) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c ResponseMcpCallCompleted) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ResponseMcpCallFailed) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c ResponseMcpCallInProgress) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c ResponseMcpListToolsCompleted) MarshalJSON() ([]byte, error)      { return marshalString(c) }
func (c ResponseMcpListToolsFailed) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ResponseMcpListToolsInProgress) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ResponseOutputItemAdded) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ResponseOutputItemDone) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c ResponseOutputTextAnnotationAdded) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c ResponseOutputTextDelta) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ResponseOutputTextDone) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c ResponseReasoningSummaryPartAdded) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c ResponseReasoningSummaryPartDone) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ResponseReasoningSummaryTextDelta) MarshalJSON() ([]byte, error)  { return marshalString(c) }
func (c ResponseReasoningSummaryTextDone) MarshalJSON() ([]byte, error)   { return marshalString(c) }
func (c ResponseReasoningTextDelta) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c ResponseReasoningTextDone) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c ResponseRefusalDelta) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c ResponseRefusalDone) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ResponseWebSearchCallCompleted) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ResponseWebSearchCallInProgress) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ResponseWebSearchCallSearching) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Rows) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c ScoringFunction) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Shield) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Static) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c SummaryText) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c TextCompletion) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Tool) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c ToolGroup) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c TopK) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c TopP) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Uri) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c URLCitation) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c User) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c VectorStoreFileContentPage) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c WebSearchCall) MarshalJSON() ([]byte, error)                      { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
