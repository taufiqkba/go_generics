package go_generics

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"testing"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, 100.0, ExperimentalMin(100.0, 200.0))
}

// Experimental Maps
func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Taufiq",
	}
	second := map[string]string{
		"Name": "Taufiq",
	}
	assert.True(t, maps.Equal(first, second))
}

// Experimental Maps
func TestExperimentalSlice(t *testing.T) {
	first := []string{"Taufiq"}
	second := []string{"Taufiq"}
	assert.True(t, slices.Equal(first, second))
}
