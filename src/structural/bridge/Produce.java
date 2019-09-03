package structural.bridge;

public class Produce implements Workshop {
    @Override
    public void work() {
        System.out.println("Produced");
    }
}
