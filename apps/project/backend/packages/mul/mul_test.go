package mul

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMul(t *testing.T) {
	result := Mul(2, 3)

	assert.Equal(t, 6, result)
}
