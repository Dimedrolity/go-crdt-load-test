package main

import (
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
	rr := schedule.NewRoundRobin([]string{
		"http://localhost:8000",
		"http://localhost:8001",
		"http://localhost:8002",
	})

	l := loader.NewLoader(loaderConfig, rr)
	responseSeries, err := l.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = report.WriteSeriesToFile(responseSeries, "report.txt")
	if err != nil {
		log.Fatal(err)
	}

	incStats := statistic.CalcIncStats(responseSeries)
	err = report.WriteStatsToFile(incStats, "inc.txt")
	if err != nil {
		log.Fatal(err)
	}

	countStats := statistic.CalcCountStats(responseSeries)
	err = report.WriteStatsToFile(countStats, "count.txt")
	if err != nil {
		log.Fatal(err)
	}
}
