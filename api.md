# Shared Params Types

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InterleavedContentUnionParam">InterleavedContentUnionParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InterleavedContentItemUnionParam">InterleavedContentItemUnionParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SamplingParams">SamplingParams</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SystemMessageParam">SystemMessageParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InterleavedContentUnion">InterleavedContentUnion</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InterleavedContentItemUnion">InterleavedContentItemUnion</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SafetyViolation">SafetyViolation</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringResult">ScoringResult</a>

# Toolgroups

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListToolGroupsResponse">ListToolGroupsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolGroup">ToolGroup</a>

Methods:

- <code title="get /v1/toolgroups">client.Toolgroups.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolgroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolGroup">ToolGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/toolgroups/{toolgroup_id}">client.Toolgroups.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolgroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolgroupID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolGroup">ToolGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

Methods:

- <code title="get /v1/tools">client.Tools.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolListParams">ToolListParams</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolDef">ToolDef</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tools/{tool_name}">client.Tools.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolDef">ToolDef</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ToolRuntime

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolDef">ToolDef</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolInvocationResult">ToolInvocationResult</a>

Methods:

- <code title="post /v1/tool-runtime/invoke">client.ToolRuntime.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolRuntimeService.InvokeTool">InvokeTool</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolRuntimeInvokeToolParams">ToolRuntimeInvokeToolParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolInvocationResult">ToolInvocationResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/tool-runtime/list-tools">client.ToolRuntime.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolRuntimeService.ListTools">ListTools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolRuntimeListToolsParams">ToolRuntimeListToolsParams</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolDef">ToolDef</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## RagTool

# Responses

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseObject">ResponseObject</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseObjectStreamUnion">ResponseObjectStreamUnion</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseListResponse">ResponseListResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseDeleteResponse">ResponseDeleteResponse</a>

Methods:

- <code title="post /v1/responses">client.Responses.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseNewParams">ResponseNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseObject">ResponseObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/responses/{response_id}">client.Responses.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseObject">ResponseObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/responses">client.Responses.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseListParams">ResponseListParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseListResponse">ResponseListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/responses/{response_id}">client.Responses.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseDeleteResponse">ResponseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InputItems

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseInputItemListResponse">ResponseInputItemListResponse</a>

Methods:

- <code title="get /v1/responses/{response_id}/input_items">client.Responses.InputItems.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseInputItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseInputItemListParams">ResponseInputItemListParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ResponseInputItemListResponse">ResponseInputItemListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Prompts

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListPromptsResponse">ListPromptsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Prompt">Prompt</a>

Methods:

- <code title="post /v1/prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptNewParams">PromptNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptGetParams">PromptGetParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptUpdateParams">PromptUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/prompts/{prompt_id}/set-default-version">client.Prompts.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptService.SetDefaultVersion">SetDefaultVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptSetDefaultVersionParams">PromptSetDefaultVersionParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Methods:

- <code title="get /v1/prompts/{prompt_id}/versions">client.Prompts.Versions.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PromptVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Conversations

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationObject">ConversationObject</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationDeleteResponse">ConversationDeleteResponse</a>

Methods:

- <code title="post /v1/conversations">client.Conversations.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationNewParams">ConversationNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationObject">ConversationObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/conversations/{conversation_id}">client.Conversations.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationObject">ConversationObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/conversations/{conversation_id}">client.Conversations.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationUpdateParams">ConversationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationObject">ConversationObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/conversations/{conversation_id}">client.Conversations.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationDeleteResponse">ConversationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Items

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemNewResponse">ConversationItemNewResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemListResponseUnion">ConversationItemListResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemGetResponseUnion">ConversationItemGetResponseUnion</a>

Methods:

- <code title="post /v1/conversations/{conversation_id}/items">client.Conversations.Items.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemNewParams">ConversationItemNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemNewResponse">ConversationItemNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/conversations/{conversation_id}/items">client.Conversations.Items.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemListParams">ConversationItemListParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemListResponseUnion">ConversationItemListResponseUnion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/conversations/{conversation_id}/items/{item_id}">client.Conversations.Items.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemGetParams">ConversationItemGetParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ConversationItemGetResponseUnion">ConversationItemGetResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inspect

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#HealthInfo">HealthInfo</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ProviderInfo">ProviderInfo</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#RouteInfo">RouteInfo</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VersionInfo">VersionInfo</a>

Methods:

- <code title="get /v1/health">client.Inspect.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InspectService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#HealthInfo">HealthInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/version">client.Inspect.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InspectService.Version">Version</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VersionInfo">VersionInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CreateEmbeddingsResponse">CreateEmbeddingsResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CreateEmbeddingsResponse">CreateEmbeddingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionChunk">ChatCompletionChunk</a>

## Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionNewResponseUnion">ChatCompletionNewResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionGetResponse">ChatCompletionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionListResponse">ChatCompletionListResponse</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionNewResponseUnion">ChatCompletionNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chat/completions/{completion_id}">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, completionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionGetResponse">ChatCompletionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionListParams">ChatCompletionListParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ChatCompletionListResponse">ChatCompletionListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CompletionNewResponse">CompletionNewResponse</a>

Methods:

- <code title="post /v1/completions">client.Completions.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CompletionNewParams">CompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CompletionNewResponse">CompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VectorIo

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#QueryChunksResponse">QueryChunksResponse</a>

Methods:

- <code title="post /v1/vector-io/insert">client.VectorIo.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorIoService.Insert">Insert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorIoInsertParams">VectorIoInsertParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/vector-io/query">client.VectorIo.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorIoService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorIoQueryParams">VectorIoQueryParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#QueryChunksResponse">QueryChunksResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VectorStores

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListVectorStoresResponse">ListVectorStoresResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStore">VectorStore</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreDeleteResponse">VectorStoreDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreSearchResponse">VectorStoreSearchResponse</a>

Methods:

- <code title="post /v1/vector_stores">client.VectorStores.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreNewParams">VectorStoreNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStore">VectorStore</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}">client.VectorStores.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStore">VectorStore</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector_stores/{vector_store_id}">client.VectorStores.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreUpdateParams">VectorStoreUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStore">VectorStore</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores">client.VectorStores.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreListParams">VectorStoreListParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStore">VectorStore</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/vector_stores/{vector_store_id}">client.VectorStores.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreDeleteResponse">VectorStoreDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector_stores/{vector_store_id}/search">client.VectorStores.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreSearchParams">VectorStoreSearchParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreSearchResponse">VectorStoreSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Files

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFile">VectorStoreFile</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileDeleteResponse">VectorStoreFileDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileContentResponse">VectorStoreFileContentResponse</a>

Methods:

- <code title="post /v1/vector_stores/{vector_store_id}/files">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileNewParams">VectorStoreFileNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/files/{file_id}">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileGetParams">VectorStoreFileGetParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector_stores/{vector_store_id}/files/{file_id}">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileUpdateParams">VectorStoreFileUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/files">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileListParams">VectorStoreFileListParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFile">VectorStoreFile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/vector_stores/{vector_store_id}/files/{file_id}">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileDeleteParams">VectorStoreFileDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileDeleteResponse">VectorStoreFileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/files/{file_id}/content">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileContentParams">VectorStoreFileContentParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileContentResponse">VectorStoreFileContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FileBatches

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListVectorStoreFilesInBatchResponse">ListVectorStoreFilesInBatchResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatches">VectorStoreFileBatches</a>

Methods:

- <code title="post /v1/vector_stores/{vector_store_id}/file_batches">client.VectorStores.FileBatches.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatchNewParams">VectorStoreFileBatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatches">VectorStoreFileBatches</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/file_batches/{batch_id}">client.VectorStores.FileBatches.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatchGetParams">VectorStoreFileBatchGetParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatches">VectorStoreFileBatches</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector_stores/{vector_store_id}/file_batches/{batch_id}/cancel">client.VectorStores.FileBatches.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatchService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatchCancelParams">VectorStoreFileBatchCancelParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatches">VectorStoreFileBatches</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/file_batches/{batch_id}/files">client.VectorStores.FileBatches.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatchService.ListFiles">ListFiles</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileBatchListFilesParams">VectorStoreFileBatchListFilesParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFile">VectorStoreFile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListModelsResponse">ListModelsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelGetResponse">ModelGetResponse</a>

Methods:

- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelGetResponse">ModelGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OpenAI

Methods:

- <code title="get /v1/models">client.Models.OpenAI.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelOpenAIService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListProvidersResponse">ListProvidersResponse</a>

Methods:

- <code title="get /v1/providers/{provider_id}">client.Providers.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ProviderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ProviderInfo">ProviderInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/providers">client.Providers.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ProviderInfo">ProviderInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Routes

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListRoutesResponse">ListRoutesResponse</a>

Methods:

- <code title="get /v1/inspect/routes">client.Routes.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#RouteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#RouteListParams">RouteListParams</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#RouteInfo">RouteInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Moderations

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CreateResponse">CreateResponse</a>

Methods:

- <code title="post /v1/moderations">client.Moderations.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModerationNewParams">ModerationNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CreateResponse">CreateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Safety

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#RunShieldResponse">RunShieldResponse</a>

Methods:

- <code title="post /v1/safety/run-shield">client.Safety.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SafetyService.RunShield">RunShield</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SafetyRunShieldParams">SafetyRunShieldParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#RunShieldResponse">RunShieldResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Shields

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListShieldsResponse">ListShieldsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Shield">Shield</a>

Methods:

- <code title="get /v1/shields/{identifier}">client.Shields.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ShieldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/shields">client.Shields.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ShieldService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Scoring

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringScoreResponse">ScoringScoreResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringScoreBatchResponse">ScoringScoreBatchResponse</a>

Methods:

- <code title="post /v1/scoring/score">client.Scoring.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringService.Score">Score</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringScoreParams">ScoringScoreParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringScoreResponse">ScoringScoreResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/scoring/score-batch">client.Scoring.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringService.ScoreBatch">ScoreBatch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringScoreBatchParams">ScoringScoreBatchParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringScoreBatchResponse">ScoringScoreBatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ScoringFunctions

Params Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFnParamsUnion">ScoringFnParamsUnion</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListScoringFunctionsResponse">ListScoringFunctionsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFn">ScoringFn</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFnParamsUnionResp">ScoringFnParamsUnionResp</a>

Methods:

- <code title="get /v1/scoring-functions/{scoring_fn_id}">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, scoringFnID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFn">ScoringFn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/scoring-functions">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFn">ScoringFn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#DeleteFileResponse">DeleteFileResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#File">File</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListFilesResponse">ListFilesResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileContentResponse">FileContentResponse</a>

Methods:

- <code title="post /v1/files">client.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileNewParams">FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files">client.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#File">File</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#DeleteFileResponse">DeleteFileResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/{file_id}/content">client.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#FileContentResponse">FileContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Alpha

## Inference

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaInferenceRerankResponse">AlphaInferenceRerankResponse</a>

Methods:

- <code title="post /v1alpha/inference/rerank">client.Alpha.Inference.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaInferenceService.Rerank">Rerank</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaInferenceRerankParams">AlphaInferenceRerankParams</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaInferenceRerankResponse">AlphaInferenceRerankResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PostTraining

Params Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlgorithmConfigUnionParam">AlgorithmConfigUnionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListPostTrainingJobsResponse">ListPostTrainingJobsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PostTrainingJob">PostTrainingJob</a>

Methods:

- <code title="post /v1alpha/post-training/preference-optimize">client.Alpha.PostTraining.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingService.PreferenceOptimize">PreferenceOptimize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingPreferenceOptimizeParams">AlphaPostTrainingPreferenceOptimizeParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PostTrainingJob">PostTrainingJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1alpha/post-training/supervised-fine-tune">client.Alpha.PostTraining.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingService.SupervisedFineTune">SupervisedFineTune</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingSupervisedFineTuneParams">AlphaPostTrainingSupervisedFineTuneParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#PostTrainingJob">PostTrainingJob</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Job

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobArtifactsResponse">AlphaPostTrainingJobArtifactsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobStatusResponse">AlphaPostTrainingJobStatusResponse</a>

Methods:

- <code title="get /v1alpha/post-training/jobs">client.Alpha.PostTraining.Job.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListPostTrainingJobsResponseData">ListPostTrainingJobsResponseData</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1alpha/post-training/job/artifacts">client.Alpha.PostTraining.Job.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobService.Artifacts">Artifacts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobArtifactsParams">AlphaPostTrainingJobArtifactsParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobArtifactsResponse">AlphaPostTrainingJobArtifactsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1alpha/post-training/job/cancel">client.Alpha.PostTraining.Job.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobCancelParams">AlphaPostTrainingJobCancelParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1alpha/post-training/job/status">client.Alpha.PostTraining.Job.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobStatusParams">AlphaPostTrainingJobStatusParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaPostTrainingJobStatusResponse">AlphaPostTrainingJobStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Benchmarks

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Benchmark">Benchmark</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListBenchmarksResponse">ListBenchmarksResponse</a>

Methods:

- <code title="get /v1alpha/eval/benchmarks/{benchmark_id}">client.Alpha.Benchmarks.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaBenchmarkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Benchmark">Benchmark</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1alpha/eval/benchmarks">client.Alpha.Benchmarks.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaBenchmarkService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Benchmark">Benchmark</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1alpha/eval/benchmarks">client.Alpha.Benchmarks.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaBenchmarkService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaBenchmarkRegisterParams">AlphaBenchmarkRegisterParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Eval

Params Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BenchmarkConfigParam">BenchmarkConfigParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#EvaluateResponse">EvaluateResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Job">Job</a>

Methods:

- <code title="post /v1alpha/eval/benchmarks/{benchmark_id}/evaluations">client.Alpha.Eval.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalService.EvaluateRows">EvaluateRows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalEvaluateRowsParams">AlphaEvalEvaluateRowsParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#EvaluateResponse">EvaluateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1alpha/eval/benchmarks/{benchmark_id}/evaluations">client.Alpha.Eval.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalService.EvaluateRowsAlpha">EvaluateRowsAlpha</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalEvaluateRowsAlphaParams">AlphaEvalEvaluateRowsAlphaParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#EvaluateResponse">EvaluateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1alpha/eval/benchmarks/{benchmark_id}/jobs">client.Alpha.Eval.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalService.RunEval">RunEval</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalRunEvalParams">AlphaEvalRunEvalParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1alpha/eval/benchmarks/{benchmark_id}/jobs">client.Alpha.Eval.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalService.RunEvalAlpha">RunEvalAlpha</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, benchmarkID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalRunEvalAlphaParams">AlphaEvalRunEvalAlphaParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Jobs

Methods:

- <code title="get /v1alpha/eval/benchmarks/{benchmark_id}/jobs/{job_id}/result">client.Alpha.Eval.Jobs.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalJobGetParams">AlphaEvalJobGetParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#EvaluateResponse">EvaluateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1alpha/eval/benchmarks/{benchmark_id}/jobs/{job_id}">client.Alpha.Eval.Jobs.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalJobService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalJobCancelParams">AlphaEvalJobCancelParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1alpha/eval/benchmarks/{benchmark_id}/jobs/{job_id}">client.Alpha.Eval.Jobs.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalJobService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#AlphaEvalJobStatusParams">AlphaEvalJobStatusParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Beta

## Datasets

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListDatasetsResponse">ListDatasetsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetGetResponse">BetaDatasetGetResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetIterrowsResponse">BetaDatasetIterrowsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetRegisterResponse">BetaDatasetRegisterResponse</a>

Methods:

- <code title="get /v1beta/datasets/{dataset_id}">client.Beta.Datasets.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetGetResponse">BetaDatasetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1beta/datasets">client.Beta.Datasets.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListDatasetsResponseData">ListDatasetsResponseData</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1beta/datasetio/append-rows/{dataset_id}">client.Beta.Datasets.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetService.Appendrows">Appendrows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetAppendrowsParams">BetaDatasetAppendrowsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /v1beta/datasetio/iterrows/{dataset_id}">client.Beta.Datasets.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetService.Iterrows">Iterrows</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetIterrowsParams">BetaDatasetIterrowsParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetIterrowsResponse">BetaDatasetIterrowsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1beta/datasets">client.Beta.Datasets.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetRegisterParams">BetaDatasetRegisterParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetRegisterResponse">BetaDatasetRegisterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1beta/datasets/{dataset_id}">client.Beta.Datasets.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#BetaDatasetService.Unregister">Unregister</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, datasetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
