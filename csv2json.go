package csv2json

import (
	"fmt"
)

func ReadFile(filePath string) string { // This will actually return a json array
	fmt.Printf("Got: %s\n", filePath)
	return filePath
}
