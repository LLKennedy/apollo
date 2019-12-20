package sine

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/llkennedy/apollo"
)

const defaultFrequency = 1000

// Generator generates sine waves
type Generator struct {
	name       string
	id         uuid.UUID
	sampleRate int
	bitDepth   int
	frequency  int
	magnitude  int
}

// New creates a new Sine Wave generator
func New(sampleRate, bitDepth, frequency, magnitude int) *Generator {
	g := &Generator{
		name:       "Sine Wave",
		id:         uuid.New(),
		sampleRate: sampleRate,
		bitDepth:   bitDepth,
		frequency:  frequency,
		magnitude:  magnitude,
	}
	if g.frequency == 0 {
		g.frequency = defaultFrequency
	}
	if g.magnitude < 0 || g.magnitude > 100 {
		g.magnitude = 100
	}
	return g
}

// Name is a human-readable descriptor for the module.
func (g *Generator) Name() string {
	if g == nil {
		return ""
	}
	return g.name
}

// ID is a randomly-generated ID for the module in case of duplicate names.
func (g *Generator) ID() uuid.UUID {
	if g == nil {
		var nilID uuid.UUID
		return nilID
	}
	return g.id
}

// InputStreams returns all inputs - this list must be consistently ordered and of consistent length for a given module.
func (g *Generator) InputStreams() []apollo.WriteStream {
	return []apollo.WriteStream{}
}

// OutputStreams returns all outputs - this list must be consistently ordered and of consisten length for a given module.
func (g *Generator) OutputStreams() []apollo.ReadStream {
	return []apollo.ReadStream{output{source: g}}
}

// Fiddles returns all the things you can fiddle with on the module - buttons, dials, etc.
func (g *Generator) Fiddles() []apollo.Fiddle {
	return []apollo.Fiddle{frequencyDial{source: g}}
}

// SampleRate is the sample rate of the audio stream in Hz
func (g *Generator) SampleRate() int {
	if g == nil {
		return 0
	}
	return g.sampleRate
}

// BitDepth is the number of bits per sample
func (g *Generator) BitDepth() int {
	if g == nil {
		return 0
	}
	return g.bitDepth
}

func (g *Generator) Read(p []byte) (n int, err error) {
	//TODO: generate a sine wave
	return 0, fmt.Errorf("not implemented")
}
