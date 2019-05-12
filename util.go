package main

import (
	"fmt"
	"strconv"
	"strings"
)

func printResult(word []byte, count int) {
	fmt.Println(pad(strconv.Itoa(count), 4), string(word))
}

func pad(s string, n int) string {
	return padChar(s, n, ' ')
}

func padChar(s string, n int, r rune) string {
	if len(s) > n {
		return s
	}

	return strings.Repeat(string(r), n-len(s)) + s
}
