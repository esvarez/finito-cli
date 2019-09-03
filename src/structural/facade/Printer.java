package structural.facade;

public class Printer {

    private String documentType;
    private String paper;
    private boolean color;
    private String text;

    public String getDocumentType() {
        return documentType;
    }

    public void setDocumentType(String documentType) {
        this.documentType = documentType;
    }

    public String getPaper() {
        return paper;
    }

    public void setPaper(String paper) {
        this.paper = paper;
    }

    public boolean isColor() {
        return color;
    }

    public void setColor(boolean color) {
        this.color = color;
    }

    public String getText() {
        return text;
    }

    public void setText(String text) {
        this.text = text;
    }

    public void print(){
        System.out.println("Printing...");
        System.out.println("Paper: " + paper);
        System.out.println("Color: " + color);
        System.out.println(text);
    }
}
