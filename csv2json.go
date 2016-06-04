package csv2json

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
)

func ReadFile(filePath string) ([]byte, error) {
	fileReader, err := os.Open(filePath)

	if err != nil {
		return make([]byte, 0), err
	}

	json, err := Read(fileReader)

	if err != nil {
		return make([]byte, 0), err
	}

	return json, nil
}

func Read(reader io.Reader) ([]byte, error) {
	csvReader := csv.NewReader(reader)
	entries := make([]map[string]string, 0)

	headers, err := csvReader.Read()
	if err != nil {
		return make([]byte, 0), err
	}

	for {
		row, err := ReadRow(csv.Reader(*csvReader), headers)

		if err == io.EOF {
			break
		}

		if err != nil {
			return make([]byte, 0), err
		}

		entries = append(entries, row)
	}

	json, err := json.Marshal(entries)

	return json, nil
}

func ReadRow(csvReader csv.Reader, headers []string) (map[string]string, error) {
	row, err := csvReader.Read()
	contents := make(map[string]string)

	if err != nil {
		return make(map[string]string, 0), err
	}

	for index, cell := range row {
		contents[headers[index]] = cell
	}

	return contents, nil
}
