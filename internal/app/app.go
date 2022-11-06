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

	incStart := config.IncsPerCountCall[0]
	incStop := config.IncsPerCountCall[1]
	incStep := config.IncsPerCountCall[2]

	countStart := config.CountsCount[0]
	countStop := config.CountsCount[1]
	countStep := config.CountsCount[2]

	const dirName = "experiments/"

	csv, err := report.NewCsvWriter(fmt.Sprintf(dirName+"counts-%d-%v-%v", len(gcounters), config.CountsCount, config.IncsPerCountCall))
	if err != nil {
		return err
	}

	for countsCount := countStart; countsCount <= countStop; countsCount *= countStep {
		for incsCount := incStart; incsCount <= incStop; incsCount *= incStep {
			l := loader.NewLoader(countsCount, incsCount, rr)
			responseSeries, err := l.Load()
			if err != nil {
				return err
			}

			err = report.WriteSeriesToFile(responseSeries, fmt.Sprintf(dirName+"report-%d-%d-%d.txt", len(gcounters), countsCount, incsCount))
			if err != nil {
				return err
			}

			incStats := statistic.CalcIncStats(responseSeries)
			err = report.WriteStatsToFile(incStats, fmt.Sprintf(dirName+"inc-%d-%d-%d.txt", len(gcounters), countsCount, incsCount))
			if err != nil {
				return err
			}

			countStats := statistic.CalcCountStats(responseSeries)
			err = report.WriteStatsToFile(countStats, fmt.Sprintf(dirName+"count-%d-%d-%d.txt", len(gcounters), countsCount, incsCount))
			if err != nil {
				return err
			}

			err = csv.WriteStatsToCsv(countStats, countsCount, incsCount)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
