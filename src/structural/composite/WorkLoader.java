package structural.composite;

import javax.imageio.IIOException;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStream;
import java.util.*;

public class WorkLoader {

    protected Properties properties = new Properties();

    public WorkLoader(String fileName) {
        try (InputStream input = new FileInputStream(fileName)){
            properties.load(input);
        } catch (IIOException | FileNotFoundException exp) {
            exp.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public Properties getProperties(){
        return properties;
    }

    public List<Work> getWorkList(){
        List<Work> workList = new ArrayList<>();
        Set<Object> keys = properties.keySet();
        for (Object key: keys) {
            String workType = key.toString().substring("Calculate".length() + 1).toUpperCase();
            String values = properties.getProperty(key.toString());
            Work work = new Work(Calculator.valueOf(workType), Arrays.asList(values.split(", ")));
            workList.add(work);
        }
        return workList;
    }
}
