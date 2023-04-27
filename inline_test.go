package go_generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, 100.0, FindMin(100.0, 200.0))
}

func GetFirst[T []E, E any](data T) E {
	first := data[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"taufiq", "kurniawan", "bayu",
	}
	first := GetFirst[[]string, string](names)
	assert.Equal(t, "taufiq", first)
}
