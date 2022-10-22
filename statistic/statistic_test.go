package statistic_test

import (
	"github.com/stretchr/testify/assert"
	"go-crdt-load-test/report"
	"go-crdt-load-test/statistic"
	"testing"
	"time"
)

func TestMean(t *testing.T) {
	var series report.ResponseSeries = []report.Response{
		{Time: 1 * time.Microsecond},
		{Time: 2 * time.Microsecond},
		{Time: 7 * time.Microsecond},
	}
	mean := statistic.Mean(series)
	assert.Equal(t, 3*time.Microsecond, mean)
}

func TestMedian(t *testing.T) {
	var series report.ResponseSeries = []report.Response{
		{Time: 7 * time.Microsecond},
		{Time: 1 * time.Microsecond},
		{Time: 2 * time.Microsecond},
	}
	median := statistic.Median(series)
	assert.Equal(t, 2*time.Microsecond, median)

	// Source slice is unsorted
	assert.ElementsMatch(t, series, []report.Response{
		{Time: 7 * time.Microsecond},
		{Time: 1 * time.Microsecond},
		{Time: 2 * time.Microsecond},
	})
}

func TestP25(t *testing.T) {
	var series report.ResponseSeries = []report.Response{
		{Time: 1 * time.Microsecond},
		{Time: 2 * time.Microsecond},
		{Time: 3 * time.Microsecond},
		{Time: 4 * time.Microsecond},
		{Time: 5 * time.Microsecond},
	}
	p25 := statistic.P25(series)
	assert.Equal(t, 2*time.Microsecond, p25)
}

func TestP75(t *testing.T) {
	var series report.ResponseSeries = []report.Response{
		{Time: 1 * time.Microsecond},
		{Time: 2 * time.Microsecond},
		{Time: 3 * time.Microsecond},
		{Time: 4 * time.Microsecond},
		{Time: 5 * time.Microsecond},
	}
	p75 := statistic.P75(series)
	assert.Equal(t, 4*time.Microsecond, p75)
}
