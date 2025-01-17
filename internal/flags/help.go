package flags

import "fmt"

func Help() {
	fmt.Println("Usage:")
	fmt.Println("  -h, --help: Print this help message")
	fmt.Println("  -v, --version: Print the version")
	fmt.Println("  -i, --input: OPML input file")
	fmt.Println("  -o, --output: CSV output file (default: ./export.csv)")
	fmt.Println("")
	fmt.Println("Example:")
	fmt.Println("  opml2csv -h")
	fmt.Println("  opml2csv -i feeds.opml -o feeds.csv")
	fmt.Println("  opml2csv -v")
}
