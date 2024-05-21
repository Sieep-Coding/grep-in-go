package command

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/sieep-coding/grep-in-go/locator"
)

func ParseArgs() (Arguments, bool) {
	if len(os.Args) <= 2 {
		Usage()
		os.Exit(0)
	}

	args := os.Args[1:]
	opts := args[2:]

	mainArgs := Arguments{}

	if len(opts) != 0 {
		mainArgs.Options = ParseOpts(opts)
	}

	if args[0] == "usage" {
		Usage()
		os.Exit(0)
		return Arguments{}, false
	}

	if len(args) < 2 {
		Usage()
		os.Exit(0)
		return Arguments{}, false
	}

	mainArgs.SearchString, mainArgs.Directory = args[0], args[1]

	ok := validateDir(mainArgs.Directory)

	if !ok {
		fmt.Printf("%s is not a valid directory\n", mainArgs.Directory)
		return Arguments{}, false
	}

	return mainArgs, true
}

func ParseOpts(opts []string) locator.OptionConfig {
	var options locator.OptionConfig
	for _, opt := range opts {
		switch opt {
		case "-h":
			options.Hidden = true
		case "v":
			options.Verbose = true
		default:
			color.Blue("Unregonized Argument %s", opt)
		}
	}
	return options
}
