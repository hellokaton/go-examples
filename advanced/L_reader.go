package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("size = %v, err = %v, byte = %v\n", n, err, b)
		fmt.Printf("value = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

