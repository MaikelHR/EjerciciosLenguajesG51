package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type any interface{}

func main() {
	// `rand.Seed(time.Now().UnixNano())` está configurando la semilla para el generador de números
	// aleatorios en Go. Al usar `time.Now().UnixNano()`, se garantiza que la semilla sea diferente cada
	// vez que se ejecuta el programa, lo que da como resultado números aleatorios diferentes cada vez.
	rand.Seed(time.Now().UnixNano())

	// La línea `listaMatrices := generarMatriz(3, 8)` está llamando a la función `generarMatriz` con los
	// argumentos `3` y `8`. Esta función genera una lista de matrices con dimensiones aleatorias y
	// valores aleatorios. El `3` representa el número de matrices a generar y el `8` representa la
	// dimensión máxima (tanto filas como columnas) de cada matriz. Las matrices generadas luego se
	// almacenan en la variable `listaMatrices`.
	listaMatrices := generarMatriz(3, 8)

	// `verticalFunc` es una función que toma cuatro parámetros: `x`, `y`, `n` y `matriz`. Devuelve una
	// porción de valores de "runa".
	verticalFunc := func(x, y, n int, matriz [][]rune) []rune {
		if x >= len(matriz) || y >= len(matriz[0]) {
			return nil
		}

		resultado := make([]rune, 0)
		for i := x; i < len(matriz) && i-x < n; i++ {
			resultado = append(resultado, matriz[i][y])
		}
		return resultado
	}

	// `horizontalFunc` es una función que toma cuatro parámetros: `x`, `y`, `n` y `matriz`. Devuelve una
	// porción de valores de "runa".
	horizontalFunc := func(x, y, n int, matriz [][]rune) []rune {
		if x >= len(matriz) || y >= len(matriz[0]) {
			return nil
		}

		resultado := make([]rune, 0)
		for j := y; j < len(matriz[0]) && j-y < n; j++ {
			resultado = append(resultado, matriz[x][j])
		}
		return resultado
	}

	// La función `diagonalFunc` se define como un cierre que toma cuatro parámetros: `x`, `y`, `n` y
	// `matriz`. Devuelve una porción de valores de "runa".
	diagonalFunc := func(x, y, n int, matriz [][]rune) []rune {
		if x >= len(matriz) || y >= len(matriz[0]) {
			return nil
		}

		resultado := make([]rune, 0)
		for i, j := x, y; i < len(matriz) && j < len(matriz[0]) && i-x < n; i, j = i+1, j+1 {
			resultado = append(resultado, matriz[i][j])
		}
		return resultado
	}

	// La línea `funciones := []func(int, int, int, [][]rune) []rune{verticalFunc, horizontalFunc,
	// diagonalFunc}` está creando una porción de funciones. Cada función en el segmento toma cuatro
	// parámetros: `int`, `int`, `int` y `[][]rune`, y devuelve un segmento de `rune`. Las funciones
	// incluidas en el segmento son "verticalFunc", "horizontalFunc" y "diagonalFunc".
	funciones := []func(int, int, int, [][]rune) []rune{verticalFunc, horizontalFunc, diagonalFunc}
	funcionNombres := []string{"Vertical", "Horizontal", "Diagonal"}

	for i, function := range funciones {
		fmt.Printf("Function: %s\n", funcionNombres[i])
		fmt.Println("Results:")

		for j, matriz := range listaMatrices {
			fmt.Printf("  Matrix %d:\n", j+1)
			resultado := map3(matriz, function)
			for _, elementos := range resultado {
				// El código `para _, elemento := elementos de rango {
				// fmt.Print(cadena(elemento), " ")
				// }` itera sobre cada elemento en el segmento `elementos` y lo imprime como una cadena seguida de
				// un espacio. El `_` se utiliza como marcador de posición para el valor del índice, que no es
				// necesario en este caso. La `cadena(elemento)` convierte el valor de la `runa` en una cadena
				// antes de imprimirla.
				for _, elemento := range elementos {
					fmt.Print(string(elemento), " ")
				}
				fmt.Println()
			}
		}
		fmt.Println()
	}
}

// La función `generarMatriz` genera un número específico de matrices con dimensiones aleatorias y las
// llena con caracteres o números aleatorios.
// 
// @param cont El parámetro "cont" representa el número de matrices a generar. Determina cuántas
// matrices se crearán y devolverán.
// @param maxDi El parámetro `maxDi` representa la dimensión máxima de las matrices que se generarán.
// Las dimensiones de las matrices se elegirán aleatoriamente entre 3 y `maxDi` (inclusive).
// @return una porción de matrices de runas 3D.
func generarMatriz(cont, maxDi int) [][][]rune {
	matrices := make([][][]rune, cont)
	for i := 0; i < cont; i++ {
		filas := rand.Intn(maxDi-2) + 3 // Rango de 3 a dimensión máxima 8x8
		columnas := rand.Intn(maxDi-2) + 3 // Rango de 3 a dimensión máxima 8x8
		matriz := make([][]rune, filas)
		for j := 0; j < filas; j++ {
			matriz[j] = make([]rune, columnas)
			for k := 0; k < columnas; k++ {
				if rand.Intn(2) == 0 { // Probabilidad 50% de ser carácter o número
					matriz[j][k] = rune(rand.Intn(26) + 'a') // Carácter ASCII minúsculo
				} else {
					matriz[j][k] = rune(rand.Intn(100)) // Número entre 0 y 99
				}
			}
		}
		matrices[i] = matriz
	}
	return matrices
}

// La función "map3" toma una lista de matrices y una función, y aplica la función a cada matriz de la
// lista, devolviendo una nueva lista de matrices.
// 
// @param lista El parámetro "lista" es de tipo "interfaz{}", lo que significa que puede aceptar
// cualquier tipo de valor. En este caso, se espera que sea una porción de tipo [][]runa.
// @param f El parámetro `f` es una función que toma cuatro argumentos: `int`, `int`, `int` y
// `[][]rune`. Devuelve una porción de "runa".
// @return La función `map3` devuelve una porción de rodajas de runas (`[][]runa`).
func map3(lista interface{}, f func(int, int, int, [][]rune) []rune) [][]rune {
	resultado := make([][]rune, 0)
	switch reflect.TypeOf(lista).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(lista)
		for i := 0; i < s.Len(); i++ {
			matriz := s.Index(i).Interface().([][]rune)
			resultado = append(resultado, f(1, 2, 3, matriz))
		}
	}
	return resultado
}
