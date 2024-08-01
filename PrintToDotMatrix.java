import java.io.*;

public class PrintToDotMatrix {

    public static void main(String[] args) {
        // File path to the text file
        String filePath = "example.txt";
        
        // Printer name or port (this might be different depending on your setup)
        String printerName = "LPT1"; // Example for parallel port
        
        try {
            // Read the text file
            BufferedReader reader = new BufferedReader(new FileReader(filePath));
            StringBuilder content = new StringBuilder();
            String line;
            while ((line = reader.readLine()) != null) {
                content.append(line).append("\n");
            }
            reader.close();
            
            // Send the content to the printer
            PrintWriter printerWriter = new PrintWriter(new OutputStreamWriter(new FileOutputStream(printerName)));
            printerWriter.print(content.toString());
            printerWriter.close();
            
            System.out.println("Printed successfully!");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
