package EjercicioIndividual;

public class main {
    public static void main(String[] args) {
        Agenda agenda = new Agenda();

        ContactoP contactoPersonal = new ContactoP("Luna", "Miranda", "70854467",22, "San Jose, Costa Rica");
        ContactoEmpresa contactoEmpresarial = new ContactoEmpresa("Max", "Lopez", "708904467", 25, "max.lopez@example.com");

        EventoP eventoPersonal = new EventoP("Cumpleannos", "18/09/23", "Restaurante Example");
        EventoEmpresa eventoEmpresarial = new EventoEmpresa("Reunión Zoom", "22/10/23", "Obligatoria");

        agenda.agregarContacto(contactoPersonal);
        agenda.agregarContacto(contactoEmpresarial);
        agenda.agregarEvento(eventoPersonal);
        agenda.agregarEvento(eventoEmpresarial);

        agenda.mostrarContactos();

        agenda.mostrarEventos();
    }
}
