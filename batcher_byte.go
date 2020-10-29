package osx

// ByteBatcher main batch object
type ByteBatcher struct {
	buffer    *[]byte
	batchSize int64
	offset    int64
	size      int64
	closed    bool
}

// NewByteBatcher is a file iterator
func NewByteBatcher(buffer []byte, batchSize int64) (*ByteBatcher, error) {
	b := &ByteBatcher{buffer: &buffer, batchSize: batchSize, size: int64(len(buffer))}
	return b, nil
}

// Next get next batch as byte array
func (b *ByteBatcher) Next() ([]byte, bool) {
	if b.closed {
		return nil, false
	}
	delt := b.size - b.batchSize
	if delt < 0 {
		b.batchSize = b.size
	}
	var buffer []byte = *b.buffer
	data := buffer[b.offset : b.offset+b.batchSize]
	b.offset += b.batchSize
	if b.offset >= b.size {
		b.closed = true
	}
	return data, true
}
