package main

import "fmt"

func rotate(arr []string, rotationCount int, direction int) {
	length := len(arr)

	if rotationCount > length {
		rotationCount %= length
	}

	if direction == 0 { // Rotación a la izquierda
		temp := append([]string{}, arr[:rotationCount]...)
		copy(arr[:], arr[rotationCount:])
		copy(arr[length-rotationCount:], temp)
	} else if direction == 1 { // Rotación a la derecha
		temp := append([]string{}, arr[length-rotationCount:]...)
		copy(arr[rotationCount:], arr[:length-rotationCount])
		copy(arr[:], temp)
	}
}

func main() {
	secuenciaOriginal := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("Secuencia Original:", secuenciaOriginal)

	rotaciones := []struct {
		rotacion int
		direccion int
	}{
		{3, 0}, // Rotación a la izquierda
		{2, 1}, // Rotación a la derecha
	}

	for _, r := range rotaciones {
		nuevaSecuencia := make([]string, len(secuenciaOriginal))
		copy(nuevaSecuencia, secuenciaOriginal)

		rotate(nuevaSecuencia, r.rotacion, r.direccion)

		direccion := "izquierda"
		if r.direccion == 1 {
			direccion = "derecha"
		}

		fmt.Printf("Rotación de %d hacia %s: %v\n", r.rotacion, direccion, nuevaSecuencia)
	}
}
