// Package loader loads GCounters using client and schedule packages.
package loader

import (
	"go-crdt-load-test/client"
	"go-crdt-load-test/schedule"
)

type Loader struct {
	// TODO use client as dependency
	//client
	config    Config
	scheduler schedule.Scheduler[string]
}

func NewLoader(config Config, scheduler schedule.Scheduler[string]) *Loader {
	return &Loader{
		config:    config,
		scheduler: scheduler,
	}
}

// TODO measure request time
// TODO aggregate times

func (l *Loader) Load() error {

	for i := 0; i < l.config.CountsCount; i++ {
		for j := 0; j < l.config.IncsPerCountCall; j++ {
			address := l.scheduler.Next()
			err := client.Inc(address)
			if err != nil {
				return err
			}
		}

		address := l.scheduler.Next()
		_, err := client.GetCount(address)
		if err != nil {
			return err
		}
	}

	return nil
}
