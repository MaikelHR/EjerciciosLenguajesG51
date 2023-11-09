package Ejercicio1;

public class Linea extends ObjetoGeometrico{
    private int x2;
    private int y2;

    public Linea(int posX, int posY, String color, int x2, int y2) {
        super(posX, posY, color, Math.abs(posX - x2), Math.abs(posY - y2));
        this.x2 = x2;
        this.y2 = y2;
    }

    public int getX2() {
        return this.x2;
    }

    public void setX2(int x2) {
        this.x2 = x2;
    }

    public int getY2() {
        return this.y2;
    }

    public void setY2(int y2) {
        this.y2 = y2;
    }

}
