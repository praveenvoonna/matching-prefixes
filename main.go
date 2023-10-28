package main

import (
	"bufio"
	"fmt"
	"matching_prefixes/prefixes"
	"os"
)

func main() {
	err := prefixes.LoadPrefixes("prefixes.txt")
	if err != nil {
		fmt.Println("Error loading prefixes from prefixes.txt file")
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the input string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reader read string from terminal")
	}

	input = input[:len(input)-1]

	fmt.Println("Input String:", input)
	longestPrefix := prefixes.FindLongestPrefix(input)
	fmt.Println("Longest Matching Prefix:", longestPrefix)
}
