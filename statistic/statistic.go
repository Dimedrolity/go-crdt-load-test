package statistic

import (
	"go-crdt-load-test/report"
	"sort"
	"time"
)

// TODO use lib https://github.com/montanaflynn/stats, don't create own.

func Mean(series report.ResponseSeries) time.Duration {
	var sum int64

	for _, item := range series {
		sum += item.Time.Microseconds()
	}

	mean := sum / int64(len(series))

	return time.Duration(mean) * time.Microsecond
}

// P25 is 25 percentile.
func P25(series report.ResponseSeries) time.Duration {
	return percentile(series, 25)
}

// Median is 50 percentile.
func Median(series report.ResponseSeries) time.Duration {
	return percentile(series, 50)
}

// P75 is 75 percentile.
func P75(series report.ResponseSeries) time.Duration {
	return percentile(series, 75)
}

// percentile returns element with index n by formula n = (p/100) x N,
// where where N = number of values in the data set, p = percentile.
// TODO check that has no side effect, not sorting source slice.
func percentile(series report.ResponseSeries, p int) time.Duration {
	sort.Slice(series, func(i, j int) bool {
		return series[i].Time < series[j].Time
	})
	n := p * len(series) / 100
	return series[n].Time
}

func CalcIncStats(series report.ResponseSeries) *report.StatisticReport {
	incs := series.FilterByOperation(report.OperationInc)
	mean := Mean(incs)
	median := Median(incs)
	p25 := P25(incs)
	p75 := P75(incs)

	return &report.StatisticReport{
		Operation: report.OperationInc,
		Count:     len(incs),
		Mean:      mean,
		P25:       p25,
		Median:    median,
		P75:       p75,
	}
}

func CalcCountStats(series report.ResponseSeries) *report.StatisticReport {
	counts := series.FilterByOperation(report.OperationCount)
	mean := Mean(counts)
	median := Median(counts)
	p25 := P25(counts)
	p75 := P75(counts)

	return &report.StatisticReport{
		Operation: report.OperationCount,
		Count:     len(counts),
		Mean:      mean,
		P25:       p25,
		Median:    median,
		P75:       p75,
	}
}
