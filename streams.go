package apollo

import "io"

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
