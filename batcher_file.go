package osx

import "os"

const (
	kb               int64 = 1024
	mb               int64 = 1024 * 1024
	defaultBatchSize int64 = 10 * mb
)

// FileBatcher main batch object
type FileBatcher struct {
	dir       string
	file      *os.File
	batchSize int64
	offset    int64
	size      int64
	closed    bool
}

// NewFileBatcher is a file iterator
func NewFileBatcher(dir string, batchSize int64) (*FileBatcher, error) {
	if isDir(dir) {
		return nil, createError("NewFileBatcher", errStringDir, dir)
	}
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	b := &FileBatcher{dir: dir, file: file, batchSize: batchSize, size: getFileSize(dir)}
	return b, nil
}

// Next get next batch as byte array
func (b *FileBatcher) Next() ([]byte, bool) {
	if b.closed {
		return nil, false
	}
	data := make([]byte, b.batchSize)
	_, err := b.file.Seek(b.offset, 0)
	if err != nil {
		return nil, false
	}
	var n int
	n, err = b.file.Read(data)
	if err != nil {
		return nil, false
	}
	b.offset += int64(n)
	if b.offset >= b.size || n == 0 {
		b.closed = true
	}
	return data[:n], true

}

// Close close batcher
func (b *FileBatcher) Close() {
	b.closed = true
	b.file.Close()
}

// Refresh refresh batcher
func (b *FileBatcher) Refresh() error {
	b.closed = false
	b.offset = 0
	file, err := os.Open(b.dir)
	if err != nil {
		return err
	}
	b.file = file
	return nil
}
