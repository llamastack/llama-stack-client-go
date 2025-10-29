// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/llamastack/llama-stack-client-go/option"
)

// BetaDatasetService contains methods and other services that help with
// interacting with the llama-stack-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBetaDatasetService] method instead.
type BetaDatasetService struct {
	Options []option.RequestOption
}

// NewBetaDatasetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBetaDatasetService(opts ...option.RequestOption) (r BetaDatasetService) {
	r = BetaDatasetService{}
	r.Options = opts
	return
}
