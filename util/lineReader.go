// Package util provides an example of how create packages
package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// LineReader defines a line reader representation
type LineReader struct {
	bufferReader *bufio.Reader
}

// NewLineReader creeates a new line reader
func NewLineReader() LineReader {
	return LineReader{bufferReader: bufio.NewReader(os.Stdin)}
}

// Read the next line
func (lineReader *LineReader) Read(message string) string {
	if lineReader.bufferReader == nil {
		lineReader.bufferReader = bufio.NewReader(os.Stdin)
	}

	fmt.Println(message)

	line, err := lineReader.bufferReader.ReadString('\n')

	if err != nil {
		log.Panicln("Ops, the program can't read the current line. Cause:", err)
	}

	return strings.Replace(line, "\n", "", 1)
}
