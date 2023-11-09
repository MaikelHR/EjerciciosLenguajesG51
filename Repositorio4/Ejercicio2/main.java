package Ejercicio2;

import java.util.ArrayList;
import java.util.List;

public class main {
    public static void main(String[] args) {
        List<Socio> socios = new ArrayList<>();
        socios.add(new Socio(1, "Juan", "Calle A"));
        socios.add(new Socio(2, "María", "Calle B"));

        Biblioteca biblioteca = new Biblioteca(socios);

        List<Socio> sociosConMasDeTresLibros = biblioteca.obtenerSociosConMasDeTresLibrosPrestados();

        for (Socio socio : sociosConMasDeTresLibros) {
            System.out.println("Socio con más de 3 libros prestados: " + socio.getNumeroSocio());
        }
    }
}
