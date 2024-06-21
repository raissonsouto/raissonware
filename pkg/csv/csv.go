package csv

import (
	"encoding/csv"
	"os"
	"sync"
)

var (
	mu sync.Mutex
)

// LoadCSV loads CSV data from a file and returns a slice of records.
func LoadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

// AppendLine appends a row to the existing CSV file or creates a new CSV file with the content.
func AppendLine(csvPath string, csvRow []string) error {
	mu.Lock()
	defer mu.Unlock()

	var file *os.File
	defer file.Close()

	switch _, err := os.Stat(csvPath); {

	case err == nil:
		file, err = os.OpenFile(csvPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return err
		}

	case os.IsNotExist(err):
		file, err = os.Create(csvPath)
		if err != nil {
			return err
		}

	default:
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err := writer.Write(csvRow)
	return err
}
