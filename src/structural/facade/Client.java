package structural.facade;

public class Client {

    public static void main(String[] args) {

        // Without Facade
        Printer printer = new Printer();
        printer.setPaper("A4");
        printer.setColor(true);
        printer.setDocumentType("pdf");
        printer.setText("Text 1");
        printer.print();

        //With Facade
        FacadeNormalPrinter printer1 = new FacadeNormalPrinter("Text 2");
        printer1.print();


    }


}
