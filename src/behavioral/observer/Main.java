package behavioral.observer;

public class Main {
    public static void main(String[] args) {

        Subject subject = new Subject();

        new HexaObserver(subject);
        new OctalObserver(subject);
        new BinaryObserver(subject);

        System.out.println("First stage change: 15");
        subject.setState(15);

        System.out.println("Second stage change: 10");
        subject.setState(10);
    }
}
