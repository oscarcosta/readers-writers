# Readers-Writers Problem

### Two different solutions to solve the Readers-Writers Problem from The Little Book of Semaphores

Both solutions are implemented in a similar way in different languages (Java and Go). In this implementation a Resource allows the writers write and the readers read. To do this, the Resource is using Mutexes and a *Lightswitch* component to guarantee the mutual exclusion on the resource as described in the book.

+ **readersWritersGo** is the Go implementation

+ **readersWritersJava** is the Java implementation
