package command

import "fmt"

func Usage() {
	fmt.Println(
		`

		Usage:

		Ggrep <searchString> ( <searchDir> | . ) [-opt]
			
		Arguments:
		
			<searchString>	  The desired text you want to search for
		
			<searchDir>   	  The directory in which you'd like to search. Use '.' to search in the current directory
		
		Options:
			
			-h 			  Show hidden folders and files

        `,
	)
}
