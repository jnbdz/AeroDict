package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the CSV file
	csvFile, err := os.Open("config/noaa.csv")
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

	// Find the maximum length of each column
	maxLengths := make([]int, len(records[0]))
	for _, record := range records {
		for i, field := range record {
			field = strings.ReplaceAll(field, `""`, `"`)
			field = strings.ReplaceAll(field, "\n", " ")
			record[i] = field
			if len(field) > maxLengths[i] {
				maxLengths[i] = len(field)
			}
		}
	}

	// Print the records with padding for alignment
	for _, record := range records {
		for i, field := range record {
			fmt.Printf("%-*s", maxLengths[i]+2, field) // +2 for padding
		}
		fmt.Println()
	}
}
