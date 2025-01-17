package flags

import "flag"

func Parse() (showHelp, showVersion bool, inputFile, outputFile string) {
	helpDescription := "Print this help message"
	versionDescription := "Print the version"
	inputDescription := "OPML input file"
	outputDescription := "CSV output file (default: ./export.csv)"

	help := flag.Bool("h", false, helpDescription)
	helpLong := flag.Bool("help", false, helpDescription)
	version := flag.Bool("v", false, versionDescription)
	versionLong := flag.Bool("version", false, versionDescription)
	input := flag.String("i", "", inputDescription)
	inputLong := flag.String("input", "", inputDescription)
	output := flag.String("o", "./export.csv", outputDescription)
	outputLong := flag.String("output", "./export.csv", outputDescription)

	flag.Parse()

	showHelp = *help || *helpLong
	showVersion = *version || *versionLong

	if *inputLong != "" {
		inputFile = *inputLong
	} else {
		inputFile = *input
	}

	if *outputLong != "./export.csv" {
		outputFile = *outputLong
	} else {
		outputFile = *output
	}

	return showHelp, showVersion, inputFile, outputFile
}
