package sine

type output struct {
	source *Module
}

// Name is a human-readable descriptor for the stream.
func (o output) Name() string {
	return "Output"
}

// SampleRate is the sample rate of the audio stream in Hz
func (o output) SampleRate() int {
	return o.source.SampleRate()
}

// BitDepth is the number of bits per sample
func (o output) BitDepth() int {
	return o.source.BitDepth()
}

func (o output) Data() <-chan byte {
	return o.source.data
}