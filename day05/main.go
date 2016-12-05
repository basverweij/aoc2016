package main

import "fmt"

const (
	prefix    string = "00000"
	prefixLen int    = len(prefix) + 1
	pwdLen    int    = 8
)

func main() {
	input := "wtnhxymk"
	pwd := calcPassword(input)

	fmt.Printf("found password '%s' for input '%s'\n", pwd, input)
}
