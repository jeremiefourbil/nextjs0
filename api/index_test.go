package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	total := Sum(4, 4)
	assert.Equal(t, 8, total, "Sum was incorrect, got: %d, want: %d.", total, 8)
}