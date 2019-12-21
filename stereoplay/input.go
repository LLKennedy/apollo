package sine

type leftInput struct {
	source *Module
}

// Name is a human-readable descriptor for the stream.
func (i leftInput) Name() string {
	return "Left"
}

// SampleRate is the sample rate of the audio stream in Hz
func (i leftInput) SampleRate() int {
	return i.source.SampleRate()
}

// BitDepth is the number of bits per sample
func (i leftInput) BitDepth() int {
	return i.source.BitDepth()
}

func (i leftInput) Data() chan<- byte {
	return i.source.leftData
}

type rightInput struct {
	source *Module
}

// Name is a human-readable descriptor for the stream.
func (i rightInput) Name() string {
	return "Right"
}

// SampleRate is the sample rate of the audio stream in Hz
func (i rightInput) SampleRate() int {
	return i.source.SampleRate()
}

// BitDepth is the number of bits per sample
func (i rightInput) BitDepth() int {
	return i.source.BitDepth()
}

func (i rightInput) Data() chan<- byte {
	return i.source.rightData
}
