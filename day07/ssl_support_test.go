package main

import "testing"

func testSupportSSL(t *testing.T, ip string, expectedResult bool) {
	actualResult, err := supportsSSL(ip)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
		return
	}

	if actualResult != expectedResult {
		t.Errorf("expected result %v for IP '%s', but got %v", expectedResult, ip, actualResult)
	}
}

func TestSupportSSL(t *testing.T) {
	testSupportSSL(t, "aba[bab]xyz", true)
	testSupportSSL(t, "xyx[xyx]xyx", false)
	testSupportSSL(t, "aaa[kek]eke", true)
	testSupportSSL(t, "zazbz[bzb]cdb", true)
}
