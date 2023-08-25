package main

import "fmt"

type calzado struct {
	Modelo string
	Precio float64
	Talla  int
	Stock  int
}

func venderZapato(z *calzado) bool {
	// El código verifica si el stock de un zapato determinado es mayor que 0. Si es así, disminuye el
	// stock en 1 y devuelve verdadero, lo que indica que el zapato se vendió exitosamente. Si el stock es
	// 0 o menos, devuelve falso, lo que indica que el zapato está agotado y no se puede vender.
	if z.Stock > 0 {
		z.Stock--
		return true
	}
	return false
}

func main() {
	inventario := []calzado{
		{"Nike", 50000, 42, 5},
		{"Adidas", 60000, 39, 10},
		{"Puma", 45000, 37, 8},
		// Se puede agregar más zapatos aquí
	}

	fmt.Println("Inventario de la tienda:")
	// El código utiliza un bucle `for` para iterar sobre cada elemento en el segmento `inventario`. El
	// bucle asigna cada elemento a la variable `zapato`, pero como no necesitamos el valor del índice,
	// usamos el guión bajo `_` para ignorarlo.
	for _, zapato := range inventario {
		fmt.Printf("Modelo: %s, Talla: %d, Precio: %.2f, Stock: %d\n", zapato.Modelo, zapato.Talla, zapato.Precio, zapato.Stock)
	}

	fmt.Println("\nRealizando ventas:")
	fmt.Println("Vendiendo 2 pares de Nike talla 42")
	// Este código simula el proceso de venta de 2 pares de zapatos Nike con talla 42.
	marca := &inventario[0]
	for i := 0; i < 2; i++ {
		if venderZapato(marca) {
			fmt.Println("Venta exitosa")
		} else {
			fmt.Println("No hay stock disponible")
		}
	}

	fmt.Println("\nInventario después de las ventas:")
	// El código utiliza un bucle `for` para iterar sobre cada elemento en el segmento `inventario`. El
	// bucle asigna cada elemento a la variable `zapato`, pero como no necesitamos el valor del índice,
	// usamos el guión bajo `_` para ignorarlo.
	for _, zapato := range inventario {
		fmt.Printf("Modelo: %s, Talla: %d, Precio: %.2f, Stock: %d\n", zapato.Modelo, zapato.Talla, zapato.Precio, zapato.Stock)
	}
}
