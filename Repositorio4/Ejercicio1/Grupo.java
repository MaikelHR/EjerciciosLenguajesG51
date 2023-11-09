package Ejercicio1;

import java.util.List;

public class Grupo extends ObjetoRepresentable {
    private List<ObjetoRepresentable> objetosEnGrupo;

    public Grupo(int posX, int posY, List<ObjetoRepresentable> objetosEnGrupo) {
        super(posX, posY);
        this.objetosEnGrupo = objetosEnGrupo;
    }


    public List<ObjetoRepresentable> getObjetosEnGrupo() {
        return this.objetosEnGrupo;
    }

    public void setObjetosEnGrupo(List<ObjetoRepresentable> objetosEnGrupo) {
        this.objetosEnGrupo = objetosEnGrupo;
    }

    public void agregarObjeto(ObjetoRepresentable objeto) {
        objetosEnGrupo.add(objeto);
    }

    public void removerObjeto(ObjetoRepresentable objeto) {
        objetosEnGrupo.remove(objeto);
    }
}
