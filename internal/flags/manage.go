package flags

import (
	"fmt"
	"os"

	"github.com/Kharec/opml2csv/internal/utils"
)

func HandleHelp(showHelp bool) {
	if showHelp {
		Help()
		os.Exit(0)
	}
}

func HandleVersion(showVersion bool) {
	if showVersion {
		fmt.Printf("Version: %s\n", utils.GetVersion())
		os.Exit(0)
	}
}
