package creational.prototype;

public class Main {
    public static void main(String[] args) {
        ShapeCache.loadCache();

        Shape cloneShape = (Shape) ShapeCache.getShape("1");
        System.out.println("Shape : " + cloneShape.getType());

        Shape cloneShape2 = (Shape) ShapeCache.getShape("2");
        System.out.println("Shape : " + cloneShape2.getType());

        Shape clonedShaped3 = (Shape) ShapeCache.getShape("3");
        System.out.println("Shape : " + clonedShaped3.getType());
    }
}
