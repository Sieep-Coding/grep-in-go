package main

import (
	command "github.com/sieep-coding/grep-in-go/command"
	"github.com/sieep-coding/grep-in-go/locator"
)

func main() {
	args, ok := command.ParseArgs()

	if !ok {
		return
	}

	locator := locator.NewLocator(args.Directory)
	locator.Options = args.Options

	locator.Dig(args.SearchString)
}
