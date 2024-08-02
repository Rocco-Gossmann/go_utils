package go_utils

import "io"

type t_CopyTargetWithProgress struct {
	m_file    io.Writer
	m_channel *chan int
}

func (c t_CopyTargetWithProgress) Write(data []byte) (int, error) {
	written, err := c.m_file.Write(data)

	if written > 0 && err == nil {
		*(c.m_channel) <- written
	}

	return written, err
}

func CopyWithProgress(reader io.Reader, writer io.Writer, onProgress func(bytesCopied int)) (copied int64, err error) {
	var progressChannel chan int = make(chan int)
	var progress int
	var totalProgress int

	outFile := t_CopyTargetWithProgress{
		m_file:    writer,
		m_channel: &progressChannel,
	}

	go func() {
		copied, err = io.Copy(outFile, reader)
		progressChannel <- -1
	}()

	for {
		progress = <-progressChannel

		if progress == -1 {
			break
		} else {
			totalProgress += progress
			onProgress(totalProgress)
		}
	}

	return

}
