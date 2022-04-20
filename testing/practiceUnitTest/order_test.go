package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestOrdenamientoDeEnteros(t *testing.T) {
	// Dado
	var slice []int
	slice = append(slice, 50, 20, 30, 10)

	// Cuando
	ordenamientoDeSlice(&slice)

	// Entonces
	assert.IsIncreasing(t, slice)
}
