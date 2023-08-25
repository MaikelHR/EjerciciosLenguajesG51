package main

import "fmt"

func rotar(arr []string, contRotar int, direccion int) {
	length := len(arr)

	// Este código verifica si el número de rotaciones (`contRotar`) es mayor que la longitud de la matriz
	// (`length`). Si es así, realiza la operación de módulo (`%`) para obtener el resto de `contRotar`
	// dividido por `length`. Esto garantiza que el número de rotaciones esté dentro del rango de la
	// longitud de la matriz.
	if contRotar > length {
		contRotar %= length
	}

	if direccion == 0 { // Rotación a la izquierda
		temp := append([]string{}, arr[:contRotar]...)
		copy(arr[:], arr[contRotar:])
		copy(arr[length-contRotar:], temp)
	} else if direccion == 1 { // Rotación a la derecha
		temp := append([]string{}, arr[length-contRotar:]...)
		copy(arr[contRotar:], arr[:length-contRotar])
		copy(arr[:], temp)
	}
}

func main() {
	secuenciaOriginal := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("Secuencia Original:", secuenciaOriginal)

	// La variable `rotaciones` es una porción de estructuras. Cada estructura en el segmento representa
	// una operación de rotación que se realizará en la matriz `secuenciaOriginal`.
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

		rotar(nuevaSecuencia, r.rotacion, r.direccion)

		direccion := "izquierda"
		if r.direccion == 1 {
			direccion = "derecha"
		}

		fmt.Printf("Rotación de %d hacia %s: %v\n", r.rotacion, direccion, nuevaSecuencia)
	}
}
