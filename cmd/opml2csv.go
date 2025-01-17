package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/Kharec/opml2csv/internal/flags"
	"github.com/Kharec/opml2csv/internal/utils"
)

func main() {
	showHelp, showVersion, inputFile, outputFile := flags.Parse()
	flags.HandleHelp(showHelp)
	flags.HandleVersion(showVersion)
	flags.ValidateInputFile(inputFile)
	outputFile = flags.SetDefaultOutputFile(outputFile)

	opmlInput, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer opmlInput.Close()
	csvOutput, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer csvOutput.Close()

	csvWriter := csv.NewWriter(csvOutput)
	defer csvWriter.Flush()

	header := utils.BuildCSVHeader()
	if err = csvWriter.Write(header); err != nil {
		fmt.Printf("Error writing CSV header: %v\n", err)
		return
	}

	var outLine *utils.Outline
	scanner := bufio.NewScanner(opmlInput)

	for scanner.Scan() {
		line := scanner.Text()
		if !utils.IsStructuredLine(line) {
			continue
		}
		outLine, err = utils.ParseOpmlLine(line)
		if err != nil {
			fmt.Printf("Failed to parse line: %v\n", err)
			continue
		}
		csvLine := utils.BuildCSVLine(outLine)
		if err = csvWriter.Write(csvLine); err != nil {
			fmt.Printf("Error writing to CSV: %v\n", err)
			return
		}
	}

	if err = scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	fmt.Printf("CSV written to %s successfully.\n", csvOutput.Name())
}
