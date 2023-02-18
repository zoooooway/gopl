package counter

import (
	"bufio"
	"bytes"
)

type LineCounter int

const (
	ONE_LINE LineCounter = 1
	ONE_WORD WordCount   = 1
)

type WordCount int

func (c *LineCounter) write(p []byte) (n int, err error) {
	reader := bytes.NewReader(p)
	scan := bufio.NewScanner(reader)
	scan.Split(bufio.ScanLines)
	for scan.Scan() {
		*c += ONE_LINE
		n += 1
	}
	return n, nil
}

func (c *WordCount) write(p []byte) (n int, err error) {
	reader := bytes.NewReader(p)
	scan := bufio.NewScanner(reader)
	scan.Split(bufio.ScanWords)
	for scan.Scan() {
		*c += ONE_WORD
		n += 1
	}
	return n, nil
}
