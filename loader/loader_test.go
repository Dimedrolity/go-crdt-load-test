package loader_test

import (
	"github.com/stretchr/testify/assert"
	"go-crdt-load-test/client"
	"log"
	"testing"

	"go-crdt-load-test/loader"
	"go-crdt-load-test/schedule"
)

func TestLoad(t *testing.T) {
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

	err := l.Load()
	if err != nil {
		log.Fatal(err)
	}

	count, err := client.GetCount("http://localhost:8002")
	assert.Nil(t, err)
	assert.Equal(t, 90, count)
}
