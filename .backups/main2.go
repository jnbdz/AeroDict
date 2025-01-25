package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the CSV file
	csvFile, err := os.Open("config/test.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {

		}
	}(csvFile)

	// Create a new CSV reader
	reader := csv.NewReader(csvFile)
	reader.Comma = ','       // Define the delimiter
	reader.LazyQuotes = true // Allow lazy quotes

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Process records
	for _, record := range records {
		for i, field := range record {
			// Replace escaped quotes with a single quote
			field = strings.ReplaceAll(field, `""`, `"`)
			// Remove newlines from the field
			field = strings.ReplaceAll(field, "\n", " ")
			// Optional: Replace any specific processing you need here
			record[i] = field
		}
		// Print the processed record
		fmt.Println(strings.Join(record, ", "))
	}
}
