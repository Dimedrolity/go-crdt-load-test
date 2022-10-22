package report

import (
	"fmt"
	"strings"
	"time"
)

type Operation string

const OperationInc Operation = "INC"
const OperationCount Operation = "COUNT"

type Response struct {
	Operation Operation
	Address   string
	Time      time.Duration
}

func (item *Response) String() string {
	return fmt.Sprintf("%s: %s %s", item.Operation, item.Address, item.Time)
}

type ResponseSeries []Response

func (series ResponseSeries) String() string {
	builder := strings.Builder{}
	for _, item := range series {
		builder.WriteString(item.String() + "\n")
	}
	return builder.String()
}

func (series ResponseSeries) CalcAverage() int64 {
	var sum int64

	for _, item := range series {
		sum += item.Time.Microseconds()
	}

	avg := sum / int64(len(series))
	return avg
}
