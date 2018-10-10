package main

import (
	"sync"

	log "github.com/Sirupsen/logrus"
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

	log.Info("--Writer entered the room ", id)
	defer log.Info("--Writer left the room ", id)

	//time.Sleep(100 * time.Millisecond)
	r.content = data
	log.Info("----Writer ", id, " wrote ", r.content)
	return r.content
}

func (r *Resource) read(id int) string {
	r.turnstile.Lock()
	r.turnstile.Unlock()
	r.readSwitch.Lock(&r.roomEmpty)
	// critical section occurs here
	defer r.readSwitch.Unlock(&r.roomEmpty)

	log.Info("--Reader entered the room ", id)
	defer log.Info("--Reader left the room ", id)

	//time.Sleep(100 * time.Millisecond)
	log.Info("----Reader ", id, " read ", r.content)
	return r.content
}
