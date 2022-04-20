package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func TestDividir(t *testing.T) {
	// Dado
	var dividendo float64
	var divisor float64
	dividendo = 10
	divisor = 2
	resultadoEsperado := 5.0

	// Cuando
	resultado, err := dividir(dividendo, divisor)

	// Entonces
	assert.Equal(t, resultado, resultadoEsperado)
	assert.Nil(t, err)
}

func TestDividirPorCero(t *testing.T) {
	// Dado
	var dividendo float64
	var divisor float64
	dividendo = 10
	divisor = 0
	resultadoEsperado := 0.0

	// Cuando
	resultado, err := dividir(dividendo, divisor)

	// Entonces
	assert.Equal(t, resultado, resultadoEsperado)
	assert.NotNil(t, err)
}

func TestDevolverNoDivisionPorCero(t *testing.T) {
	// Dado
	var dividendo float64
	var divisor float64
	dividendo = 10
	divisor = 0

	// Cuando
	_, err := dividir(dividendo, divisor)

	// Entonces
	assert.ErrorIs(t, err, NoDivisionPorCero)
}
