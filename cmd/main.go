package main

import (
	"go-crdt-load-test/loader"
	"go-crdt-load-test/report"
	"go-crdt-load-test/schedule"
	"log"
)

func main() {
	// Call 9 Increments and 1 Count, 10 times.
	// 9*10 = 90 is a total Increments counts.
	// 1*10 = 10 is a total Count calls count.
	// (9+1)*10 = 100 is a total requests count.
	loaderConfig := loader.Config{
		CountsCount:      10,
		IncsPerCountCall: 9,
	}
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
