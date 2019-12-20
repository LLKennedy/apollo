package apollo

// Fiddle is a thing you can fiddle with on the module
type Fiddle interface {
	Name() string
	// Values are the only valid values if discreet is true, or the upper and lower bounds as well as recommended values if discreet is false
	Range() (values []int, discrete bool)
	Current() int
	Default() int
	Set(int) error
	ReadOnly() bool
}
