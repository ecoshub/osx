package osx

import (
	"bufio"
	"os"
)

// LineReader line by line reader for iteration
type LineReader struct {
	dir    string
	file   *os.File
	scan   *bufio.Scanner
	offset int64
	closed bool
}

// NewLineReader is constructor of LineReader object
func NewLineReader(dir string) (*LineReader, error) {
	if isDir(dir) {
		return nil, createError("NewLineReader", errStringDir, dir)
	}
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	rl := LineReader{dir: dir, file: file, scan: bufio.NewScanner(file)}
	return &rl, nil
}

// Next gives next line
func (r *LineReader) Next() ([]byte, bool) {
	if r.closed {
		return nil, false
	}
	if r.scan.Scan() {
		buffer := r.scan.Bytes()
		r.offset += int64(len(buffer)) + int64(len(newLine()))
		return buffer, true
	}
	r.Close()
	return nil, false
}

// Close close the ReadLine object
func (r *LineReader) Close() {
	r.closed = true
	r.file.Close()
}

// Refresh resets the reader
func (r *LineReader) Refresh() error {
	r.closed = false
	r.offset = 0
	file, err := os.Open(r.dir)
	if err != nil {
		return err
	}
	r.file = file
	r.scan = bufio.NewScanner(file)
	return nil
}
