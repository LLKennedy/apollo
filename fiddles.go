package apollo

// Fiddle is a thing you can fiddle with on the module
type Fiddle interface {
	Name() string
	Range() []int
	Current() int
	Set(int) error
	ReadOnly() bool
}
