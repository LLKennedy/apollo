package sine

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/llkennedy/apollo"
)

const defaultBalance = 50

// Module generates sine waves
type Module struct {
	name       string
	id         uuid.UUID
	leftData   chan byte
	rightData  chan byte
	sampleRate int
	bitDepth   int
	balance    int
	volume     int
}

// New creates a new Sine Wave Module
func New(sampleRate, bitDepth, balance, volume int) *Module {
	m := &Module{
		name:       "Sine Wave",
		id:         uuid.New(),
		leftData:   make(chan byte),
		rightData:  make(chan byte),
		sampleRate: sampleRate,
		bitDepth:   bitDepth,
		balance:    balance,
		volume:     volume,
	}
	if m.balance == 0 {
		m.balance = defaultBalance
	}
	if m.volume < 0 || m.volume > 100 {
		m.volume = 100
	}
	return m
}

// Name is a human-readable descriptor for the module.
func (m *Module) Name() string {
	if m == nil {
		return ""
	}
	return m.name
}

// ID is a randomly-generated ID for the module in case of duplicate names.
func (m *Module) ID() uuid.UUID {
	if m == nil {
		var nilID uuid.UUID
		return nilID
	}
	return m.id
}

// InputStreams returns all inputs - this list must be consistently ordered and of consistent length for a given module.
func (m *Module) InputStreams() []apollo.WriteStream {
	return []apollo.WriteStream{leftInput{source: m}, rightInput{source: m}}
}

// OutputStreams returns all outputs - this list must be consistently ordered and of consisten length for a given module.
func (m *Module) OutputStreams() []apollo.ReadStream {
	return []apollo.ReadStream{}
}

// Fiddles returns all the things you can fiddle with on the module - buttons, dials, etc.
func (m *Module) Fiddles() []apollo.Fiddle {
	return []apollo.Fiddle{balanceDial{source: m}}
}

// SampleRate is the sample rate of the audio stream in Hz
func (m *Module) SampleRate() int {
	if m == nil {
		return 0
	}
	return m.sampleRate
}

// BitDepth is the number of bits per sample
func (m *Module) BitDepth() int {
	if m == nil {
		return 0
	}
	return m.bitDepth
}

func (m *Module) Read(p []byte) (n int, err error) {
	//TODO: generate a sine wave
	return 0, fmt.Errorf("not implemented")
}
