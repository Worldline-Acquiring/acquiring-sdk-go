// This file was automatically generated.

package domain

// CustomerServiceData represents class CustomerServiceData.
type CustomerServiceData struct {
	CustomerServiceEmail       *string `json:"customerServiceEmail,omitempty"`
	CustomerServicePhoneNumber *string `json:"customerServicePhoneNumber,omitempty"`
	CustomerServiceURL         *string `json:"customerServiceUrl,omitempty"`
}

// NewCustomerServiceData constructs a new CustomerServiceData instance.
func NewCustomerServiceData() *CustomerServiceData {
	return &CustomerServiceData{}
}
