package structural.adapter;

public class BirdAdapter implements ToyDuck{

    Bird bird;
    public BirdAdapter(Bird bird){
        this.bird = bird;
    }

    public void squeak() {
        bird.makeSound();
    }
}
