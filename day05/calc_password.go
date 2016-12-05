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
	for found := 0; found < 8; found++ {
		for {
			hash := calcHash(input, i)
			i++

			if strings.HasPrefix(hash, prefix) {
				pos := hash[prefixLen-2]
				if pos >= '0' && pos < '8' {
					pos -= '0'
				} else {
					// not a valid position
					continue
				}

				if pwd[pos] != 0 {
					// position already filled
					continue
				}

				pwd[pos] = hash[prefixLen-1]
				break
			}
		}
	}

	return string(pwd)
}

func calcHash(input string, index int) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s%d", input, index)))
	s := hex.EncodeToString(hash[:useBytes])

	return s[:prefixLen]
}
