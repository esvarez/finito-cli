package structural.bridge;

public abstract class Vehicle {
    protected Workshop workShop1;
    protected Workshop workShop2;

    protected Vehicle (Workshop workshop1, Workshop worksho2){
        this.workShop1 = workshop1;
        this.workShop2 = worksho2;
    }

    abstract public void manufacture();
}
