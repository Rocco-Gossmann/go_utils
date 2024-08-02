package testdata

import "testing"

type t_TestWriter struct {
	m_t *testing.T
}

// Creates a writer, that simulates writing 1 byte per Write opperation
func NewTestWriter(t *testing.T) t_TestWriter {
	return t_TestWriter{m_t: t}
}

func (w t_TestWriter) Write([]byte) (written int, err error) {
	return 1, nil
}
