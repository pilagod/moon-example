package add

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result := Add(1, 1)

	assert.Equal(t, 2, result)
}
