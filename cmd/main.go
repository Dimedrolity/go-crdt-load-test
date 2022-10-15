package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"

	"go-crdt-load-test/loader"
	"go-crdt-load-test/report"
	"go-crdt-load-test/schedule"
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
	rep, err := l.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = report.WriteToFile(rep, "report.txt")
	if err != nil {
		log.Fatal(err)
	}
}
