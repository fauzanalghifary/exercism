package circular

import (
	"errors"
)

type Buffer struct {
	data       []byte
	readIndex  int
	writeIndex int
	size       int
	isFull     bool
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		size: size,
		data: make([]byte, size),
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.isEmpty() {
		return 0, errors.New("empty Buffer")
	}

	val := b.data[b.readIndex]
	b.readIndex = (b.readIndex + 1) % b.size
	b.isFull = false

	return val, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.isBufferFull() {
		return errors.New("cannot write to full buffer")
	}

	b.data[b.writeIndex] = c
	b.writeIndex = (b.writeIndex + 1) % b.size

	if b.writeIndex == b.readIndex {
		b.isFull = true
	}
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.isBufferFull() {
		b.readIndex = (b.readIndex + 1) % b.size
	}

	b.data[b.writeIndex] = c
	b.writeIndex = (b.writeIndex + 1) % b.size
	b.isFull = b.writeIndex == b.readIndex
}

func (b *Buffer) Reset() {
	b.data = make([]byte, b.size)
	b.readIndex = 0
	b.writeIndex = 0
	b.isFull = false
}

func (b *Buffer) isEmpty() bool {
	return !b.isFull && b.readIndex == b.writeIndex
}

func (b *Buffer) isBufferFull() bool {
	return b.isFull
}

//type Buffer struct {
//	data       []byte
//	writeIndex int
//	readIndex  int
//	size       int
//}
//
//func NewBuffer(size int) *Buffer {
//	return &Buffer{
//		size: size,
//		data: make([]byte, 0, size),
//	}
//}
//
//func (b *Buffer) ReadByte() (byte, error) {
//	if len(b.data) == 0 {
//		return 0, errors.New("empty Buffer")
//	}
//	val := b.data[0]
//	b.data = slices.Delete(b.data, 0, 1)
//	b.writeIndex--
//	return val, nil
//}
//
//func (b *Buffer) WriteByte(c byte) error {
//	if len(b.data) == b.size {
//		return errors.New("cannot write to full byte")
//	}
//	b.data = slices.Insert(b.data, b.writeIndex, c)
//	b.writeIndex++
//	return nil
//}
//
//func (b *Buffer) Overwrite(c byte) {
//	if len(b.data) == b.size {
//		b.data = slices.Delete(b.data, 0, 1)
//		b.writeIndex--
//	}
//
//	b.data = slices.Insert(b.data, b.writeIndex, c)
//
//	b.writeIndex++
//}
//
//func (b *Buffer) Reset() {
//	b.data = []byte{}
//	b.writeIndex = 0
//}
