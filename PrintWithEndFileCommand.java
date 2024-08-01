import java.io.*;
import javax.print.*;

public class PrintWithEndFileCommand {

    public static void main(String[] args) {
        // Path to the text file
        String filePath = "example.txt";
        
        // Name of the printer to use
        String printerName = "Your Network Printer Name";

        // Print the file with end-of-file command
        printFileWithEndFileCommand(filePath, printerName);
    }

    private static void printFileWithEndFileCommand(String filePath, String printerName) {
        try {
            // Create a FileInputStream for the file
            FileInputStream fis = new FileInputStream(filePath);

            // Read the file content into a byte array
            byte[] fileContent = fis.readAllBytes();
            fis.close();

            // Define ESC/POS command for initialization (ESC @)
            byte[] escPosInit = new byte[] { 0x1B, 0x40 }; // ESC @

            // Define end-of-job command (Form Feed)
            byte[] endOfJob = new byte[] { 0x0C }; // Form Feed (FF)

            // Combine ESC/POS command, file content, and end-of-job command
            ByteArrayOutputStream baos = new ByteArrayOutputStream();
            baos.write(escPosInit);
            baos.write(fileContent);
            baos.write(endOfJob);

            // Create a byte array from the output stream
            byte[] dataToPrint = baos.toByteArray();

            // Get the print service (you might need to specify a specific printer)
            PrintService[] printServices = PrintServiceLookup.lookupPrintServices(DocFlavor.BYTE_ARRAY.AUTOSENSE, null);
            PrintService selectedPrinter = null;
            for (PrintService printService : printServices) {
                if (printService.getName().equalsIgnoreCase(printerName)) {
                    selectedPrinter = printService;
                    break;
                }
            }

            if (selectedPrinter != null) {
                // Create a Doc from the byte array
                Doc doc = new SimpleDoc(dataToPrint, DocFlavor.BYTE_ARRAY.AUTOSENSE, null);

                // Create a DocPrintJob
                DocPrintJob printJob = selectedPrinter.createPrintJob();

                // Print the document
                printJob.print(doc, null);

                System.out.println("Print job sent to " + printerName + " successfully.");
            } else {
                System.out.println("Printer not found: " + printerName);
            }

        } catch (IOException | PrintException e) {
            e.printStackTrace();
        }
    }
}
