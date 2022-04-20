package calculator

func ordenamientoDeSlice(array *[]int) {
	var valorArray = *array
	var n = len(valorArray)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if valorArray[j-1] > valorArray[j] {
				valorArray[j-1], valorArray[j] = valorArray[j], valorArray[j-1]
			}
			j = j - 1
		}
	}
	*array = valorArray
}
