
public class ReadersWriters {

    private final int nReaders;
    private final int nWriters;
    private final Resource resource;

    public ReadersWriters(int nReaders, int nWriters) {
        this.nReaders = nReaders;
        this.nWriters = nWriters;
        this.resource = new Resource("No Content");
    }

    public void execute() {
        // launch the writers
        for (int i = 0; i < nWriters; i++) {
            new Thread(new Writer(i, resource)).start();
        }

        // launch the readers
        for (int i = 0; i < nReaders; i++) {
            new Thread(new Reader(i, resource)).start();
        }
    }
    
    public static void main(String[] args) {
        // get the number of readers and writers from arguments
        int nReaders = Integer.valueOf(args[0]);
        int nWriters = Integer.valueOf(args[1]);
        
        ReadersWriters rw = new ReadersWriters(nReaders, nWriters);
        rw.execute();
    }
}
