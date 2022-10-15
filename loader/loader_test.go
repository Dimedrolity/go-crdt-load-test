package loader_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-crdt-load-test/client"
	"os"
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

	report, err := l.Load()
	assert.Nil(t, err)
	f, err := os.Create("report.txt")
	assert.Nil(t, err)
	defer f.Close()
	for _, item := range report {
		_, err := f.WriteString(fmt.Sprintf("%s: %s %s\n", item.Operation, item.Address, item.ResponseTime))
		assert.Nil(t, err)
	}

	_, err = f.WriteString(fmt.Sprint("avg in mcs\n"))
	_, err = f.WriteString(fmt.Sprintf("avg all: %d\n", report.CalcAverage()))
	assert.Nil(t, err)
	incs := filterInc(report)
	_, err = f.WriteString(fmt.Sprintf("avg inc: %d\n", incs.CalcAverage()))
	assert.Nil(t, err)
	counts := filterCount(report)
	_, err = f.WriteString(fmt.Sprintf("avg count: %d\n", counts.CalcAverage()))
	assert.Nil(t, err)

	count, err := client.GetCount("http://localhost:8002")
	assert.Nil(t, err)
	assert.Equal(t, 90, count)
}

func filterInc(report loader.Report) loader.Report {
	var incs loader.Report
	for _, item := range report {
		if item.Operation == loader.IncOperation {
			incs = append(incs, item)
		}
	}
	return incs
}
func filterCount(report loader.Report) loader.Report {
	var counts loader.Report
	for _, item := range report {
		if item.Operation == loader.CountOperation {
			counts = append(counts, item)
		}
	}
	return counts
}
