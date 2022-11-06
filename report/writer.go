package report

import (
	"fmt"
	"os"
)

// WriteSeriesToFile writes full report to file.
// TODO change param to Report Object with all info and write to file.
func WriteSeriesToFile(series ResponseSeries, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(series.String())
	if err != nil {
		return err
	}

	return nil
}

func WriteStatsToFile(report *StatisticReport, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(report.String())
	if err != nil {
		return err
	}

	return nil
}

const csvSuffix = ".csv"

// CsvWriter writes all experiments of one operation to csv file.
type CsvWriter struct {
	file *os.File
}

func NewCsvWriter(filename string) (*CsvWriter, error) {
	file, err := os.Create(filename + csvSuffix)
	if err != nil {
		return nil, err
	}
	_, err = file.WriteString("Counts,Incs,Mean,P25,Median,P75\n")
	if err != nil {
		return nil, err
	}

	return &CsvWriter{
		file: file,
	}, nil
}

func (csv *CsvWriter) WriteStatsToCsv(report *StatisticReport, countsCount int, incsCount int) error {
	_, err := csv.file.WriteString(fmt.Sprintf("%d,%d,%d,%d,%d,%d\n", countsCount, incsCount, report.Mean.Microseconds(), report.P25.Microseconds(), report.Median.Microseconds(), report.P75.Microseconds()))
	if err != nil {
		return err
	}

	return nil
}
