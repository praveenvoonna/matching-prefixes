package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findLongestPrefix(input string, prefixes []string) string {
	longestPrefix := ""
	for _, prefix := range prefixes {
		if strings.HasPrefix(input, prefix) && len(prefix) > len(longestPrefix) {
			longestPrefix = prefix
		}
	}
	return longestPrefix
}

func main() {
	file, err := os.Open("sample_prefixes_(3).txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var prefixes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prefixes = append(prefixes, scanner.Text())
	}

	input := "applepie"
	fmt.Println("Input String:", input)
	longestPrefix := findLongestPrefix(input, prefixes)
	fmt.Println("Longest Matching Prefix:", longestPrefix)
}
