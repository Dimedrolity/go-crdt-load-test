package report

import (
	"fmt"
	"strings"
	"time"
)

type StatisticReport struct {
	Operation Operation
	Count     int
	Mean      time.Duration
	P25       time.Duration
	Median    time.Duration
	P75       time.Duration
}

func (r *StatisticReport) String() string {
	// TODO use template/text
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Operation: %s\n", r.Operation))
	builder.WriteString(fmt.Sprintf("Mean: %d\n", r.Mean.Microseconds()))
	builder.WriteString(fmt.Sprintf("P25: %d\n", r.P25.Microseconds()))
	builder.WriteString(fmt.Sprintf("Median: %d\n", r.Median.Microseconds()))
	builder.WriteString(fmt.Sprintf("P75: %d\n", r.P75.Microseconds()))
	return builder.String()
}
