package Ejercicio1;

public class Elipse extends ObjetoGeometrico{
    private double radioMayor;
    private double radioMenor;

    public Elipse(int posX, int posY, String color, double radioMayor, double radioMenor) {
        super(posX, posY, color, radioMayor * 2, radioMenor * 2);
        this.radioMayor = radioMayor;
        this.radioMenor = radioMenor;
    }

    public double getRadioMayor() {
        return this.radioMayor;
    }

    public void setRadioMayor(double radioMayor) {
        this.radioMayor = radioMayor;
    }

    public double getRadioMenor() {
        return this.radioMenor;
    }

    public void setRadioMenor(double radioMenor) {
        this.radioMenor = radioMenor;
    }

}
