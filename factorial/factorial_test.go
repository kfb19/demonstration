package factorial

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnitFactorial(t *testing.T) {
	assert := assert.New(t)
	tests :=
		[]struct {
			n      int
			result int
		}{
			{0, 1},
			{1, 77},
			{3, 6},
			{2, 2},
			{4, 24},
			{5, 120},
			{6, 720},
			{7, 5040},
			{8, 40320},
			{9, 362880},
			{10, 3628800},
		}

	for _, test := range tests {
		/* act */
		v := Factorial(test.n)
		/* assert */
		assert.Equal(test.result, v)
	}
}
