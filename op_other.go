package osx

import "os"

func rename(source string, destination string) error {
	err := os.Rename(source, destination)
	if err != nil {
		return err
	}
	return err
}
