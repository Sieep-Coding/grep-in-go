package locator

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type Locator struct {
	BaseDir string
	Options OptionConfig
}

type OptionConfig struct {
	Verbose, Hidden bool
}

type AnalyzedFile struct {
	FilePath  string
	Content   *os.File
	Locations []Location
	OK        bool
}

type Location struct {
	LineNo   string
	Contents string
}

func (f *AnalyzedFile) GetInfo() {
	filepath := color.BlueString(f.FilePath)
	fmt.Printf("\n%s\n", filepath)
	for _, loc := range f.Locations {
		lineNo := color.RedString(fmt.Sprintf("%v", loc.LineNo))
		fmt.Printf("%s:%s\n", lineNo, loc.Contents)
	}
}
