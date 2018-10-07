package main

// Implements the Lightswitch pattern

import "sync"

// Lightswitch definition
type Lightswitch struct {
	counter int
	mutex   sync.Mutex
}

// Lock locks the semaphore if its the first thread to call this function
func (l *Lightswitch) Lock(semaphore *sync.Mutex) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.counter++
	if l.counter == 1 {
		semaphore.Lock()
	}
}

// Unlock unlocks the semaphore when the last thread calls this function
func (l *Lightswitch) Unlock(semaphore *sync.Mutex) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.counter--
	if l.counter == 0 {
		semaphore.Unlock()
	}
}
