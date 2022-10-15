package report

import (
	"fmt"
	"os"
)

func WriteToFile(report Report, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	for _, item := range report {
		_, err := f.WriteString(fmt.Sprintf("%s: %s %s\n", item.Operation, item.Address, item.ResponseTime))
		if err != nil {
			return err
		}
	}

	_, err = f.WriteString(fmt.Sprint("avg in mcs\n"))
	_, err = f.WriteString(fmt.Sprintf("avg ALL: %d\n", report.CalcAverage()))
	if err != nil {
		return err
	}
	incs := filterInc(report)

	_, err = f.WriteString(fmt.Sprintf("avg INC: %d\n", incs.CalcAverage()))
	if err != nil {
		return err
	}
	counts := filterCount(report)
	_, err = f.WriteString(fmt.Sprintf("avg COUNT: %d\n", counts.CalcAverage()))
	if err != nil {
		return err
	}

	return nil
}

func filterInc(report Report) Report {
	var incs Report
	for _, item := range report {
		if item.Operation == IncOperation {
			incs = append(incs, item)
		}
	}
	return incs
}
func filterCount(report Report) Report {
	var counts Report
	for _, item := range report {
		if item.Operation == CountOperation {
			counts = append(counts, item)
		}
	}
	return counts
}
