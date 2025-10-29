// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/option"
)

// AlphaPostTrainingJobService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlphaPostTrainingJobService] method instead.
type AlphaPostTrainingJobService struct {
	Options []option.RequestOption
}

// NewAlphaPostTrainingJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAlphaPostTrainingJobService(opts ...option.RequestOption) (r AlphaPostTrainingJobService) {
	r = AlphaPostTrainingJobService{}
	r.Options = opts
	return
}
