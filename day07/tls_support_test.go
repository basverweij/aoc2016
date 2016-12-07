package main

import "testing"
import "strings"

func testParseIP(t *testing.T, ip string, expectedIPs, expectedHypernets []string) {
	actualIPs, actualHypernets, err := parseIP(ip)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}

	if strings.Join(actualIPs, "|") != strings.Join(expectedIPs, "|") {
		t.Errorf("expected IPs %v for '%s', but got %v", expectedIPs, ip, actualIPs)
	}

	if strings.Join(actualHypernets, "|") != strings.Join(expectedHypernets, "|") {
		t.Errorf("expected hypernets %v for '%s', but got %v", expectedHypernets, ip, actualHypernets)
	}
}

func TestParseIP(t *testing.T) {
	testParseIP(t, "abba[mnop]qrst", []string{"abba", "qrst"}, []string{"[mnop]"})

	testParseIP(t,
		"vwwekuavrftztxb[aywyoempmajrdkxpsc]eibnjbszsfsapujqn[rxpcsihuzszefcdzl]gsahdvozzgxjhontxk",
		[]string{"vwwekuavrftztxb", "eibnjbszsfsapujqn", "gsahdvozzgxjhontxk"},
		[]string{"[aywyoempmajrdkxpsc]", "[rxpcsihuzszefcdzl]"})
}

func testContainsABBA(t *testing.T, s string, expectedResult bool) {
	actualResult := containsABBA(s)
	if actualResult != expectedResult {
		t.Errorf("expected result %v for '%s', but got %v", expectedResult, s, actualResult)
	}
}

func TestContainsABBA(t *testing.T) {
	testContainsABBA(t, "abba", true)
	testContainsABBA(t, "mnop", false)
	testContainsABBA(t, "qrst", false)

	testContainsABBA(t, "abcd", false)
	testContainsABBA(t, "bddb", true)
	testContainsABBA(t, "xyyx", true)

	testContainsABBA(t, "ioxxoj", true)
	testContainsABBA(t, "asdfgh", false)
	testContainsABBA(t, "zxcvbn", false)
}

func testSupportsTLS(t *testing.T, ip string, expectedResult bool) {
	actualResult, err := supportsTLS(ip)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}

	if actualResult != expectedResult {
		t.Errorf("expected result %v for '%s', but got %v", expectedResult, ip, actualResult)
	}
}

func TestSupportsTLS(t *testing.T) {
	testSupportsTLS(t, "abba[mnop]qrst", true)
	testSupportsTLS(t, "abcd[bddb]xyyx", false)
	testSupportsTLS(t, "aaaa[qwer]tyui", false)
	testSupportsTLS(t, "ioxxoj[asdfgh]zxcvbn", true)
}
