# Shared Params Types

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#CompletionMessageParam">CompletionMessageParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#DocumentParam">DocumentParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InterleavedContentUnionParam">InterleavedContentUnionParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InterleavedContentItemUnionParam">InterleavedContentItemUnionParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#MessageUnionParam">MessageUnionParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#QueryConfigParam">QueryConfigParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SystemMessageParam">SystemMessageParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolCallParam">ToolCallParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolResponseMessageParam">ToolResponseMessageParam</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#UserMessageParam">UserMessageParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InterleavedContentUnion">InterleavedContentUnion</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#InterleavedContentItemUnion">InterleavedContentItemUnion</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#QueryResult">QueryResult</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SafetyViolation">SafetyViolation</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringResult">ScoringResult</a>

# Toolgroups

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ListToolGroupsResponse">ListToolGroupsResponse</a>
- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolGroup">ToolGroup</a>

Methods:

- <code title="get /v1/toolgroups">client.Toolgroups.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolgroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolGroup">ToolGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/toolgroups/{toolgroup_id}">client.Toolgroups.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolgroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolgroupID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolGroup">ToolGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/toolgroups">client.Toolgroups.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolgroupService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolgroupRegisterParams">ToolgroupRegisterParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /v1/toolgroups/{toolgroup_id}">client.Toolgroups.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolgroupService.Unregister">Unregister</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolgroupID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

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

Methods:

- <code title="post /v1/tool-runtime/rag-tool/insert">client.ToolRuntime.RagTool.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolRuntimeRagToolService.Insert">Insert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolRuntimeRagToolInsertParams">ToolRuntimeRagToolInsertParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/tool-runtime/rag-tool/query">client.ToolRuntime.RagTool.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolRuntimeRagToolService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ToolRuntimeRagToolQueryParams">ToolRuntimeRagToolQueryParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#QueryResult">QueryResult</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
- <code title="get /v1/vector_stores/{vector_store_id}/files/{file_id}/content">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileContentParams">VectorStoreFileContentParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#VectorStoreFileContentResponse">VectorStoreFileContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

Methods:

- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelRegisterParams">ModelRegisterParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ModelService.Unregister">Unregister</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

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
- <code title="delete /v1/shields/{identifier}">client.Shields.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ShieldService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/shields">client.Shields.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ShieldService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ShieldRegisterParams">ShieldRegisterParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SyntheticDataGeneration

Response Types:

- <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SyntheticDataGenerationResponse">SyntheticDataGenerationResponse</a>

Methods:

- <code title="post /v1/synthetic-data-generation/generate">client.SyntheticDataGeneration.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SyntheticDataGenerationService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SyntheticDataGenerationGenerateParams">SyntheticDataGenerationGenerateParams</a>) (<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#SyntheticDataGenerationResponse">SyntheticDataGenerationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
- <code title="post /v1/scoring-functions">client.ScoringFunctions.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFunctionService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go">llamastackclient</a>.<a href="https://pkg.go.dev/github.com/llamastack/llama-stack-client-go#ScoringFunctionRegisterParams">ScoringFunctionRegisterParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

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

## PostTraining

### Job

## Benchmarks

## Eval

### Jobs

## Agents

### Session

### Steps

### Turn

# Beta

## Datasets
