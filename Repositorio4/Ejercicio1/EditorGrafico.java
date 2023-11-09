package Ejercicio1;

public class EditorGrafico {
    public static void main(String[] args) {
        // Crear instancias de objetos representables, círculos, grupos, etc.
        ObjetoRepresentable objeto1 = new ObjetoRepresentable(10, 20);
        Circulo circulo1 = new Circulo(30, 40, 15);
        Grupo grupo1 = new Grupo(50, 60, null);
        // Agregar elementos al grupo
        grupo1.agregarObjeto(objeto1);
        grupo1.agregarObjeto(circulo1);

        // Crear una hoja y agregar elementos a la hoja
        Hoja hoja = new Hoja(null);
        hoja.agregarElemento(objeto1);
        hoja.agregarElemento(circulo1);
        hoja.agregarElemento(grupo1);

        for (ObjetoRepresentable elemento : hoja.elementosEnHoja) {
            elemento.dibujar();
        }
    }
}
