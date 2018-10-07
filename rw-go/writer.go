package main

import (
	"fmt"
	"sync"
)

// Writer definition
type Writer struct {
	id       int
	resource *Resource
	wg       *sync.WaitGroup
}

// NewWriter creates a Writer
func NewWriter(id int, resource *Resource, wg *sync.WaitGroup) *Writer {
	return &Writer{id: id, resource: resource, wg: wg}
}

func (w *Writer) write() {
	defer w.wg.Done()

	fmt.Printf("Writer %d starting\n", w.id)

	w.resource.write(w.id, fmt.Sprintf("Content #%d", w.id))

	fmt.Printf("Writer %d is done\n", w.id)
}
