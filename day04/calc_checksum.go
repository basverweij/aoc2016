package main

import "sort"

type runeFreq struct {
	Rune rune
	Freq int
}

type runeFreqs []*runeFreq

func (fs runeFreqs) Len() int      { return len(fs) }
func (fs runeFreqs) Swap(i, j int) { fs[i], fs[j] = fs[j], fs[i] }
func (fs runeFreqs) Less(i, j int) bool {
	if fs[i].Freq == fs[j].Freq {
		// if frequencies are equal: compare alphabetically
		return fs[i].Rune < fs[j].Rune
	}

	return fs[i].Freq > fs[j].Freq
}

func calcChecksum(roomID string) string {
	freqs := make(runeFreqs, 26)
	for i := 0; i < 26; i++ {
		freqs[i] = &runeFreq{Rune: rune('a' + i)}
	}

	for _, r := range roomID {
		if r == '-' {
			continue
		}

		freqs[r-'a'].Freq++
	}

	sort.Sort(freqs)

	return string(freqs[0].Rune) + string(freqs[1].Rune) + string(freqs[2].Rune) + string(freqs[3].Rune) + string(freqs[4].Rune)
}
