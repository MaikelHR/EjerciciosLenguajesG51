package main

import (
	"fmt"
	"strings"
)

// La función principal de este programa Go cuenta el número de caracteres, palabras y líneas de un
// texto determinado.
func main() {
	texto := `Este es un texto de ejemplo
		para probar el programa.
		Este es un texto de ejemplo
		para probar el programa.`

	numCaracteres := contarCaracteres(texto)
	numPalabras := contarPalabras(texto)
	numLineas := contarLineas(texto)

	fmt.Printf("Número de caracteres: %d\n", numCaracteres)
	fmt.Printf("Número de palabras: %d\n", numPalabras)
	fmt.Printf("Número de líneas: %d\n", numLineas)
}

func contarCaracteres(texto string) int {
	return len(texto)
}

func contarPalabras(texto string) int {
	palabras := strings.Fields(texto)
	return len(palabras)
}

func contarLineas(texto string) int {
	lineas := strings.Split(texto, "\n")
	return len(lineas)
}
