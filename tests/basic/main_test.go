package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	var (
		input  = 1
		output = 2
	)

	actual := AddOne(1)
	if actual != output {
		t.Errorf("AddOne(%d), output %d, actual = %d", input, output, actual)
	}
}

func TestSomething(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")

	assert.NotEqual(t, 123, 456, "they should not be equal")
}
