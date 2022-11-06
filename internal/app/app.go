package app

import (
	"fmt"
	"strconv"

	"go-crdt-load-test/gcounter"
	"go-crdt-load-test/loader"
	"go-crdt-load-test/report"
	"go-crdt-load-test/schedule"
	"go-crdt-load-test/statistic"
)

func Run(config *loader.Config) error {
	const host = "http://localhost:"
	var gcounters []gcounter.GCounter
	for port := config.StartPort; port <= config.EndPort; port++ {
		gcounters = append(gcounters, gcounter.NewHttp(host+strconv.Itoa(port)))
	}

	rr := schedule.NewRoundRobin(gcounters)

	start := config.IncsPerCountCall[0]
	stop := config.IncsPerCountCall[1]
	step := config.IncsPerCountCall[2]

	for incsCount := start; incsCount <= stop; incsCount += step {
		l := loader.NewLoader(config.CountsCount, incsCount, rr)
		responseSeries, err := l.Load()
		if err != nil {
			return err
		}

		const dirName = "experiments/"
		err = report.WriteSeriesToFile(responseSeries, fmt.Sprintf(dirName+"report-%d-%d-%d.txt", len(gcounters), config.CountsCount, incsCount))
		if err != nil {
			return err
		}

		incStats := statistic.CalcIncStats(responseSeries)
		err = report.WriteStatsToFile(incStats, fmt.Sprintf(dirName+"inc-%d-%d-%d.txt", len(gcounters), config.CountsCount, incsCount))
		if err != nil {
			return err
		}

		countStats := statistic.CalcCountStats(responseSeries)
		err = report.WriteStatsToFile(countStats, fmt.Sprintf(dirName+"count-%d-%d-%d.txt", len(gcounters), config.CountsCount, incsCount))
		if err != nil {
			return err
		}
	}

	return nil
}
