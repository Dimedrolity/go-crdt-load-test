package main

import (
	"fmt"
	"go-crdt-load-test/gcounter"
	"go-crdt-load-test/loader"
	"go-crdt-load-test/report"
	"go-crdt-load-test/schedule"
	"go-crdt-load-test/statistic"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func main() {
	yamlFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	var loaderConfig loader.Config
	err = yaml.Unmarshal(yamlFile, &loaderConfig)
	if err != nil {
		log.Fatal(err)
	}

	// TODO move addresses to config
	rr := schedule.NewRoundRobin([]gcounter.GCounter{
		gcounter.NewHttp("http://localhost:8000"),
		gcounter.NewHttp("http://localhost:8001"),
		gcounter.NewHttp("http://localhost:8002"),
	})
	// TODO move all logic from main to don't repeat in E2E test
	l := loader.NewLoader(loaderConfig, rr)
	responseSeries, err := l.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = report.WriteSeriesToFile(responseSeries, fmt.Sprintf("report-%d-%d.txt", loaderConfig.CountsCount, loaderConfig.IncsPerCountCall))
	if err != nil {
		log.Fatal(err)
	}

	incStats := statistic.CalcIncStats(responseSeries)
	err = report.WriteStatsToFile(incStats, fmt.Sprintf("inc-%d-%d.txt", loaderConfig.CountsCount, loaderConfig.IncsPerCountCall))
	if err != nil {
		log.Fatal(err)
	}

	countStats := statistic.CalcCountStats(responseSeries)
	err = report.WriteStatsToFile(countStats, fmt.Sprintf("count-%d-%d.txt", loaderConfig.CountsCount, loaderConfig.IncsPerCountCall))
	if err != nil {
		log.Fatal(err)
	}
}
