package go_generics

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestLength(t *testing.T) {
	var result string = Length[string]("taufiqkba")
	assert.Equal(t, "taufiqkba", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}
