package main

import (
	"fmt"
	"matching_prefixes/prefixes"
)

func main() {
	prefixes.LoadPrefixes("prefixes.txt")
	input := "mDHjp"
	fmt.Println("Input String:", input)
	longestPrefix := prefixes.FindLongestPrefix(input)
	fmt.Println("Longest Matching Prefix:", longestPrefix)
}
