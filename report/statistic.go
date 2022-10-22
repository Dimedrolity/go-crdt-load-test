package report

import (
	"fmt"
	"strings"
)

type StatisticReport struct {
	Operation Operation
	Mean      int64
}

func (r *StatisticReport) String() string {
	// TODO use template/text
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("Operation: %s\n", r.Operation))
	builder.WriteString(fmt.Sprintf("Mean: %d", r.Mean))
	return builder.String()
}
