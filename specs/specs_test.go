package specs

import "testing"
import "github.com/stretchr/testify/assert"

func TestReadFromString(t *testing.T) {

	specString := "1000\n750\n-0.75\n-0.16\n-0.73\n-0.145\n1000"

	specs := ReadFromString(specString)

	assert.Equal(t, 1000, specs.Width)
	assert.Equal(t, 750, specs.Height)
	assert.Equal(t, -0.75, specs.MinR)
	assert.Equal(t, -0.16, specs.MinI)
	assert.Equal(t, -0.73, specs.MaxR)
	assert.Equal(t, -0.145, specs.MaxI)
	assert.Equal(t, 1000, specs.MaximumNumberOfIterations)
}
