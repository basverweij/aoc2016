package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func calcPassword(input string) string {
	pwd := make([]byte, 8)

	i := 0
	for n := 0; n < 8; n++ {
		for {
			hash := calcHash(input, i)
			i++

			if strings.HasPrefix(hash, prefix) {
				pwd[n] = hash[prefixLen-1]
				break
			}
		}
	}

	return string(pwd)
}

func calcHash(input string, index int) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, index)))

	return hex.EncodeToString(hash[:prefixLen])
}
