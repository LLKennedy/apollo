package apollo

import (
	"io"

	"github.com/google/uuid"
)

// Module performs arbitrary transformations on numerous inputs and outputs.
type Module interface {
	// Name is a human-readable descriptor for the module.
	Name() string
	// ID is a randomly-generated ID for the module in case of duplicate names.
	ID() uuid.UUID
	// Inputs returns all inputs - this list must be consistently ordered and of consistent length for a given module.
	Inputs() []WriteStream
	// Outputs returns all outputs - this list must be consistently ordered and of consisten length for a given module.
	Outputs() []ReadStream
}

// Stream is an audio stream
type Stream interface {
	// Name is a human-readable descriptor for the stream.
	Name() string
	// SampleRate is the sample rate of the audio stream in Hz
	SampleRate() int
	// BitDepth is the number of bits per sample
	BitDepth() int
}

// ReadStream is a read-only stream
type ReadStream interface {
	Stream
	io.Reader
}

// WriteStream is a read-only stream
type WriteStream interface {
	Stream
	io.Writer
}
