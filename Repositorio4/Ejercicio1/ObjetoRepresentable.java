package Ejercicio1;

class ObjetoRepresentable {
    
    private int posX;
    private int posY;

    // Constructor
    public ObjetoRepresentable(int posX, int posY) {
        this.posX = posX;
        this.posY = posY;
    }

    // Métodos get/set
    public int getPosX() {
        return posX;
    }

    public void setPosX(int posX) {
        this.posX = posX;
    }

    public int getPosY() {
        return posY;
    }

    public void setPosY(int posY) {
        this.posY = posY;
    }

    // Método para dibujar el objeto
    public void dibujar() {
        System.out.println("Objeto representable dibujado en la posición (" + posX + ", " + posY + ")");
    }
}