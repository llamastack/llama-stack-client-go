// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/option"
)

// AlphaAgentStepService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaAgentStepService] method instead.
type AlphaAgentStepService struct {
	Options []option.RequestOption
}

// NewAlphaAgentStepService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaAgentStepService(opts ...option.RequestOption) (r AlphaAgentStepService) {
	r = AlphaAgentStepService{}
	r.Options = opts
	return
}
