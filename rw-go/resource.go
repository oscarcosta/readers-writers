package main

import (
	"fmt"
	"sync"
)

// Resource definition
type Resource struct {
	content    string
	readSwitch Lightswitch
	roomEmpty  sync.Mutex
	turnstile  sync.Mutex
}

// NewResource create a Resource
func NewResource(initContent string) *Resource {
	return &Resource{content: initContent}
}

func (r *Resource) write(id int, data string) string {
	r.turnstile.Lock()
	r.roomEmpty.Lock()
	// critical section occurs here
	defer r.turnstile.Unlock()
	defer r.roomEmpty.Unlock()

	fmt.Printf("--Writer %d entered the room\n", id)
	defer fmt.Printf("--Writer %d left the room\n", id)

	//time.Sleep(100 * time.Millisecond)
	r.content = data
	fmt.Printf("----Writer %d wrote '%s'\n", id, r.content)
	return r.content
}

func (r *Resource) read(id int) string {
	r.turnstile.Lock()
	r.turnstile.Unlock()
	r.readSwitch.Lock(&r.roomEmpty)
	// critical section occurs here
	defer r.readSwitch.Unlock(&r.roomEmpty)

	fmt.Printf("--Reader %d entered the room\n", id)
	defer fmt.Printf("--Reader %d left the room\n", id)

	//time.Sleep(100 * time.Millisecond)
	fmt.Printf("----Reader %d read '%s'\n", id, r.content)
	return r.content
}
