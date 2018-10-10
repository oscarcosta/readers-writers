package main

import (
	"sync"

	log "github.com/Sirupsen/logrus"
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

	log.Info("Reader starting ", r.id)

	r.resource.read(r.id)

	log.Info("Reader done ", r.id)
}
