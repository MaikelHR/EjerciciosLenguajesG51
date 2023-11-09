package Ejercicio3;

public class EventoEmpresa extends Evento {
    private String tipoAsistencia;

    public EventoEmpresa(String nombre, String fecha, String tipo) {
        super(nombre, fecha);
        this.tipoAsistencia = tipo;
    }

    public String getTipoAsistencia() {
        return tipoAsistencia;
    }

    public void setTipoAsistencia(String tipoAsistencia) {
        this.tipoAsistencia = tipoAsistencia;
    }

    @Override
    public String toString() {
        return super.toString() + ", Tipo de asistencia del evento: " + tipoAsistencia;
    }
}
