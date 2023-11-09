package Ejercicio3;

public class ContactoP extends Contacto {
    private String direccion;

    public ContactoP(String nombre, String apellido, String telefono, int edad, String direccion) {
        super(nombre, apellido, telefono, edad);
        this.direccion = direccion;
    }

    public String getDireccion() {
        return direccion;
    }

    public void setDireccion(String direccion) {
        this.direccion = direccion;
    }

    @Override
    public String toString() {
        return super.toString() + ", Dirección actual: " + direccion;
    }
}