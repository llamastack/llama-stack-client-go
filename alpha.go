// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/option"
)

// AlphaService contains methods and other services that help with interacting with
// the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaService] method instead.
type AlphaService struct {
	Options      []option.RequestOption
	Inference    AlphaInferenceService
	PostTraining AlphaPostTrainingService
	Eval         AlphaEvalService
	Agents       AlphaAgentService
}

// NewAlphaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAlphaService(opts ...option.RequestOption) (r AlphaService) {
	r = AlphaService{}
	r.Options = opts
	r.Inference = NewAlphaInferenceService(opts...)
	r.PostTraining = NewAlphaPostTrainingService(opts...)
	r.Eval = NewAlphaEvalService(opts...)
	r.Agents = NewAlphaAgentService(opts...)
	return
}
