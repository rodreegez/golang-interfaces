// An example of wrapping interfaces used by Andrew Gerrand in his tallk
// "Go For Gophers", given at Gophercon 2014 as the closing day keynote
// http://www.confreaks.com/videos/3443-gophercon2014-closing-day-keynote

package main

import (
	"fmt"
	"io"
	"log"
)

// ByteReader is a byte
type ByteReader byte

// Make byte an io.Reader that emits a stream of it's byte value into a slice
func (b ByteReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = byte(b)
	}

	return len(buf), nil
}

// LogReader embedds an io.Reader and wraps it with logging functionality
type LogReader struct {
	io.Reader
}

func (r LogReader) Read(b []byte) (int, error) {
	n, err := r.Reader.Read(b)
	log.Printf("read %d bytes, error: %w", n, err)
	return n, err
}

func main() {
	// Wrap a ByteReader for the byte 'a' in a LogReader
	r := LogReader{ByteReader('a')}
	// Create an empty slice of 10 bytes
	b := make([]byte, 10)
	// Call Read on our LogReader with our slice
	r.Read(b)
	// Print the result, which is our empty slice filled with the ByteReader byte
	// Also loggs thanks to the LogReader
	fmt.Printf("b: %q", b)
}
