package osx

import "os"

func deleteFile(dir string) error {
	return os.Remove(dir)
}

func deleteFolder(dir string) error {
	return os.RemoveAll(dir)
}
