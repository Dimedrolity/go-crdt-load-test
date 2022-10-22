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

func (series ResponseSeries) FilterByOperation(op Operation) ResponseSeries {
	var result ResponseSeries
	for _, item := range series {
		if item.Operation == op {
			result = append(result, item)
		}
	}
	return result
}
