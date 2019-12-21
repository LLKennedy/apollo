package sine

import "fmt"

type balanceDial struct {
	source *Module
}

func (d balanceDial) Name() string {
	return "Frequency (Hz)"
}

func (d balanceDial) Range() ([]int, bool) {
	return []int{1, 100000}, false
}

func (d balanceDial) Current() int {
	return d.source.balance
}

func (d balanceDial) Default() int {
	return defaultBalance
}

func (d balanceDial) Set(f int) error {
	if f < 1 || f > 100000 {
		return fmt.Errorf("out of range")
	}
	d.source.balance = f
	return nil
}

func (d balanceDial) ReadOnly() bool {
	return false
}

type magnitudeDial struct {
	source *Module
}

func (d magnitudeDial) Name() string {
	return "Magnitude (%)"
}

func (d magnitudeDial) Range() ([]int, bool) {
	return []int{0, 100}, false
}

func (d magnitudeDial) Current() int {
	return d.source.magnitude
}

func (d magnitudeDial) Default() int {
	return 100
}

func (d magnitudeDial) Set(f int) error {
	if f < 0 || f > 100 {
		return fmt.Errorf("out of range")
	}
	d.source.magnitude = f
	return nil
}

func (d magnitudeDial) ReadOnly() bool {
	return false
}
