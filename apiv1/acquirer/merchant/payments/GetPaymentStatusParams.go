// This file was automatically generated.

package payments

import (
	"strconv"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
)

// GetPaymentStatusParams represents query parameters for Retrieve payment
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Payments/operation/getPaymentStatus
type GetPaymentStatusParams struct {
	ReturnOperations *bool
}

// ToRequestParameters converts the query to communication.RequestParams
func (params *GetPaymentStatusParams) ToRequestParameters() communication.RequestParams {
	reqParams := communication.RequestParams{}

	if params.ReturnOperations != nil {
		param, _ := communication.NewRequestParam("returnOperations", strconv.FormatBool(*params.ReturnOperations))
		reqParams = append(reqParams, *param)
	}

	return reqParams
}

// NewGetPaymentStatusParams constructs a new GetPaymentStatusParams instance
func NewGetPaymentStatusParams() *GetPaymentStatusParams {
	return &GetPaymentStatusParams{}
}
