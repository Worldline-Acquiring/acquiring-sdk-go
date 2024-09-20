// This file was automatically generated.

package domain

// CardOnFileData represents class CardOnFileData
type CardOnFileData struct {
	InitialCardOnFileData    *InitialCardOnFileData    `json:"initialCardOnFileData,omitempty"`
	IsInitialTransaction     *bool                     `json:"isInitialTransaction,omitempty"`
	SubsequentCardOnFileData *SubsequentCardOnFileData `json:"subsequentCardOnFileData,omitempty"`
}

// NewCardOnFileData constructs a new CardOnFileData instance
func NewCardOnFileData() *CardOnFileData {
	return &CardOnFileData{}
}
