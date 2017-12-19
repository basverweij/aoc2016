package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

type hash struct {
	index        int
	h            string
	firstTriplet byte
	quintets     map[byte]struct{}
}

// genHash generates as hash from the salt and index.
// it only returns a non-nil result if the generated hash
// contains a triplet, and/or one or more quintents
func genHash(salt string, index int) *hash {
	b := md5.Sum([]byte(salt + strconv.Itoa(index)))
	s := hex.EncodeToString(b[:])

	firstTriplet := getFirstTriplet(s)
	quintets := getQuintets(s)

	if firstTriplet == 0 && quintets == nil {
		return nil
	}

	h := &hash{
		h:            s,
		index:        index,
		firstTriplet: firstTriplet,
		quintets:     make(map[byte]struct{}, len(quintets)),
	}

	for _, q := range quintets {
		h.quintets[q] = struct{}{}
	}

	return h
}

func getFirstTriplet(s string) byte {
	ss := getSeries(s, 3)
	if len(ss) < 1 {
		return 0
	}

	return ss[0]
}

func getQuintets(s string) []byte {
	return getSeries(s, 5)
}

func getSeries(s string, length int) []byte {
	var ss []byte
	n := len(s)

	d := s[0]
	j := 0
	i := 1
	for ; i < n; i++ {
		if i-j == length {
			ss = append(ss, s[j])
		}

		if s[i] == d {
			continue
		}

		d = s[i]
		j = i
	}

	if i-j == length {
		ss = append(ss, s[j])
	}

	return ss
}
