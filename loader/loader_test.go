package loader_test

import (
	"github.com/stretchr/testify/assert"
	"go-crdt-load-test/client"
	"go-crdt-load-test/report"
	"go-crdt-load-test/statistic"
	"testing"

	"go-crdt-load-test/loader"
	"go-crdt-load-test/schedule"
)

// TestLoad is E2E test
// TODO write unit tests
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

	responseSeries, err := l.Load()
	assert.Nil(t, err)
	err = report.WriteSeriesToFile(responseSeries, "report.txt")
	assert.Nil(t, err)

	incStats := statistic.CalcIncStats(responseSeries)
	err = report.WriteStatsToFile(incStats, "inc.txt")
	assert.Nil(t, err)

	countStats := statistic.CalcCountStats(responseSeries)
	err = report.WriteStatsToFile(countStats, "count.txt")
	assert.Nil(t, err)

	count, err := client.GetCount("http://localhost:8002")
	assert.Nil(t, err)
	assert.Equal(t, 90, count)
}
