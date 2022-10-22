package report

import (
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
