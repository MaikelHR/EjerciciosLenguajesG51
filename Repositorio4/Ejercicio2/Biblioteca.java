package Ejercicio2;

import java.util.List;
import java.util.stream.Collectors;

public class Biblioteca {
    private List<Socio> socios;

    public Biblioteca(List<Socio> socios) {
        this.socios = socios;
    }

    public List<Socio> obtenerSociosConMasDeTresLibrosPrestados() {
        return socios.stream()
                .filter(socio -> socio.getLibrosPrestados() > 3)
                .collect(Collectors.toList());
    }
}
