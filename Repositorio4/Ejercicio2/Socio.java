package Ejercicio2;

public class Socio {
    private int numeroSocio;
    private String nombre;
    private String direccion;
    private int librosPrestados;

    public Socio(int numeroSocio, String nombre, String direccion) {
        this.numeroSocio = numeroSocio;
        this.nombre = nombre;
        this.direccion = direccion;
        this.librosPrestados = 0;
    }

    public int getNumeroSocio() {
        return this.numeroSocio;
    }

    public void setNumeroSocio(int numeroSocio) {
        this.numeroSocio = numeroSocio;
    }

    public String getNombre() {
        return this.nombre;
    }

    public void setNombre(String nombre) {
        this.nombre = nombre;
    }

    public String getDireccion() {
        return this.direccion;
    }

    public void setDireccion(String direccion) {
        this.direccion = direccion;
    }

    public int getLibrosPrestados() {
        return this.librosPrestados;
    }

    public void setLibrosPrestados(int librosPrestados) {
        this.librosPrestados = librosPrestados;
    }

    public void prestarLibro() {
        librosPrestados++;
    }

    public void devolverLibro() {
        librosPrestados--;
    }
}
