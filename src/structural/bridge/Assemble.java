package structural.bridge;

public class Assemble implements Workshop {

    @Override
    public void work() {
        System.out.println(" And");
        System.out.println(" Assembled");
    }
}
