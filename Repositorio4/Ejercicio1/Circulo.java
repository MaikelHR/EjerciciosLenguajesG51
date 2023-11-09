package Ejercicio1;

public class Circulo extends ObjetoRepresentable{
    private double radio;

    public Circulo(int posX, int posY, double radio) {
        super(posX, posY);
        this.radio = radio;
    }


    public double getRadio() {
        return this.radio;
    }

    public void setRadio(double radio) {
        this.radio = radio;
    }

}
