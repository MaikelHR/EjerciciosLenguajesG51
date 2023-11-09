package Ejercicio1;

import java.util.List;

public class Hoja {
    List<ObjetoRepresentable> elementosEnHoja;


    public Hoja(List<ObjetoRepresentable> elementosEnHoja) {
        this.elementosEnHoja = elementosEnHoja;
    }

    public List<ObjetoRepresentable> getElementosEnHoja() {
        return this.elementosEnHoja;
    }

    public void setElementosEnHoja(List<ObjetoRepresentable> elementosEnHoja) {
        this.elementosEnHoja = elementosEnHoja;
    }

    public void agregarElemento(ObjetoRepresentable elemento) {
        elementosEnHoja.add(elemento);
    }

    public void removerElemento(ObjetoRepresentable elemento) {
        elementosEnHoja.remove(elemento);
    }
}
