package creational.abstractfactory;

import java.util.List;
import java.util.Random;

public class Main {

    @FunctionalInterface
    interface GUIFactory {
      public Button createButton();
    }

    private static GUIFactory factory(String appearence) {
        switch (appearence) {
            case "osx":
                return OSXButton::new;
            case "win":
                return WinButton::new;
            default:
                throw new IllegalArgumentException("unknown " + appearence);
        }
    }

    public static void main(String[] args) {
        var randomAppearance = List.of("osx", "win").get(new Random().nextInt(2));

        var factory = factory(randomAppearance);

        var button = factory.createButton();
        button.paint();
    }
}
