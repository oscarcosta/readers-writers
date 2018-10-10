package main

import (
	"fmt"
	"sync"

	log "github.com/Sirupsen/logrus"
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

	log.Info("Writer starting ", w.id)

	w.resource.write(w.id, fmt.Sprintf("Content #%d", w.id))

	log.Info("Writer done ", w.id)
}
