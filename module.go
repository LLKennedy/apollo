package apollo

import (
	"github.com/google/uuid"
)

// Module performs arbitrary transformations on numerous inputs and outputs.
type Module interface {
	// Name is a human-readable descriptor for the module.
	Name() string
	// ID is a randomly-generated ID for the module in case of duplicate names.
	ID() uuid.UUID
	// Inputs returns all inputs - this list must be consistently ordered and of consistent length for a given module.
	InputStreams() []WriteStream
	// Outputs returns all outputs - this list must be consistently ordered and of consisten length for a given module.
	OutputStreams() []ReadStream
	// Fiddles returns all the things you can fiddle with on the module - buttons, dials, etc.
	Fiddles() []Fiddle
}
