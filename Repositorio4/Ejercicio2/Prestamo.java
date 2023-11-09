package Ejercicio2;

import java.time.LocalDate;

public class Prestamo {
    private int codigoLibro;
    private int numeroSocio;
    private LocalDate fecha;

    public Prestamo(int codigoLibro, int numeroSocio) {
        this.codigoLibro = codigoLibro;
        this.numeroSocio = numeroSocio;
        this.fecha = LocalDate.now(); 
    }

    public int getCodigoLibro() {
        return this.codigoLibro;
    }

    public void setCodigoLibro(int codigoLibro) {
        this.codigoLibro = codigoLibro;
    }

    public int getNumeroSocio() {
        return this.numeroSocio;
    }

    public void setNumeroSocio(int numeroSocio) {
        this.numeroSocio = numeroSocio;
    }

    public LocalDate getFecha() {
        return this.fecha;
    }

    public void setFecha(LocalDate fecha) {
        this.fecha = fecha;
    }
}
