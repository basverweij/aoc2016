package main

import (
	"fmt"
	"regexp"
	"strings"
)

func supportsTLS(ip string) (bool, error) {
	ips, hypernets, err := parseIP(ip)
	if err != nil {
		return false, err
	}

	// first check if at least one of the IPs contains an ABBA
	abbaInIP := false
	for _, ip := range ips {
		if containsABBA(ip) {
			abbaInIP = true
			break
		}
	}

	// if none of the IPs contains an ABBA were done
	if !abbaInIP {
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

func parseIP(ip string) (ips []string, hypernets []string, err error) {
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
			// ip
			ips = append(ips, s)
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
