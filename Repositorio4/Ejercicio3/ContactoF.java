package Ejercicio3;

public class ContactoF implements AgendaF{
    public Contacto crearContacto(String nombre, String apellido, String telefono, int edad, String detalles) {
        if (detalles.equals("Personal")) {
            return new ContactoP(nombre, apellido, telefono, edad, "Detalles personales");
        } else if (detalles.equals("Empresarial")) {
            return new ContactoEmpresa(nombre, apellido, telefono, edad, "Detalles empresariales");
        }
        return null;
    }

    
    public Evento crearEvento(String nombre, String fecha, String detalles) {
        return null;
    }
}
