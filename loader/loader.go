// Package loader loads GCounters using client and schedule packages.
package loader

import (
	"go-crdt-load-test/client"
	"go-crdt-load-test/report"
	"go-crdt-load-test/schedule"
	"time"
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

func (l *Loader) Load() (report.Report, error) {

	var rep report.Report

	for i := 0; i < l.config.CountsCount; i++ {
		for j := 0; j < l.config.IncsPerCountCall; j++ {
			address := l.scheduler.Next()

			start := time.Now()
			err := client.Inc(address)
			if err != nil {
				return nil, err
			}

			finish := time.Now()
			delta := finish.Sub(start)
			reportItem := report.Item{
				Operation:    report.IncOperation,
				Address:      address,
				ResponseTime: delta,
			}
			rep = append(rep, reportItem)
		}

		address := l.scheduler.Next()
		start := time.Now()
		_, err := client.GetCount(address)
		if err != nil {
			return nil, err
		}
		finish := time.Now()
		delta := finish.Sub(start) // TODO refactor time.Since()

		reportItem := report.Item{
			Operation:    report.CountOperation,
			Address:      address,
			ResponseTime: delta,
		}
		rep = append(rep, reportItem)
	}

	return rep, nil
}
