package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}

type listaProductos []producto

var lProductos listaProductos

const existenciaMinima int = 10

// La función `agregarProducto` se utiliza para agregar un nuevo producto a la lista de productos. Toma
// tres parámetros: `nombre` (nombre del producto), `cantidad` (cantidad del producto) y `precio`
// (precio del producto).
func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	encontrado := false
	for i, prod := range *l {
		if prod.nombre == nombre {
			(*l)[i].cantidad += cantidad
			if prod.precio != precio {
				(*l)[i].precio = precio
			}
			encontrado = true
			break
		}
	}

	if !encontrado {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
}

// La función `agregarVariosProductos` se utiliza para agregar múltiples productos a la lista de
// productos. Toma un número variable de argumentos de `producto` y los itera, llamando a la función
// `agregarProducto` para cada producto para agregarlo a la lista.
func (l *listaProductos) agregarVariosProductos(productos ...producto) {
	for _, p := range productos {
		l.agregarProducto(p.nombre, p.cantidad, p.precio)
	}
}

// La función `buscarProducto` se utiliza para buscar un producto en la lista de productos. Toma un
// parámetro `nombre` y devuelve tres valores: el producto en sí, el índice del producto en la lista y
// un código de error.
func (l *listaProductos) buscarProducto(nombre string) (producto, int, int) {
	for i, prod := range *l {
		if prod.nombre == nombre {
			return prod, i, 0
		}
	}
	return producto{}, -1, -1
}

// La función `venderProducto` se utiliza para vender un producto de la lista de productos. Toma dos
// parámetros: `nombre` (nombre del producto) y `cantidad` (cantidad a vender).
func (l *listaProductos) venderProducto(nombre string, cantidad int) int {
	_, prodIndex, err := l.buscarProducto(nombre)
	// El código `if err != 0 { return err }` comprueba si la variable `err` no es igual a 0. Si no es
	// igual a 0, significa que ocurrió un error durante la ejecución de una función. En este caso, la
	// función es `buscarProducto` o `venderProducto`.
	if err != 0 {
		return err
	}

	// Este bloque de código verifica si la cantidad de un producto en la lista es mayor o igual a la
	// cantidad a vender. Si es así, resta la cantidad vendida de la cantidad del producto. Si la cantidad
	// del producto llega a ser cero después de la resta, llama a la función `eliminarProductoPorIndice`
	// para eliminar el producto de la lista.
	if (*l)[prodIndex].cantidad >= cantidad {
		(*l)[prodIndex].cantidad -= cantidad
		if (*l)[prodIndex].cantidad == 0 {
			l.eliminarProductoPorIndice(prodIndex)
		}
		return 0
	}
	return -2
}

// La función `eliminarProductoPorIndice` se utiliza para eliminar un producto de la lista de productos
// en un índice específico. Se necesitan dos parámetros: `index` (el índice del producto que se va a
// eliminar) y `l` (un puntero a la lista de productos).
func (l *listaProductos) eliminarProductoPorIndice(index int) {
	*l = append((*l)[:index], (*l)[index+1:]...)
}

// La función `modificarPrecio` se utiliza para modificar el precio de un producto en la lista de
// productos. Toma dos parámetros: `nombre` (nombre del producto) y `nuevoPrecio` (nuevo precio del
// producto).
func (l *listaProductos) modificarPrecio(nombre string, nuevoPrecio int) int {
	_, prodIndex, err := l.buscarProducto(nombre)
	// El código `if err != 0 { return err }` está verificando si la variable `err` no es igual a 0. Si no
	// es igual a 0, significa que ocurrió un error durante la ejecución de una función. En este caso las
	// funciones que se están verificando son `buscarProducto` y `venderProducto`. Si ocurre un error, la
	// función devuelve el valor de "err", que puede usarse para manejar el error en el código de llamada.
	if err != 0 {
		return err
	}

	(*l)[prodIndex].precio = nuevoPrecio
	return 0
}

// La función `modificarPrecioProducto` se utiliza para modificar el precio de un producto en la lista
// de productos. Toma dos parámetros: `nombre` (nombre del producto) y `nuevoPrecio` (nuevo precio del
// producto).
func (l *listaProductos) modificarPrecioProducto(nombre string, nuevoPrecio int) int {
	_, prodIndex, err := l.buscarProducto(nombre)
	if err != 0 {
		return err
	}

	// La línea `(*l)[prodIndex].precio = nuevoPrecio` está actualizando el precio de un producto en la
	// lista de productos. Es acceder al producto en la posición `prodIndex` de la lista (`*l`) y fijar su
	// `precio` al valor de `nuevoPrecio`.
	(*l)[prodIndex].precio = nuevoPrecio
	return 0
}

func (l *listaProductos) listarProductosMínimos() []producto {
	productosMinimos := []producto{}

	// Este código se repite sobre cada producto en la lista de productos (`*l`). Comprueba si la cantidad
	// (“prod.cantidad”) del producto es menor o igual al stock mínimo (“existenciaMinima”). Si es así, el
	// producto se agrega al segmento `productosMinimos` usando la función `append`. Este código se
	// utiliza para buscar y recoger todos los productos que tengan una cantidad inferior al nivel mínimo
	// de stock.
	for _, prod := range *l {
		if prod.cantidad <= existenciaMinima {
			productosMinimos = append(productosMinimos, prod)
		}
	}

	return productosMinimos
}

func (l *listaProductos) aumentarInventarioDeMinimos(listaMinimos []producto) {
	// El código se itera sobre cada producto en el segmento `listaMinimos`. Para cada producto calcula la
	// cantidad necesaria ("cantidadNecesaria") para alcanzar el nivel de stock mínimo
	// ("existenciaMinima") restando la cantidad actual del producto al nivel de stock mínimo.
	for _, prod := range listaMinimos {
		cantidadNecesaria := existenciaMinima - prod.cantidad
		if cantidadNecesaria > 0 {
			l.agregarProducto(prod.nombre, cantidadNecesaria, prod.precio)
		}
	}
}

func (l *listaProductos) ordenarPorPrecio() {
	// La función `sort.SliceStable` se utiliza para ordenar la porción `listaProductos` de estructuras
	// `producto` por el campo `precio` en orden ascendente.
	sort.SliceStable(*l, func(i, j int) bool {
		return (*l)[i].precio < (*l)[j].precio
	})
}

func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("arroz", 20, 2000)
}

func main() {
	llenarDatos()
	fmt.Println("Inventario inicial:")
	fmt.Println(lProductos)

	// Modificar precio de un producto
	lProductos.modificarPrecio("arroz", 2700)
	lProductos.modificarPrecioProducto("frijoles", 2100)

	// Agregar varios productos
	lProductos.agregarVariosProductos(
		producto{nombre: "azúcar", cantidad: 5, precio: 1000},
		producto{nombre: "sal", cantidad: 10, precio: 500},
	)

	// Vender productos
	lProductos.venderProducto("arroz", 1)
	lProductos.venderProducto("frijoles", 3)

	fmt.Println("Inventario después de modificaciones, agregados y ventas:")
	fmt.Println(lProductos)

	// Listar productos con existencia mínima
	productosMinimos := lProductos.listarProductosMínimos()
	fmt.Println("Productos con existencia mínima:")
	fmt.Println(productosMinimos)

	// Aumentar inventario de productos mínimos
	lProductos.aumentarInventarioDeMinimos(productosMinimos)
	fmt.Println("Inventario después de aumentar inventario de productos mínimos:")
	fmt.Println(lProductos)

	// Ordenar la lista de productos por precio
	lProductos.ordenarPorPrecio()
	fmt.Println("Inventario ordenado por precio:")
	fmt.Println(lProductos)
}
