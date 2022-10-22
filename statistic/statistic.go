package statistic

import (
	"go-crdt-load-test/report"
)

func Mean(series report.ResponseSeries) int64 {
	var sum int64

	for _, item := range series {
		sum += item.Time.Microseconds()
	}

	mean := sum / int64(len(series))
	return mean
}

func CalcIncStats(series report.ResponseSeries) *report.StatisticReport {
	incs := series.FilterByOperation(report.OperationInc)
	mean := Mean(incs)
	return &report.StatisticReport{
		Operation: report.OperationInc,
		Mean:      mean,
	}
}

func CalcCountStats(series report.ResponseSeries) *report.StatisticReport {
	incs := series.FilterByOperation(report.OperationCount)
	mean := Mean(incs)
	return &report.StatisticReport{
		Operation: report.OperationCount,
		Mean:      mean,
	}
}
