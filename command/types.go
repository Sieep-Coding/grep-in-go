package command

import "github.com/sieep-coding/grep-in-go/locator"

type Arguments struct {
	SearchString string
	Directory    string
	Options      locator.OptionConfig
}

type Parser interface {
	ParseArgs() (Arguments, bool)
	ParseOpts(opts []string) locator.OptionConfig
}
