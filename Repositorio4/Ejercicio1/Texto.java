package Ejercicio1;

public class Texto extends ObjetoRepresentable{
    private String contenido;
    private String fuente;

    public Texto(int posX, int posY, String contenido, String fuente) {
        super(posX, posY);
        this.contenido = contenido;
        this.fuente = fuente;
    }


    public String getContenido() {
        return this.contenido;
    }

    public void setContenido(String contenido) {
        this.contenido = contenido;
    }

    public String getFuente() {
        return this.fuente;
    }

    public void setFuente(String fuente) {
        this.fuente = fuente;
    }

}
