package structural.facade;

public class FacadeNormalPrinter {

    private Printer printer;

    public FacadeNormalPrinter(String text) {
        super();
        printer = new Printer();
        printer.setColor(true);
        printer.setPaper("A4");
        printer.setDocumentType("PDF");
        printer.setText(text);
    }

    public void print(){
        printer.print();
    }
}
