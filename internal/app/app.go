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
	l := loader.NewLoader(config, rr)
	responseSeries, err := l.Load()
	if err != nil {
		return err
	}

	err = report.WriteSeriesToFile(responseSeries, fmt.Sprintf("report-%d-%d-%d.txt", len(gcounters), config.CountsCount, config.IncsPerCountCall))
	if err != nil {
		return err
	}

	incStats := statistic.CalcIncStats(responseSeries)
	err = report.WriteStatsToFile(incStats, fmt.Sprintf("inc-%d-%d-%d.txt", len(gcounters), config.CountsCount, config.IncsPerCountCall))
	if err != nil {
		return err
	}

	countStats := statistic.CalcCountStats(responseSeries)
	err = report.WriteStatsToFile(countStats, fmt.Sprintf("count-%d-%d-%d.txt", len(gcounters), config.CountsCount, config.IncsPerCountCall))
	if err != nil {
		return err
	}

	return nil
}
