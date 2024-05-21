[![GoDoc](https://godoc.org/github.com/gomarkdown/markdown?status.svg)](https://pkg.go.dev/github.com/sieep-coding/grep-in-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/sieep-coding/grep-in-go)](https://goreportcard.com/report/github.com/sieep-coding/grep-in-go)
![Go version](https://img.shields.io/github/go-mod/go-version/sieep-coding/grep-in-go)


## Installation

Inside the source directory, run the following command:

```sh
source install.sh zsh 
# or
source install.sh bash
```
It will build the binary file and put it in a bin folder on your `$HOME` dir.

---

**Note**: If you have already run the command and your *.rc file is already modified, you can run the command again but without the zsh/bash part in order to not have the function repeated in it. 


## Usage

```bash
❯ Ggrep deter .      

./test/br/col/sk/cool_essay.txt
7:'which makes the output <deter>ministic but means that for'
```

```
Usage:

Ggrep <searchString> ( <searchDir> | . ) [-opt]
	
Arguments:

	<searchString>	  The desired text you want to search for

	<searchDir>   	  The directory in which you'd like to search. Use '.' to search in the current directory

Options:
	
	-h 			  Show hidden folders and files

```
