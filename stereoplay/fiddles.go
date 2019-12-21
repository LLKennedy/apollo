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

type volumeDial struct {
	source *Module
}

func (d volumeDial) Name() string {
	return "volume (%)"
}

func (d volumeDial) Range() ([]int, bool) {
	return []int{0, 100}, false
}

func (d volumeDial) Current() int {
	return d.source.volume
}

func (d volumeDial) Default() int {
	return 100
}

func (d volumeDial) Set(f int) error {
	if f < 0 || f > 100 {
		return fmt.Errorf("out of range")
	}
	d.source.volume = f
	return nil
}

func (d volumeDial) ReadOnly() bool {
	return false
}
