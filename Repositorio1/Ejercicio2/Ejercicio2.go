package main

import "fmt"

func main() {
	tamano := 5 // Cambia este valor para ajustar el tamaño del rombo
	imprimirRombo(tamano)
}

func imprimirRombo(tamano int) {
	if tamano%2 == 0 {
		tamano++ // Asegurarse de que el tamaño sea impar para que el rombo sea simétrico
	}

	mitad := tamano / 2

	// Imprimir la parte superior del rombo
	for i := 0; i < mitad; i++ {
		for j := 0; j < mitad-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Imprimir la parte inferior del rombo
	for i := mitad; i >= 0; i-- {
		for j := 0; j < mitad-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
