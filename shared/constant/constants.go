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

type Agent string                              // Always "agent"
type AgentTurnInput string                     // Always "agent_turn_input"
type Array string                              // Always "array"
type Assistant string                          // Always "assistant"
type Auto string                               // Always "auto"
type Basic string                              // Always "basic"
type Benchmark string                          // Always "benchmark"
type Boolean string                            // Always "boolean"
type ChatCompletionInput string                // Always "chat_completion_input"
type ChatCompletion string                     // Always "chat.completion"
type ChatCompletionChunk string                // Always "chat.completion.chunk"
type CompletionInput string                    // Always "completion_input"
type ContainerFileCitation string              // Always "container_file_citation"
type CreatedAt string                          // Always "created_at"
type Dataset string                            // Always "dataset"
type Default string                            // Always "default"
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
type Grammar string                            // Always "grammar"
type Greedy string                             // Always "greedy"
type Image string                              // Always "image"
type ImageURL string                           // Always "image_url"
type Inference string                          // Always "inference"
type InputImage string                         // Always "input_image"
type InputText string                          // Always "input_text"
type Json string                               // Always "json"
type JsonObject string                         // Always "json_object"
type JsonSchema string                         // Always "json_schema"
type List string                               // Always "list"
type Llm string                                // Always "llm"
type LlmAsJudge string                         // Always "llm_as_judge"
type LoRa string                               // Always "LoRA"
type Mcp string                                // Always "mcp"
type McpCall string                            // Always "mcp_call"
type McpListTools string                       // Always "mcp_list_tools"
type MemoryRetrieval string                    // Always "memory_retrieval"
type Message string                            // Always "message"
type Metric string                             // Always "metric"
type Model string                              // Always "model"
type Number string                             // Always "number"
type Object string                             // Always "object"
type OutputText string                         // Always "output_text"
type Qat string                                // Always "QAT"
type Refusal string                            // Always "refusal"
type RegexParser string                        // Always "regex_parser"
type Response string                           // Always "response"
type ResponseCompleted string                  // Always "response.completed"
type ResponseContentPartAdded string           // Always "response.content_part.added"
type ResponseContentPartDone string            // Always "response.content_part.done"
type ResponseCreated string                    // Always "response.created"
type ResponseFunctionCallArgumentsDelta string // Always "response.function_call_arguments.delta"
type ResponseFunctionCallArgumentsDone string  // Always "response.function_call_arguments.done"
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
type ResponseOutputTextDelta string            // Always "response.output_text.delta"
type ResponseOutputTextDone string             // Always "response.output_text.done"
type ResponseWebSearchCallCompleted string     // Always "response.web_search_call.completed"
type ResponseWebSearchCallInProgress string    // Always "response.web_search_call.in_progress"
type ResponseWebSearchCallSearching string     // Always "response.web_search_call.searching"
type Rows string                               // Always "rows"
type Rrf string                                // Always "rrf"
type ScoringFunction string                    // Always "scoring_function"
type Shield string                             // Always "shield"
type ShieldCall string                         // Always "shield_call"
type SpanEnd string                            // Always "span_end"
type SpanStart string                          // Always "span_start"
type Static string                             // Always "static"
type StepComplete string                       // Always "step_complete"
type StepProgress string                       // Always "step_progress"
type StepStart string                          // Always "step_start"
type String string                             // Always "string"
type StructuredLog string                      // Always "structured_log"
type System string                             // Always "system"
type Text string                               // Always "text"
type TextCompletion string                     // Always "text_completion"
type Tool string                               // Always "tool"
type ToolCall string                           // Always "tool_call"
type ToolExecution string                      // Always "tool_execution"
type ToolGroup string                          // Always "tool_group"
type TopK string                               // Always "top_k"
type TopP string                               // Always "top_p"
type TurnAwaitingInput string                  // Always "turn_awaiting_input"
type TurnComplete string                       // Always "turn_complete"
type TurnStart string                          // Always "turn_start"
type Union string                              // Always "union"
type UnstructuredLog string                    // Always "unstructured_log"
type Uri string                                // Always "uri"
type URLCitation string                        // Always "url_citation"
type User string                               // Always "user"
type VectorDB string                           // Always "vector_db"
type WebSearchCall string                      // Always "web_search_call"
type Weighted string                           // Always "weighted"

func (c Agent) Default() Agent                                 { return "agent" }
func (c AgentTurnInput) Default() AgentTurnInput               { return "agent_turn_input" }
func (c Array) Default() Array                                 { return "array" }
func (c Assistant) Default() Assistant                         { return "assistant" }
func (c Auto) Default() Auto                                   { return "auto" }
func (c Basic) Default() Basic                                 { return "basic" }
func (c Benchmark) Default() Benchmark                         { return "benchmark" }
func (c Boolean) Default() Boolean                             { return "boolean" }
func (c ChatCompletionInput) Default() ChatCompletionInput     { return "chat_completion_input" }
func (c ChatCompletion) Default() ChatCompletion               { return "chat.completion" }
func (c ChatCompletionChunk) Default() ChatCompletionChunk     { return "chat.completion.chunk" }
func (c CompletionInput) Default() CompletionInput             { return "completion_input" }
func (c ContainerFileCitation) Default() ContainerFileCitation { return "container_file_citation" }
func (c CreatedAt) Default() CreatedAt                         { return "created_at" }
func (c Dataset) Default() Dataset                             { return "dataset" }
func (c Default) Default() Default                             { return "default" }
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
func (c Grammar) Default() Grammar                             { return "grammar" }
func (c Greedy) Default() Greedy                               { return "greedy" }
func (c Image) Default() Image                                 { return "image" }
func (c ImageURL) Default() ImageURL                           { return "image_url" }
func (c Inference) Default() Inference                         { return "inference" }
func (c InputImage) Default() InputImage                       { return "input_image" }
func (c InputText) Default() InputText                         { return "input_text" }
func (c Json) Default() Json                                   { return "json" }
func (c JsonObject) Default() JsonObject                       { return "json_object" }
func (c JsonSchema) Default() JsonSchema                       { return "json_schema" }
func (c List) Default() List                                   { return "list" }
func (c Llm) Default() Llm                                     { return "llm" }
func (c LlmAsJudge) Default() LlmAsJudge                       { return "llm_as_judge" }
func (c LoRa) Default() LoRa                                   { return "LoRA" }
func (c Mcp) Default() Mcp                                     { return "mcp" }
func (c McpCall) Default() McpCall                             { return "mcp_call" }
func (c McpListTools) Default() McpListTools                   { return "mcp_list_tools" }
func (c MemoryRetrieval) Default() MemoryRetrieval             { return "memory_retrieval" }
func (c Message) Default() Message                             { return "message" }
func (c Metric) Default() Metric                               { return "metric" }
func (c Model) Default() Model                                 { return "model" }
func (c Number) Default() Number                               { return "number" }
func (c Object) Default() Object                               { return "object" }
func (c OutputText) Default() OutputText                       { return "output_text" }
func (c Qat) Default() Qat                                     { return "QAT" }
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
func (c ResponseFunctionCallArgumentsDelta) Default() ResponseFunctionCallArgumentsDelta {
	return "response.function_call_arguments.delta"
}
func (c ResponseFunctionCallArgumentsDone) Default() ResponseFunctionCallArgumentsDone {
	return "response.function_call_arguments.done"
}
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
func (c ResponseOutputTextDelta) Default() ResponseOutputTextDelta {
	return "response.output_text.delta"
}
func (c ResponseOutputTextDone) Default() ResponseOutputTextDone { return "response.output_text.done" }
func (c ResponseWebSearchCallCompleted) Default() ResponseWebSearchCallCompleted {
	return "response.web_search_call.completed"
}
func (c ResponseWebSearchCallInProgress) Default() ResponseWebSearchCallInProgress {
	return "response.web_search_call.in_progress"
}
func (c ResponseWebSearchCallSearching) Default() ResponseWebSearchCallSearching {
	return "response.web_search_call.searching"
}
func (c Rows) Default() Rows                           { return "rows" }
func (c Rrf) Default() Rrf                             { return "rrf" }
func (c ScoringFunction) Default() ScoringFunction     { return "scoring_function" }
func (c Shield) Default() Shield                       { return "shield" }
func (c ShieldCall) Default() ShieldCall               { return "shield_call" }
func (c SpanEnd) Default() SpanEnd                     { return "span_end" }
func (c SpanStart) Default() SpanStart                 { return "span_start" }
func (c Static) Default() Static                       { return "static" }
func (c StepComplete) Default() StepComplete           { return "step_complete" }
func (c StepProgress) Default() StepProgress           { return "step_progress" }
func (c StepStart) Default() StepStart                 { return "step_start" }
func (c String) Default() String                       { return "string" }
func (c StructuredLog) Default() StructuredLog         { return "structured_log" }
func (c System) Default() System                       { return "system" }
func (c Text) Default() Text                           { return "text" }
func (c TextCompletion) Default() TextCompletion       { return "text_completion" }
func (c Tool) Default() Tool                           { return "tool" }
func (c ToolCall) Default() ToolCall                   { return "tool_call" }
func (c ToolExecution) Default() ToolExecution         { return "tool_execution" }
func (c ToolGroup) Default() ToolGroup                 { return "tool_group" }
func (c TopK) Default() TopK                           { return "top_k" }
func (c TopP) Default() TopP                           { return "top_p" }
func (c TurnAwaitingInput) Default() TurnAwaitingInput { return "turn_awaiting_input" }
func (c TurnComplete) Default() TurnComplete           { return "turn_complete" }
func (c TurnStart) Default() TurnStart                 { return "turn_start" }
func (c Union) Default() Union                         { return "union" }
func (c UnstructuredLog) Default() UnstructuredLog     { return "unstructured_log" }
func (c Uri) Default() Uri                             { return "uri" }
func (c URLCitation) Default() URLCitation             { return "url_citation" }
func (c User) Default() User                           { return "user" }
func (c VectorDB) Default() VectorDB                   { return "vector_db" }
func (c WebSearchCall) Default() WebSearchCall         { return "web_search_call" }
func (c Weighted) Default() Weighted                   { return "weighted" }

func (c Agent) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c AgentTurnInput) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Array) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Assistant) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Auto) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Basic) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Benchmark) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Boolean) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c ChatCompletionInput) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c ChatCompletion) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c ChatCompletionChunk) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c CompletionInput) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c ContainerFileCitation) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c CreatedAt) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Dataset) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c Default) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
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
func (c Grammar) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c Greedy) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Image) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c ImageURL) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c Inference) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c InputImage) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c InputText) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Json) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c JsonObject) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c JsonSchema) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c List) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Llm) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c LlmAsJudge) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c LoRa) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Mcp) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c McpCall) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c McpListTools) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c MemoryRetrieval) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Message) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c Metric) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Model) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c Number) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Object) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c OutputText) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c Qat) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c Refusal) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c RegexParser) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c Response) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c ResponseCompleted) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c ResponseContentPartAdded) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c ResponseContentPartDone) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ResponseCreated) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c ResponseFunctionCallArgumentsDelta) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c ResponseFunctionCallArgumentsDone) MarshalJSON() ([]byte, error)  { return marshalString(c) }
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
func (c ResponseOutputTextDelta) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c ResponseOutputTextDone) MarshalJSON() ([]byte, error)             { return marshalString(c) }
func (c ResponseWebSearchCallCompleted) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c ResponseWebSearchCallInProgress) MarshalJSON() ([]byte, error)    { return marshalString(c) }
func (c ResponseWebSearchCallSearching) MarshalJSON() ([]byte, error)     { return marshalString(c) }
func (c Rows) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c Rrf) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c ScoringFunction) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Shield) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c ShieldCall) MarshalJSON() ([]byte, error)                         { return marshalString(c) }
func (c SpanEnd) MarshalJSON() ([]byte, error)                            { return marshalString(c) }
func (c SpanStart) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Static) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c StepComplete) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c StepProgress) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c StepStart) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c String) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c StructuredLog) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)                             { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c TextCompletion) MarshalJSON() ([]byte, error)                     { return marshalString(c) }
func (c Tool) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c ToolCall) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c ToolExecution) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c ToolGroup) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c TopK) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c TopP) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c TurnAwaitingInput) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c TurnComplete) MarshalJSON() ([]byte, error)                       { return marshalString(c) }
func (c TurnStart) MarshalJSON() ([]byte, error)                          { return marshalString(c) }
func (c Union) MarshalJSON() ([]byte, error)                              { return marshalString(c) }
func (c UnstructuredLog) MarshalJSON() ([]byte, error)                    { return marshalString(c) }
func (c Uri) MarshalJSON() ([]byte, error)                                { return marshalString(c) }
func (c URLCitation) MarshalJSON() ([]byte, error)                        { return marshalString(c) }
func (c User) MarshalJSON() ([]byte, error)                               { return marshalString(c) }
func (c VectorDB) MarshalJSON() ([]byte, error)                           { return marshalString(c) }
func (c WebSearchCall) MarshalJSON() ([]byte, error)                      { return marshalString(c) }
func (c Weighted) MarshalJSON() ([]byte, error)                           { return marshalString(c) }

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
