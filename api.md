# Shared Response Types

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#HealthInfo">HealthInfo</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListProvidersResponse">ListProvidersResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListRoutesResponse">ListRoutesResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ProviderInfo">ProviderInfo</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#RouteInfo">RouteInfo</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#SafetyViolation">SafetyViolation</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VersionInfo">VersionInfo</a>

# Responses

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CompactedResponse">CompactedResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseMessage">ResponseMessage</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseObject">ResponseObject</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseObjectStreamUnion">ResponseObjectStreamUnion</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseListResponse">ResponseListResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseDeleteResponse">ResponseDeleteResponse</a>

Methods:

- <code title="post /v1/responses">client.Responses.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseNewParams">ResponseNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseObject">ResponseObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/responses/{response_id}">client.Responses.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseObject">ResponseObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/responses">client.Responses.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseListParams">ResponseListParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseListResponse">ResponseListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/responses/{response_id}">client.Responses.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseDeleteResponse">ResponseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/responses/compact">client.Responses.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseService.Compact">Compact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseCompactParams">ResponseCompactParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CompactedResponse">CompactedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InputItems

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseInputItemListResponse">ResponseInputItemListResponse</a>

Methods:

- <code title="get /v1/responses/{response_id}/input_items">client.Responses.InputItems.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseInputItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseInputItemListParams">ResponseInputItemListParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ResponseInputItemListResponse">ResponseInputItemListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Prompts

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListPromptsResponse">ListPromptsResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Prompt">Prompt</a>

Methods:

- <code title="post /v1/prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptNewParams">PromptNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptGetParams">PromptGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptUpdateParams">PromptUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/prompts">client.Prompts.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/prompts/{prompt_id}">client.Prompts.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /v1/prompts/{prompt_id}/set-default-version">client.Prompts.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptService.SetDefaultVersion">SetDefaultVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptSetDefaultVersionParams">PromptSetDefaultVersionParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Versions

Methods:

- <code title="get /v1/prompts/{prompt_id}/versions">client.Prompts.Versions.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#PromptVersionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, promptID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Prompt">Prompt</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Conversations

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationObject">ConversationObject</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationDeleteResponse">ConversationDeleteResponse</a>

Methods:

- <code title="post /v1/conversations">client.Conversations.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationNewParams">ConversationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationObject">ConversationObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/conversations/{conversation_id}">client.Conversations.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationObject">ConversationObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/conversations/{conversation_id}">client.Conversations.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationUpdateParams">ConversationUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationObject">ConversationObject</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/conversations/{conversation_id}">client.Conversations.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationDeleteResponse">ConversationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Items

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemNewResponse">ConversationItemNewResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemListResponseUnion">ConversationItemListResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemDeleteResponse">ConversationItemDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemGetResponseUnion">ConversationItemGetResponseUnion</a>

Methods:

- <code title="post /v1/conversations/{conversation_id}/items">client.Conversations.Items.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemNewParams">ConversationItemNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemNewResponse">ConversationItemNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/conversations/{conversation_id}/items">client.Conversations.Items.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, conversationID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemListParams">ConversationItemListParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemListResponseUnion">ConversationItemListResponseUnion</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/conversations/{conversation_id}/items/{item_id}">client.Conversations.Items.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemDeleteParams">ConversationItemDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemDeleteResponse">ConversationItemDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/conversations/{conversation_id}/items/{item_id}">client.Conversations.Items.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, itemID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemGetParams">ConversationItemGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ConversationItemGetResponseUnion">ConversationItemGetResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Inspect

Methods:

- <code title="get /v1/health">client.Inspect.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#InspectService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#HealthInfo">HealthInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/version">client.Inspect.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#InspectService.Version">Version</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VersionInfo">VersionInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CreateEmbeddingsResponse">CreateEmbeddingsResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#EmbeddingNewParams">EmbeddingNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CreateEmbeddingsResponse">CreateEmbeddingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionChunk">ChatCompletionChunk</a>

## Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionNewResponse">ChatCompletionNewResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionGetResponse">ChatCompletionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionListResponse">ChatCompletionListResponse</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionNewParams">ChatCompletionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionNewResponse">ChatCompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chat/completions/{completion_id}">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, completionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionGetResponse">ChatCompletionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionListParams">ChatCompletionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ChatCompletionListResponse">ChatCompletionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CompletionNewResponse">CompletionNewResponse</a>

Methods:

- <code title="post /v1/completions">client.Completions.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CompletionNewParams">CompletionNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CompletionNewResponse">CompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VectorIo

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#QueryChunksResponse">QueryChunksResponse</a>

Methods:

- <code title="post /v1/vector-io/insert">client.VectorIo.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorIoService.Insert">Insert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorIoInsertParams">VectorIoInsertParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/vector-io/query">client.VectorIo.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorIoService.Query">Query</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorIoQueryParams">VectorIoQueryParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#QueryChunksResponse">QueryChunksResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VectorStores

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListVectorStoresResponse">ListVectorStoresResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStore">VectorStore</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreDeleteResponse">VectorStoreDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreSearchResponse">VectorStoreSearchResponse</a>

Methods:

- <code title="post /v1/vector_stores">client.VectorStores.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreNewParams">VectorStoreNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStore">VectorStore</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}">client.VectorStores.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStore">VectorStore</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector_stores/{vector_store_id}">client.VectorStores.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreUpdateParams">VectorStoreUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStore">VectorStore</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores">client.VectorStores.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreListParams">VectorStoreListParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStore">VectorStore</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/vector_stores/{vector_store_id}">client.VectorStores.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreDeleteResponse">VectorStoreDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector_stores/{vector_store_id}/search">client.VectorStores.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreSearchParams">VectorStoreSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreSearchResponse">VectorStoreSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Files

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFile">VectorStoreFile</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileDeleteResponse">VectorStoreFileDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileContentResponse">VectorStoreFileContentResponse</a>

Methods:

- <code title="post /v1/vector_stores/{vector_store_id}/files">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileNewParams">VectorStoreFileNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/files/{file_id}">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileGetParams">VectorStoreFileGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector_stores/{vector_store_id}/files/{file_id}">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileUpdateParams">VectorStoreFileUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFile">VectorStoreFile</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/files">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileListParams">VectorStoreFileListParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFile">VectorStoreFile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/vector_stores/{vector_store_id}/files/{file_id}">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileDeleteParams">VectorStoreFileDeleteParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileDeleteResponse">VectorStoreFileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/files/{file_id}/content">client.VectorStores.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileContentParams">VectorStoreFileContentParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileContentResponse">VectorStoreFileContentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## FileBatches

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListVectorStoreFilesInBatchResponse">ListVectorStoreFilesInBatchResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatches">VectorStoreFileBatches</a>

Methods:

- <code title="post /v1/vector_stores/{vector_store_id}/file_batches">client.VectorStores.FileBatches.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, vectorStoreID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatchNewParams">VectorStoreFileBatchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatches">VectorStoreFileBatches</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/file_batches/{batch_id}">client.VectorStores.FileBatches.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatchGetParams">VectorStoreFileBatchGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatches">VectorStoreFileBatches</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/vector_stores/{vector_store_id}/file_batches/{batch_id}/cancel">client.VectorStores.FileBatches.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatchService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatchCancelParams">VectorStoreFileBatchCancelParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatches">VectorStoreFileBatches</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/vector_stores/{vector_store_id}/file_batches/{batch_id}/files">client.VectorStores.FileBatches.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatchService.ListFiles">ListFiles</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFileBatchListFilesParams">VectorStoreFileBatchListFilesParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VectorStoreFile">VectorStoreFile</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListModelsResponse">ListModelsResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ModelGetResponse">ModelGetResponse</a>

Methods:

- <code title="get /v1/models/{model_id}">client.Models.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ModelGetResponse">ModelGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListModelsResponse">ListModelsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OpenAI

Methods:

- <code title="get /v1/models">client.Models.OpenAI.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ModelOpenAIService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListModelsResponse">ListModelsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Providers

Methods:

- <code title="get /v1/providers/{provider_id}">client.Providers.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ProviderService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ProviderInfo">ProviderInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/providers">client.Providers.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ProviderService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ProviderInfo">ProviderInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Routes

Methods:

- <code title="get /v1/inspect/routes">client.Routes.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#RouteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#RouteListParams">RouteListParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#RouteInfo">RouteInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Moderations

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CreateResponse">CreateResponse</a>

Methods:

- <code title="post /v1/moderations">client.Moderations.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ModerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ModerationNewParams">ModerationNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#CreateResponse">CreateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Safety

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#RunShieldResponse">RunShieldResponse</a>

Methods:

- <code title="post /v1/safety/run-shield">client.Safety.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#SafetyService.RunShield">RunShield</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#SafetyRunShieldParams">SafetyRunShieldParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#RunShieldResponse">RunShieldResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Shields

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListShieldsResponse">ListShieldsResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Shield">Shield</a>

Methods:

- <code title="get /v1/shields/{identifier}">client.Shields.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ShieldService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/shields">client.Shields.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ShieldService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/shields/{identifier}">client.Shields.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ShieldService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /v1/shields">client.Shields.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ShieldService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ShieldRegisterParams">ShieldRegisterParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#Shield">Shield</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#DeleteFileResponse">DeleteFileResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#File">File</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ListFilesResponse">ListFilesResponse</a>

Methods:

- <code title="post /v1/files">client.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#FileNewParams">FileNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#File">File</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files">client.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#FileListParams">FileListParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#File">File</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#DeleteFileResponse">DeleteFileResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/files/{file_id}/content">client.Files.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#FileService.Content">Content</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batches

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchNewResponse">BatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchGetResponse">BatchGetResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchListResponse">BatchListResponse</a>
- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchCancelResponse">BatchCancelResponse</a>

Methods:

- <code title="post /v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchNewParams">BatchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchNewResponse">BatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/batches/{batch_id}">client.Batches.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchGetResponse">BatchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchListParams">BatchListParams</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go/packages/pagination#OpenAICursorPage">OpenAICursorPage</a>[<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchListResponse">BatchListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/batches/{batch_id}/cancel">client.Batches.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#BatchCancelResponse">BatchCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Alpha

## Admin

Methods:

- <code title="get /v1alpha/admin/health">client.Alpha.Admin.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaAdminService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#HealthInfo">HealthInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1alpha/admin/providers/{provider_id}">client.Alpha.Admin.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaAdminService.InspectProvider">InspectProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, providerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ProviderInfo">ProviderInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1alpha/admin/providers">client.Alpha.Admin.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaAdminService.ListProviders">ListProviders</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#ProviderInfo">ProviderInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1alpha/admin/inspect/routes">client.Alpha.Admin.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaAdminService.ListRoutes">ListRoutes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaAdminListRoutesParams">AlphaAdminListRoutesParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#RouteInfo">RouteInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1alpha/admin/version">client.Alpha.Admin.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaAdminService.Version">Version</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#VersionInfo">VersionInfo</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Inference

Response Types:

- <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaInferenceRerankResponse">AlphaInferenceRerankResponse</a>

Methods:

- <code title="post /v1alpha/inference/rerank">client.Alpha.Inference.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaInferenceService.Rerank">Rerank</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaInferenceRerankParams">AlphaInferenceRerankParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go">ogxclient</a>.<a href="https://pkg.go.dev/github.com/ogx-ai/ogx-client-go#AlphaInferenceRerankResponse">AlphaInferenceRerankResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
