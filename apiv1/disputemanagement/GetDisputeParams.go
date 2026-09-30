// This file was automatically generated.

package disputemanagement

import (
	"strconv"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
)

// GetDisputeParams represents query parameters for Retrieve Dispute.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Dispute-Management/operation/getDispute.
type GetDisputeParams struct {
	IncludeEntries *bool
}

// ToRequestParameters converts the query to communication.RequestParams.
func (params *GetDisputeParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.IncludeEntries != nil {
		param, _ := communication.NewRequestParam("includeEntries", strconv.FormatBool(*params.IncludeEntries))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewGetDisputeParams constructs a new GetDisputeParams instance.
func NewGetDisputeParams() *GetDisputeParams {
	return &GetDisputeParams{}
}
