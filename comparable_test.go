package go_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}
func TestComparable(t *testing.T) {
	assert.True(t, IsSame[string]("taufiq", "taufiq"))
	assert.True(t, IsSame[int](100, 100))
}
