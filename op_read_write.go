package osx

import (
	"io/ioutil"
	"os"

	"github.com/ecoshub/byteconv"
)

// ReadFile main's main file read function
func ReadFile(dir string) ([]byte, error) {
	buff, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

func readFileString(dir string) (string, error) {
	buff, err := ioutil.ReadFile(dir)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

//ReadAt read with specific offset of file with a length
func ReadAt(dir string, offset int64, length int64) ([]byte, int, error) {
	f, err := os.Open(dir)
	defer f.Close()
	if err != nil {
		return nil, 0, err
	}
	data := make([]byte, length)
	_, err = f.Seek(offset, 0)
	if err != nil {
		return nil, 0, err
	}
	var n int
	n, err = f.Read(data)
	if err != nil {
		return nil, 0, err
	}
	return data, n, nil
}

//WriteAt read with specific offset of file with a length
func WriteAt(dir string, offset int64, data []byte) (int, error) {
	f, err := os.OpenFile(dir, os.O_WRONLY, os.ModeAppend)
	defer f.Close()
	if err != nil {
		return 0, err
	}
	var n int
	n, err = f.WriteAt(data, offset)
	if err != nil {
		return 0, err
	}
	return n, nil
}

// WriteFileString main file write function
// it appends to end if file exists
// if its not exists it creates and writes
func WriteFileString(dir string, buff string) error {
	b, err := byteconv.ToBytes(buff)
	if err != nil {
		return err
	}
	return WriteFile(dir, b)
}

// WriteFile main file write function
// it appends to end if file exists
// if its not exists it creates and writes
func WriteFile(dir string, buff []byte) error {
	if isDir(dir) {
		return createError("writeFile", errStringDir, dir)
	}
	if isExist(dir) {
		// apppend
		appendFile(dir, buff)
	} else {
		folderDir, _, err := splitDirectory(dir)
		if err != nil {
			return err
		}
		err = os.MkdirAll(folderDir, os.ModePerm)
		if err != nil {
			return err
		}
		// create
		err = overWriteFile(dir, buff)
		if err != nil {
			return err
		}
	}
	return nil
}

// main write function
func overWriteFile(dir string, buffer []byte) error {
	if isDir(dir) {
		return createError("overWriteFile", errStringDir, dir)
	}
	err := ioutil.WriteFile(dir, buffer, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// main append function
func appendFile(dir string, buff []byte) error {
	if isDir(dir) {
		return createError("appendFile", errStringDir, dir)
	}
	f, err := os.OpenFile(dir, os.O_APPEND, os.ModeAppend)
	defer f.Close()
	if err != nil {
		return err
	}
	_, err = f.Write(buff)
	if err != nil {
		return err
	}
	return nil
}
