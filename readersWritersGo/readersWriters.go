// Implementation based on https://github.com/TheOriginalAlex/os-hw
package main

import (
	"flag"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"

	log "github.com/Sirupsen/logrus"
)

func main() {
	// profile flags
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")
	debug := flag.Bool("debug", false, "Display debug logs")

	// get the number of readers and writers from arguments
	nReaders := flag.Int("readers", 5, "defines the number of readers")
	nWriters := flag.Int("writers", 5, "defines the number of writers")

	flag.Parse()

	// utilize the max num of cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	// log level
	if *debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	// CPU Profile
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	// create the shared resource
	resource := NewResource("No Content")

	// wait until all goroutines are done
	var wg sync.WaitGroup
	wg.Add(*nReaders + *nWriters)
	defer wg.Wait()

	// launch the readers
	for i := 0; i < *nReaders; i++ {
		go NewReader(i, resource, &wg).read()
	}

	// launch the writers
	for i := 0; i < *nWriters; i++ {
		go NewWriter(i, resource, &wg).write()
	}

	// Memory Profile
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
		f.Close()
	}
}
