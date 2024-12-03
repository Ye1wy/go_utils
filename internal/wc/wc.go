package wc

import (
	"bytes"
	"io"
	"os"
)

func LineCounter(file string) (int, error) {
	count := 0

	reader, err := os.Open(file)
	if err != nil {
		return count, err
	}
	defer reader.Close()

	buf := make([]byte, 32*1024)
	lineSeparetor := []byte{'\n'}

	for {
		readed, err := reader.Read(buf)
		count += bytes.Count(buf[:readed], lineSeparetor)

		if err == io.EOF {
			return count, nil
		}

		if err != nil {
			return count, err
		}
	}
}
