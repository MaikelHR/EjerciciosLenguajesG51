package Ejercicio3;

public class EventoF implements AgendaF {
    @Override
    public Contacto crearContacto(String nombre, String apellido, String telefono, int edad, String detalles) {
        return null;
    }

    @Override
    public Evento crearEvento(String nombre, String fecha, String detalles) {
        if (detalles.equals("Personal")) {
            return new EventoP(nombre, fecha, "Detalles personales");
        } else if (detalles.equals("Empresarial")) {
            return new EventoEmpresa(nombre, fecha, "Detalles empresariales");
        }
        return null;
    }
}
