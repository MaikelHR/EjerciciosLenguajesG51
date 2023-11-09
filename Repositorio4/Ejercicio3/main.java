package Ejercicio3;

public class main {
    public static void main(String[] args) {
        // Obtén la instancia de la Agenda utilizando el Singleton
        Agenda agenda = Agenda.getInstance();

        // Utiliza la AgendaFactory para crear contactos y eventos
        AgendaF factory = new ContactoF();

        Contacto contactoPersonal = factory.crearContacto("Luna", "Miranda", "70854467", 22, "Personal");
        Contacto contactoEmpresarial = factory.crearContacto("Max", "Lopez", "708904467", 25, "Empresarial");

        factory = new EventoF();
        Evento eventoPersonal = factory.crearEvento("Cumpleannos", "18/09/23", "Personal");
        Evento eventoEmpresarial = factory.crearEvento("Reunión Zoom", "22/10/23", "Empresarial");

        // Agregar contactos y eventos a la agenda
        agenda.agregarContacto(contactoPersonal);
        agenda.agregarContacto(contactoEmpresarial);
        agenda.agregarEvento(eventoPersonal);
        agenda.agregarEvento(eventoEmpresarial);

        // Mostrar contactos y eventos
        agenda.mostrarContactos();
        agenda.mostrarEventos();
    }
}

/*
 * Eager Singleton:

En un "eager singleton", la instancia de la clase se crea en el momento en que el programa se inicia o en el momento de la carga de la clase, independientemente de si se necesitará o no en el transcurso del programa.
Esto significa que la instancia se crea anticipadamente, sin importar si se utiliza inmediatamente o no. La instancia está lista y disponible desde el principio.
Es útil cuando se está seguro de que se usará la instancia en algún momento temprano y se quiere evitar la sobrecarga de crearla más tarde.

Lazy Singleton:

En un "lazy singleton", la instancia de la clase se crea solo cuando se necesita por primera vez, es decir, se crea "perezosamente".
Esto significa que la instancia no se crea al inicio del programa, sino cuando se llama por primera vez al método que la crea. Si no se necesita, la instancia nunca se crea.
Es útil en situaciones en las que la creación de la instancia puede ser costosa en términos de recursos o tiempo, y se prefiere retrasar su creación hasta que realmente sea necesaria.

A mi parecer es mejor utilizar el singleton ya que las instacias se crean solo una vez
 */
