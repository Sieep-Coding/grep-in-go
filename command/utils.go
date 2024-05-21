package command

import "os"

// validates that a given directory is valid.
func validateDir(dir string) bool {
	_, err := os.ReadDir(dir)

	return err == nil
}
