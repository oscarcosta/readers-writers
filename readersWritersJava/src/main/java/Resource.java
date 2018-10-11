import java.util.concurrent.Semaphore;

public class Resource {

    private String content;
    private final Lightswitch readSwitch = new Lightswitch();
    private final Semaphore roomEmpty = new Semaphore(1);
    private final Semaphore turnstile = new Semaphore(1);

    public Resource(String initContent) {
        this.content = initContent;
    }

    public String write(int id, String data) throws InterruptedException {
        try {
            turnstile.acquire();
            roomEmpty.acquire();

            // critical section
            //System.out.printf("--Writer %d entered the room\n", id);

            //Thread.sleep(100);
            content = data;
            //System.out.printf("--Writer %d wrote '%s'\n", id, content);
            return content;
        } finally {
            turnstile.release();
            roomEmpty.release();
            //System.out.printf("--Writer %d left the room\n", id);
        }
    }

    public String read(int id) throws InterruptedException {
        try {
            turnstile.acquire();
            turnstile.release();
            readSwitch.lock(roomEmpty);

            // critical section
            //System.out.printf("--Reader %d entered the room\n", id);

            //Thread.sleep(100);
            //System.out.printf("--Reader %d read '%s'\n", id, content);
            return content;
        } finally {
            readSwitch.unlock(roomEmpty);
            //System.out.printf("--Reader %d left the room\n", id);
        }
    }
}