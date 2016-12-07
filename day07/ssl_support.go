package main

import "strings"

func supportsSSL(ip string) (bool, error) {
	supernets, hypernets, err := parseIP(ip)
	if err != nil {
		return false, err
	}

	for _, supernet := range supernets {
		for i := 0; i < len(supernet)-2; i++ {
			if supernet[i] != supernet[i+1] && supernet[i] == supernet[i+2] {
				// found ABA, check hypernets for BAB
				bab := supernet[i+1:i+3] + supernet[i+1:i+2]
				for _, hypernet := range hypernets {
					if strings.Contains(hypernet, bab) {
						return true, nil
					}
				}
			}
		}
	}

	return false, nil
}
