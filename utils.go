package osx

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"strings"
)

func splitDirectory(dir string) (string, string, error) {
	if dir == "" {
		return "", "", createError("splitDirectory", errStringEmptyDir, dir)
	}
	isdir := isDir(dir)
	if isdir {
		return dir, "", nil
	}
	dir = strings.TrimSpace(dir)
	tokens := strings.Split(dir, seperator)
	tokenCount := len(tokens)
	cleanTokens := make([]string, 0, tokenCount-1)
	for _, t := range tokens {
		if t != "" {
			cleanTokens = append(cleanTokens, t)
		}
	}
	tokenCount = len(cleanTokens)
	if tokenCount == 0 {
		return "", "", createError("splitDirectory", errStringEmptyDir, dir)
	}
	if tokenCount == 1 {
		return seperator, cleanTokens[0], nil
	}
	dir = strings.Join(cleanTokens[:tokenCount-1], seperator)
	name := cleanTokens[tokenCount-1]
	return dir, name, nil
}

func isDir(dir string) bool {
	fi, err := os.Stat(dir)
	if err != nil {
		return false
	}
	if fi.Mode().IsDir() {
		return true
	}
	return false
}

func isExist(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func newLine() string {
	goos := runtime.GOOS
	if goos == "windows" {
		return "\r\n"
	}
	return "\n"
}

func getHome() string {
	myself, _ := user.Current()
	return myself.HomeDir
}

func isLinux() bool {
	return runtime.GOOS == "linux"
}

func createError(funcname, errString, dir string) error {
	return fmt.Errorf("func: %v dir: %v %v", funcname, dir, errString)
}

func getFileSize(dir string) int64 {
	info, err := os.Stat(dir)
	if err != nil {
		return int64(0)
	}
	return info.Size()
}
