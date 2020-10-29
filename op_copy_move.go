package osx

import "os"

func copyFile(source, destination string) error {
	batcher, err := NewFileBatcher(source, defaultBatchSize)
	data, hasNext := batcher.Next()
	for hasNext {
		err = overWriteFile(destination, data)
		if err != nil {
			return err
		}
		data, hasNext = batcher.Next()
	}
	return nil
}

func moveFile(source, destination string) error {
	err := copyFile(source, destination)
	if err != nil {
		return err
	}
	err = deleteFile(source)
	if err != nil {
		return err
	}
	return nil
}

func copyFolder(source, destination string) error {
	return walk(source, destination, copyFile)
}

func moveFolder(source, destination string) error {
	err := walk(source, destination, moveFile)
	if err != nil {
		return err
	}
	err = deleteFolder(source)
	if err != nil {
		return err
	}
	return nil
}

func mkdir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
