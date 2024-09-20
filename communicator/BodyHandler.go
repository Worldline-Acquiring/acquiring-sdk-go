package communicator

import (
	"io"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
)

// BodyHandler is a function for handling incoming body streams
type BodyHandler func(headers []communication.Header, reader io.Reader) error
