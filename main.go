package main

import (
	"encoding/csv"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "csvviewer [filepath]",
	Short: "View CSV files in different formats",
	Args:  cobra.MinimumNArgs(1),
	Run:   runViewer,
}

var viewMode string
var columns string

func main() {
	rootCmd.PersistentFlags().StringVarP(&viewMode, "view", "v", "default", "View mode: default, json, table")
	rootCmd.PersistentFlags().StringVar(&columns, "columns", "", "Comma-separated list of column indices to display")
	cobra.CheckErr(rootCmd.Execute())
}

func runViewer(cmd *cobra.Command, args []string) {
	filePath := args[0]
	csvFile, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening CSV file: %v\n", err)
		os.Exit(1)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV file: %v\n", err)
		os.Exit(1)
	}

	if columns != "" {
		colIndices, err := parseColumns(columns)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing columns: %v\n", err)
			os.Exit(1)
		}
		records = filterColumns(records, colIndices)
	}

	switch viewMode {
	case "json":
		showJSON(records)
	case "table":
		showTable(records)
	default:
		showDefault(records)
	}

}

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
			// Optional: Replace any specific processing you need here
			record[i] = field
		}
		// Print the processed record
		fmt.Println(strings.Join(record, ", "))
	}
}
