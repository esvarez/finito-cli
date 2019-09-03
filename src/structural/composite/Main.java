package structural.composite;

public class Main {

    public static void main(String[] args) {

        Engineer ajay = new Engineer(1001l, "Ajay", "Developer", "ENG");
        Engineer vijay = new Engineer(1002l, "Vijay", "SR. Developer", "ENG");
        Engineer jay = new Engineer(1003l, "Jay", "Lead", "ENG");
        Engineer martin = new Engineer(1004l, "Martin", "QA", "ENG");
        Manager kim = new Manager(1005l, "Kim", "Manager", "ENG");
        Engineer anders = new Engineer(1006l, "Andersen", "Developer", "ENG");
        Manager niels = new Manager(1007l, "Niels", "Sr. Manager", "ENG");
        Engineer robert = new Engineer(1008l, "Robert", "Developer", "ENG");
        Manager rachelle = new Manager(1009l, "Rachelle", "Product Manager", "ENG");
        Engineer shailesh = new Engineer(1010l, "Shailesh", "Engineer", "ENG");
        kim.manages(ajay);
        kim.manages(martin);
        kim.manages(vijay);

        niels.manages(jay);
        niels.manages(anders);
        niels.manages(shailesh);

        rachelle.manages(kim);
        rachelle.manages(robert);
        rachelle.manages(niels);


        WorkLoader workLoad = new WorkLoader("work.properties");
        workLoad.getWorkList().stream().forEach(work -> {
            rachelle.assignWork(rachelle, work);
        });

        rachelle.performWork();

    }
}
