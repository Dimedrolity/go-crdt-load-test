// Package loader loads GCounter using gcounter and schedule packages.
package loader

import (
	"go-crdt-load-test/gcounter"
	"go-crdt-load-test/report"
	"go-crdt-load-test/schedule"
	"time"
)

type Loader struct {
	config    Config
	counter   gcounter.GCounter
	scheduler schedule.Scheduler[string]
}

func NewLoader(config Config, counter gcounter.GCounter, scheduler schedule.Scheduler[string]) *Loader {
	return &Loader{
		config:    config,
		counter:   counter,
		scheduler: scheduler,
	}
}

func (l *Loader) Load() (report.ResponseSeries, error) {
	var rep report.ResponseSeries

	for i := 0; i < l.config.CountsCount; i++ {
		for j := 0; j < l.config.IncsPerCountCall; j++ {
			address := l.scheduler.Next()

			start := time.Now()
			err := l.counter.Inc(address)
			if err != nil {
				return nil, err
			}

			finish := time.Now()
			delta := finish.Sub(start)
			reportItem := report.Response{
				Operation: report.OperationInc,
				Address:   address,
				Time:      delta,
			}
			rep = append(rep, reportItem)
		}

		address := l.scheduler.Next()
		start := time.Now()
		_, err := l.counter.GetCount(address)
		if err != nil {
			return nil, err
		}
		delta := time.Since(start)

		reportItem := report.Response{
			Operation: report.OperationCount,
			Address:   address,
			Time:      delta,
		}
		rep = append(rep, reportItem)
	}

	return rep, nil
}
