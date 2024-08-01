import javax.print.*;
import javax.print.attribute.*;
import javax.print.attribute.standard.*;
import java.io.*;

public class DotMatrixPrinter {
    public static void main(String[] args) {
        // Mencari printer yang tersedia
        PrintService[] printServices = PrintServiceLookup.lookupPrintServices(null, null);

        if (printServices.length == 0) {
            System.out.println("Tidak ada printer yang ditemukan.");
            return;
        }

        // Memilih printer (disesuaikan dengan nama printer yang sesuai)
        PrintService printer = null;
        for (PrintService ps : printServices) {
            if (ps.getName().equalsIgnoreCase("EPSON LX-300+II")) {
                printer = ps;
                break;
            }
        }

        if (printer == null) {
            System.out.println("Printer Epson LX-300+II tidak ditemukan.");
            return;
        }

        // Data yang akan dicetak
        String dataToPrint = "─│┌┐└┘├┤┬┴┼\n";
        
        // Membuat dokumen yang akan dicetak
        DocPrintJob job = printer.createPrintJob();
        Doc doc = new SimpleDoc(dataToPrint.getBytes(), DocFlavor.BYTE_ARRAY.AUTOSENSE, null);
        
        try {
            job.print(doc, null);
            System.out.println("Pencetakan berhasil.");
        } catch (PrintException e) {
            e.printStackTrace();
            System.out.println("Terjadi kesalahan saat mencetak.");
        }
    }
}
