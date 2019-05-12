package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	CountLimit = 20
)

func main() {
	if len(os.Args) == 1 {
		exitWithError(fmt.Sprintf("usage: %s /path/to/file", os.Args[0]))
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		exitWithError(err.Error())
	}

	defer file.Close()

	counter := NewWordsCounter(CountLimit)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		counter.ProcessChunk(scanner.Bytes())
	}

	result := counter.Result()

	for i := 0; i < CountLimit && i < len(result); i++ {
		printResult(result[i].Value, result[i].Count)
	}
}

func exitWithError(message string) {
	fmt.Println(message)
	os.Exit(1)
}
