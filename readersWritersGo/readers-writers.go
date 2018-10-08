// Implementation based on https://github.com/TheOriginalAlex/os-hw
package main

import (
	"os"
	"strconv"
	"sync"
)

func main() {
	// get the number of readers and writers from arguments
	nReaders, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	nWriters, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	// create the shared resource
	resource := NewResource("No Content")

	// wait until all goroutines are done
	var wg sync.WaitGroup
	wg.Add(nReaders + nWriters)
	defer wg.Wait()

	// launch the readers
	for i := 0; i < nReaders; i++ {
		go NewReader(i, resource, &wg).read()
	}

	// launch the writers
	for i := 0; i < nWriters; i++ {
		go NewWriter(i, resource, &wg).write()
	}
}
