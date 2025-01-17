package sub

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSub(t *testing.T) {
	result := Sub(3, 2)

	assert.Equal(t, 1, result)
}
