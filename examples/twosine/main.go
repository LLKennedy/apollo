package main

import (
	"time"

	"github.com/llkennedy/apollo/modules/sine"
	"github.com/llkennedy/apollo/modules/stereoplay"
	"github.com/llkennedy/apollo/patches/virtualcable"
)

func main() {
	// Play two sine waves, one in each ear.
	sampleRate := 44100
	bitDepth := 2
	frequency1 := 2000
	frequency2 := 4000
	sine1 := sine.New(sampleRate, bitDepth, frequency1, 100)
	sine2 := sine.New(sampleRate, bitDepth, frequency2, 100)
	output := stereoplay.New(sampleRate, bitDepth, 50, 100)
	patch1 := virtualcable.New(sine1.OutputStreams()[0], output.InputStreams()[0])
	patch2 := virtualcable.New(sine2.OutputStreams()[0], output.InputStreams()[1])
	defer patch1.Close()
	defer patch2.Close()
	time.Sleep(10 * time.Second)
}
