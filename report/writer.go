package report

import (
	"fmt"
	"os"
)

// WriteToFile writes full report to file.
// TODO dont call calculations of statistics.
// TODO change param to Report Object with all info and write to file.
func WriteToFile(series ResponseSeries, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(series.String())
	if err != nil {
		return err
	}

	_, err = f.WriteString(fmt.Sprint("avg in mcs\n"))
	_, err = f.WriteString(fmt.Sprintf("avg ALL: %d\n", series.CalcAverage()))
	if err != nil {
		return err
	}
	incs := filterInc(series)

	_, err = f.WriteString(fmt.Sprintf("avg INC: %d\n", incs.CalcAverage()))
	if err != nil {
		return err
	}
	counts := filterCount(series)
	_, err = f.WriteString(fmt.Sprintf("avg COUNT: %d\n", counts.CalcAverage()))
	if err != nil {
		return err
	}

	return nil
}

func filterInc(report ResponseSeries) ResponseSeries {
	var incs ResponseSeries
	for _, item := range report {
		if item.Operation == OperationInc {
			incs = append(incs, item)
		}
	}
	return incs
}
func filterCount(report ResponseSeries) ResponseSeries {
	var counts ResponseSeries
	for _, item := range report {
		if item.Operation == OperationCount {
			counts = append(counts, item)
		}
	}
	return counts
}
