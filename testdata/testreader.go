package testdata

import (
	"io"
	"testing"
)

type t_TestReader struct {
	m_read *int
	m_t    *testing.T
}

// creates a reader, that responds with 4 bytes "1234"  one byte per Read
func NewTestReader(t *testing.T) t_TestReader {
	var readbytes = 0
	return t_TestReader{
		m_read: &readbytes,
		m_t:    t,
	}
}

func (r t_TestReader) Read(out []byte) (int, error) {
	*r.m_read += 1

	if *r.m_read > 4 {
		return 0, io.EOF
	}

	out[0] = byte(*r.m_read)

	return 1, nil
}
