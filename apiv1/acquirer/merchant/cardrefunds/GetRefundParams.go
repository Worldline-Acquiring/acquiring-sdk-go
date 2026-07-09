// This file was automatically generated.

package cardrefunds

import (
	"strconv"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
)

// GetRefundParams represents query parameters for Retrieve card refund.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Card-Refunds/operation/getRefund.
type GetRefundParams struct {
	ReturnOperations *bool
}

// ToRequestParameters converts the query to communication.RequestParams.
func (params *GetRefundParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.ReturnOperations != nil {
		param, _ := communication.NewRequestParam("returnOperations", strconv.FormatBool(*params.ReturnOperations))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewGetRefundParams constructs a new GetRefundParams instance.
func NewGetRefundParams() *GetRefundParams {
	return &GetRefundParams{}
}
