package calculator

import "errors"

var (
	NoDivisionPorCero = errors.New("No se puede didivir por cero")
)

func dividir(a float64, b float64) (float64, error) {
	if b == 0 {
		return 0, NoDivisionPorCero
	} else {
		return a / b, nil
	}
}
