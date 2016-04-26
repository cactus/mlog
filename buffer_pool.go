package mlog

import (
	"io"
	"sync"
)

type sliceBuffer struct {
	data []byte
}

func (sb *sliceBuffer) AppendIntWidth(i int, wid int) {
	x := 0
	for i >= 10 {
		q := i / 10
		sb.data = append(sb.data, byte('0'+i-q*10))
		i = q
		x++
	}
	sb.data = append(sb.data, byte('0'+i))
	x++

	y := x
	for wid > x {
		sb.data = append(sb.data, '0')
		wid--
		y++
	}

	// twizzle
	for i, j := len(sb.data)-y, len(sb.data)-1; i < j; i, j = i+1, j-1 {
		sb.data[i], sb.data[j] = sb.data[j], sb.data[i]
	}
}

func (sb *sliceBuffer) Write(b []byte) (int, error) {
	sb.data = append(sb.data, b...)
	return len(b), nil
}

func (sb *sliceBuffer) WriteByte(c byte) error {
	sb.data = append(sb.data, c)
	return nil
}

func (sb *sliceBuffer) WriteString(s string) (int, error) {
	sb.data = append(sb.data, s...)
	return len(s), nil
}

func (sb *sliceBuffer) WriteTo(w io.Writer) (int, error) {
	return w.Write(sb.data)
}

func (sb *sliceBuffer) Bytes() []byte {
	return sb.data
}

func (sb *sliceBuffer) String() string {
	return string(sb.data)
}

func (sb *sliceBuffer) Len() int {
	return len(sb.data)
}

func (sb *sliceBuffer) Reset() {
	sb.Truncate(0)
}

func (sb *sliceBuffer) Truncate(i int) {
	sb.data = sb.data[:i]
}

type sliceBufferPool struct {
	*sync.Pool
}

func newSliceBufferPool() *sliceBufferPool {
	return &sliceBufferPool{
		&sync.Pool{New: func() interface{} {
			return &sliceBuffer{make([]byte, 0, 1024)}
		}},
	}
}

func (sp *sliceBufferPool) Get() *sliceBuffer {
	return (sp.Pool.Get()).(*sliceBuffer)
}

func (sp *sliceBufferPool) Put(c *sliceBuffer) {
	c.Reset()
	sp.Pool.Put(c)
}
