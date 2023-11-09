package Ejercicio1;

public class Cuadrado extends ObjetoGeometrico{
    private double lado;

    public Cuadrado(int posX, int posY, String color, double lado) {
        super(posX, posY, color, lado, lado);
        this.lado = lado;
    }

    public double getLado() {
        return this.lado;
    }

    public void setLado(double lado) {
        this.lado = lado;
    }

}
