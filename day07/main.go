package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	b := bufio.NewReaderSize(f, 256)

	nSupportTLS := 0
	nSupportSSL := 0

	for {
		line, isPrefix, err := b.ReadLine()
		if err == io.EOF || line == nil {
			break
		}

		if isPrefix || err != nil {
			panic("buffer too small or unexpected error")
		}

		supportsTLS, err := supportsTLS(string(line))
		if err != nil {
			panic(err)
		}

		if supportsTLS {
			nSupportTLS++
		}

		supportsSSL, err := supportsSSL(string(line))
		if err != nil {
			panic(err)
		}

		if supportsSSL {
			nSupportSSL++
		}
	}

	fmt.Printf("found %d IPs supporting TLS\n", nSupportTLS)

	fmt.Printf("found %d IPs supporting SSL\n", nSupportSSL)
}
