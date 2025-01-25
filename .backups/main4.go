package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dataharmonizer",
	Short: "DataHarmonizer harmonizes and presents CSV data in structured formats.",
}

// viewCmd represents the command to view CSV data
var viewCmd = &cobra.Command{
	// Previous setup remains the same
	Run: func(cmd *cobra.Command, args []string) {
		csvFile := args[0]
		format, _ := cmd.Flags().GetString("format")
		columns, _ := cmd.Flags().GetStringSlice("columns")

		switch format {
		case "table":
			viewAsTable(csvFile, columns)
		case "json":
			viewAsJSON(csvFile, columns)
		case "simple":
			viewAsSimple(csvFile, columns) // Implement this similarly to the initial example
		default:
			fmt.Println("Unsupported format. Available formats: 'simple', 'table', or 'json'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
	// Define the format flag once with all options included in the description
	viewCmd.Flags().String("format", "table", "Format to view the CSV data: simple, table, or json")
}

func viewAsTable(csvFile string, selectedColumns []string) {
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	if len(selectedColumns) > 0 {
		table.SetHeader(selectedColumns)
	} else if len(records) > 0 {
		table.SetHeader(records[0])
		records = records[1:]
	}

	for _, record := range records {
		if len(selectedColumns) > 0 {
			var filteredRecord []string
			for _, col := range selectedColumns {
				for i, header := range records[0] {
					if header == col {
						filteredRecord = append(filteredRecord, record[i])
						break
					}
				}
			}
			table.Append(filteredRecord)
		} else {
			table.Append(record)
		}
	}
	table.Render() // Send output
}

func viewAsJSON(csvFile string, selectedColumns []string) {
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	var jsonRecords []map[string]string

	headers := records[0]
	for _, record := range records[1:] {
		recordMap := make(map[string]string)
		for i, value := range record {
			header := headers[i]
			if len(selectedColumns) == 0 || contains(selectedColumns, header) {
				recordMap[header] = value
			}
		}
		jsonRecords = append(jsonRecords, recordMap)
	}

	jsonData, err := json.MarshalIndent(jsonRecords, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(jsonData))
}

func contains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

func viewAsSimple(csvFile string, selectedColumns []string) {
	// Open the CSV file
	file, err := os.Open(csvFile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	// Read CSV records
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV records:", err)
		return
	}

	// Determine which columns to display
	var indices []int
	if len(selectedColumns) > 0 {
		header := records[0]
		for _, col := range selectedColumns {
			for i, h := range header {
				if h == col {
					indices = append(indices, i)
					break
				}
			}
		}
	} else {
		for i := range records[0] {
			indices = append(indices, i)
		}
	}

	// Display records
	for _, record := range records {
		for _, index := range indices {
			fmt.Printf("%s\t", record[index])
		}
		fmt.Println()
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
