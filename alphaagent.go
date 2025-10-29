// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/option"
)

// AlphaAgentService contains methods and other services that help with interacting
// with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaAgentService] method instead.
type AlphaAgentService struct {
	Options []option.RequestOption
	Session AlphaAgentSessionService
	Steps   AlphaAgentStepService
	Turn    AlphaAgentTurnService
}

// NewAlphaAgentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlphaAgentService(opts ...option.RequestOption) (r AlphaAgentService) {
	r = AlphaAgentService{}
	r.Options = opts
	r.Session = NewAlphaAgentSessionService(opts...)
	r.Steps = NewAlphaAgentStepService(opts...)
	r.Turn = NewAlphaAgentTurnService(opts...)
	return
}
