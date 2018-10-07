package main

import (
	"fmt"
	"sync"
)

// Reader definition
type Reader struct {
	id       int
	resource *Resource
	wg       *sync.WaitGroup
}

// NewReader creates a Reader
func NewReader(id int, resource *Resource, wg *sync.WaitGroup) *Reader {
	return &Reader{id: id, resource: resource, wg: wg}
}

func (r *Reader) read() {
	defer r.wg.Done()

	fmt.Printf("Reader %d starting\n", r.id)

	r.resource.read(r.id)

	fmt.Printf("Reader %d is done\n", r.id)
}
