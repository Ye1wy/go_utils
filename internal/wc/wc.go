package wc

import (
	"bytes"
	"io"
)

func LienCounter(reader io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSeparetor := []byte{'\n'}

	for {
		readed, err := reader.Read(buf)
		count += bytes.Count(buf[:readed], lineSeparetor)

		switch {
		case err != nil:
			return count, err
		case err == io.EOF:
			return count, nil
		}
	}
}
