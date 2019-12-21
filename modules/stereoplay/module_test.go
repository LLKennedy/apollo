package stereoplay

import (
	"testing"

	"github.com/llkennedy/apollo"
	"github.com/stretchr/testify/assert"
)

func TestMeetsInterface(t *testing.T) {
	m := &Module{}
	var i *apollo.Module
	assert.Implements(t, i, m)
}
