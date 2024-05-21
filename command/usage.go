package command

import "fmt"

func Usage() {
	fmt.Println(
		`

		Usage:

grip <searchString> ( <searchDir> | . ) [-opt]

Arguments:

        <searchString>    The desired text you want to search for

        <searchDir>       The directory in which you'd like to search. Use '.' to search in the current directory

Options:

        -h                                Search hidden folders and files

        `,
	)
}
