package communication

import "io"

// ResponseHandler is a function for handling incoming responses
type ResponseHandler func(statusCode int, headers []Header, reader io.Reader) (interface{}, error)
