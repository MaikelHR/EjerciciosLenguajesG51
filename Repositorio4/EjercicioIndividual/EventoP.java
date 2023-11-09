package EjercicioIndividual;

public class EventoP extends Evento {
    private String lugar;

    public EventoP(String nombre, String fecha, String lugar) {
        super(nombre, fecha);
        this.lugar = lugar;
    }

    public String getLugar() {
        return lugar;
    }

    public void setLugar(String lugar) {
        this.lugar = lugar;
    }

    public String toString() {
        return super.toString() + ", Lugar: " + lugar;
    }
}
