package acquiringsdk

// NewBool returns a pointer for a boolean value.
// This can be used to assign literals to domain or query parameter objects
func NewBool(value bool) *bool {
	return &value
}

// NewInt32 returns a pointer for an int32 value.
// This can be used to assign literals to domain or query parameter objects
func NewInt32(value int32) *int32 {
	return &value
}

// NewInt64 returns a pointer for an int64 value.
// This can be used to assign literals to domain or query parameter objects
func NewInt64(value int64) *int64 {
	return &value
}

// NewFloat64 returns a pointer for a float64 value.
// This can be used to assign literals to domain or query parameter objects
func NewFloat64(value float64) *float64 {
	return &value
}

// NewString returns a pointer for a string value.
// This can be used to assign literals to domain or query parameter objects
func NewString(value string) *string {
	return &value
}
