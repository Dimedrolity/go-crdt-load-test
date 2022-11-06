// Package loader loads GCounter using gcounter and schedule packages.
package loader

import (
	"go-crdt-load-test/gcounter"
	"go-crdt-load-test/report"
	"go-crdt-load-test/schedule"
	"time"
)

type Loader struct {
	countsCount      int
	incsPerCountCall int
	scheduler        schedule.Scheduler[gcounter.GCounter]
}

func NewLoader(countsCount int, incsPerCountCall int, scheduler schedule.Scheduler[gcounter.GCounter]) *Loader {
	return &Loader{
		countsCount:      countsCount,
		incsPerCountCall: incsPerCountCall,
		scheduler:        scheduler,
	}
}

func (l *Loader) Load() (report.ResponseSeries, error) {
	var rep report.ResponseSeries

	for i := 0; i < l.countsCount; i++ {
		for j := 0; j < l.incsPerCountCall; j++ {
			counter := l.scheduler.Next()

			start := time.Now()
			err := counter.Inc()
			if err != nil {
				return nil, err
			}

			finish := time.Now()
			delta := finish.Sub(start)
			reportItem := report.Response{
				Operation: report.OperationInc,
				Address:   counter.Name(),
				Time:      delta,
			}
			rep = append(rep, reportItem)
		}

		counter := l.scheduler.Next()
		start := time.Now()
		_, err := counter.GetCount()
		if err != nil {
			return nil, err
		}
		delta := time.Since(start)

		reportItem := report.Response{
			Operation: report.OperationCount,
			Address:   counter.Name(),
			Time:      delta,
		}
		rep = append(rep, reportItem)
	}

	return rep, nil
}
