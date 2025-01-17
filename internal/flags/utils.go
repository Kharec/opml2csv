package flags

import (
	"fmt"
	"os"
)

func ValidateInputFile(inputFile string) {
	if inputFile == "" {
		fmt.Println("Error: No input file provided. Use -i or --input to specify the input.")
		os.Exit(1)
	}
}

func SetDefaultOutputFile(outputFile string) string {
	if outputFile == "" {
		return "./export.csv"
	}
	return outputFile
}
