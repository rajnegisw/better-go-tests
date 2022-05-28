package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack_Push(t *testing.T) {

	s := Stack{}
	s.Push("Test")

	assert.Equal(t, 1, s.size, "length should be equal")
	assert.Equal(t, "Test", s.head.payload, "payload should be equal")
}
