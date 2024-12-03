package wc

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"unicode/utf8"
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

func WordsCounter(file string) (int, error) {
	count := 0

	reader, err := os.Open(file)
	if err != nil {
		return count, err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)
		count += len(words)
	}

	if err := scanner.Err(); err != nil {
		return count, err
	}

	return count, nil
}

func CharCounter(file string) (int, error) {
	count := 0

	reader, err := os.Open(file)
	if err != nil {
		return count, err
	}
	defer reader.Close()

	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()

		chars := utf8.RuneCountInString(line)
		count += chars
	}

	if err := scanner.Err(); err != nil {
		return count, err
	}

	return count, nil
}
