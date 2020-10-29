package osx

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	seperator string = string(os.PathSeparator)
)

var (
	errStringEmptyDir  string = "null directory string"
	errStringDir       string = "is a folder"
	errStringNotDir    string = "is a file"
	errStringNotExists string = "not exists"
)

func mkdir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func list(dir string) ([]string, error) {
	if !isExist(dir) {
		return nil, createError("list", errStringNotExists, dir)
	}
	if !isDir(dir) {
		return nil, createError("list", errStringNotDir, dir)
	}
	list := make([]string, 0, 8)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		base := filepath.Base(path)
		list = append(list, base)
		if isDir(path) && path != dir {
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func walk(source, destination string, coreFunc func(string, string) error) error {
	if !isExist(source) {
		return createError("walk", errStringNotExists, source)
	}
	if !isDir(source) {
		return createError("walk", errStringNotDir, source)
	}
	if isExist(destination) {
		if !isDir(destination) {
			return createError("walk", errStringNotDir, destination)
		}
	}
	err := filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !isDir(path) {
				destPath := strings.Replace(path, source, destination, -1)
				dirPath, _, _ := splitDirectory(destPath)
				mkdir(dirPath)
				err = coreFunc(path, destPath)
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}
