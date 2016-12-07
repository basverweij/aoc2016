package main

import (
	"fmt"
	"regexp"
	"strings"
)

func supportsTLS(ip string) (bool, error) {
	supernets, hypernets, err := parseIP(ip)
	if err != nil {
		return false, err
	}

	// first check if at least one of the IPs contains an ABBA
	abbaInSupernet := false
	for _, supernet := range supernets {
		if containsABBA(supernet) {
			abbaInSupernet = true
			break
		}
	}

	// if none of the IPs contains an ABBA were done
	if !abbaInSupernet {
		return false, nil
	}

	// check whether none of the hypernets contains an ABBA
	for _, hypernet := range hypernets {
		if containsABBA(hypernet) {
			return false, nil
		}
	}

	// none of the hypernets matched so this IP supports TLS
	return true, nil
}

var ipPattern = regexp.MustCompile(`[a-z]+|\[[a-z]+\]`)

func parseIP(ip string) (supernets []string, hypernets []string, err error) {
	m := ipPattern.FindAllString(ip, -1)
	if m == nil || len(m) < 1 {
		err = fmt.Errorf("invalid ip: %s", ip)
		return
	}

	for _, s := range m {
		if strings.HasPrefix(s, "[") {
			// hypernet
			hypernets = append(hypernets, s)
		} else {
			// supernet
			supernets = append(supernets, s)
		}
	}

	return
}

func containsABBA(s string) bool {
	for i := 0; i < len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}

	return false
}
