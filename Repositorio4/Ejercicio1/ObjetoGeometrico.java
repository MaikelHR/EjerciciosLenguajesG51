package Ejercicio1;

public class ObjetoGeometrico extends ObjetoRepresentable {
    private String color;
    private double ancho;
    private double alto;

    public ObjetoGeometrico(int posX, int posY, String color, double ancho, double alto) {
        super(posX, posY);
        this.color = color;
        this.ancho = ancho;
        this.alto = alto;
    }


    public String getColor() {
        return this.color;
    }

    public void setColor(String color) {
        this.color = color;
    }

    public double getAncho() {
        return this.ancho;
    }

    public void setAncho(double ancho) {
        this.ancho = ancho;
    }

    public double getAlto() {
        return this.alto;
    }

    public void setAlto(double alto) {
        this.alto = alto;
    }

}
