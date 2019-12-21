package virtualcable

import "github.com/llkennedy/apollo"

// Patch is a virtual patch cable
type Patch struct {
	closer chan bool
}

// New creates a new virtual patch cable
func New(from apollo.ReadStream, to apollo.WriteStream) *Patch {
	p := new(Patch)
	go p.connect(from.Data(), to.Data())
	return p
}

func (p *Patch) connect(from <-chan byte, to chan<- byte) {
	defer recover() // Ignore all panics
	for {
		select {
		case to <- <-from:
			continue
		case shouldClose := <-p.closer:
			if shouldClose {
				return
			}
		}

	}
}

// Close ends
func (p *Patch) Close() {
	if p != nil && p.closer != nil {
		p.closer <- true
	}
}
