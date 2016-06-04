/*
	Package csv2json is a simple library that allows csv content to be parsed
	into an array of json objects
*/
package csv2json

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
)

// ReadFile reads a csv file and returns the json
// as a string
func ReadFile(filePath string) ([]byte, error) {
	fileReader, err := os.Open(filePath)

	if err != nil {
		return make([]byte, 0), err
	}

	return Read(fileReader)
}

// Read processes content from a io.Reader and treats
// each line as a row in a csv file
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

	return json.Marshal(entries)
}

// ReadRow turns the next row in a csv.Reader into a map[string]string
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
