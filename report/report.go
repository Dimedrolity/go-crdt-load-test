package report

import "time"

const IncOperation = "INC"
const CountOperation = "COUNT"

type Item struct {
	Operation    string
	Address      string
	ResponseTime time.Duration
}

type Report []Item

func (r Report) CalcAverage() int64 {
	var sum int64

	for _, item := range r {
		sum += item.ResponseTime.Microseconds()
	}

	avg := sum / int64(len(r))
	return avg
}
