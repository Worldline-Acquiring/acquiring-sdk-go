// This file was automatically generated.

package domain

// DisputeEntry represents class DisputeEntry.
type DisputeEntry struct {
	Documents               *[]DisputeDocument `json:"documents,omitempty"`
	Elaboration             *string            `json:"elaboration,omitempty"`
	EntryCategory           *string            `json:"entryCategory,omitempty"`
	EntryDateTime           *string            `json:"entryDateTime,omitempty"`
	EntryID                 *string            `json:"entryId,omitempty"`
	EntryType               *string            `json:"entryType,omitempty"`
	EntryTypeDescription    *string            `json:"entryTypeDescription,omitempty"`
	MessageText             *string            `json:"messageText,omitempty"`
	Questionnaire           *string            `json:"questionnaire,omitempty"`
	ResponseDueDate         *string            `json:"responseDueDate,omitempty"`
	SchemeReason            *string            `json:"schemeReason,omitempty"`
	SchemeReasonDescription *string            `json:"schemeReasonDescription,omitempty"`
	SettlementAmount        *AmountData        `json:"settlementAmount,omitempty"`
	TransactionAmount       *AmountData        `json:"transactionAmount,omitempty"`
	UserID                  *string            `json:"userId,omitempty"`
}

// NewDisputeEntry constructs a new DisputeEntry instance.
func NewDisputeEntry() *DisputeEntry {
	return &DisputeEntry{}
}
