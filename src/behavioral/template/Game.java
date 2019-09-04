package behavioral.template;

public abstract class Game {

    abstract void initialize();
    abstract void startPlay();
    abstract void endPLay();

    //Template method
    public final void play(){

        initialize();

        startPlay();

        endPLay();
    }
}
