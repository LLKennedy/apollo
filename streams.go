package apollo

// ReadStream is a read-only stream
type ReadStream interface {
	// Name is a human-readable descriptor for the stream.
	Name() string
	// SampleRate is the sample rate of the audio stream in Hz
	SampleRate() int
	// BitDepth is the number of bits per sample
	BitDepth() int
	// Data returns the data channel
	Data() <-chan byte
}

// WriteStream is a read-only stream
type WriteStream interface {
	// Name is a human-readable descriptor for the stream.
	Name() string
	// SampleRate is the sample rate of the audio stream in Hz
	SampleRate() int
	// BitDepth is the number of bits per sample
	BitDepth() int
	// Data returns the data channel
	Data() chan<- byte
}
