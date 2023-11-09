package EjercicioIndividual;

public class ContactoEmpresa extends Contacto {
    private String email;

    public ContactoEmpresa(String nombre, String apellido, String telefono, int edad, String email) {
        super(nombre, apellido, telefono, edad);
        this.email = email;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String toString() {
        return super.toString() + ", Email: " + email;
    }
}
