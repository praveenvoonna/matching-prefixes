package main

import (
	"bufio"
	"fmt"
	"matching_prefixes/prefixes"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var maxGoroutines int = 1000

	// Load prefixes using a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := prefixes.LoadPrefixes("prefixes.txt", maxGoroutines)
		if err != nil {
			fmt.Println("Error loading prefixes from prefixes.txt file")
		}
	}()

	// Read input string from terminal
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the input string: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input string from terminal")
	}

	input = input[:len(input)-1]

	// Use another goroutine to find the longest matching prefix
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Input String:", input)
		longestPrefix := prefixes.FindLongestPrefix(input)
		fmt.Println("Longest Matching Prefix:", longestPrefix)
	}()

	wg.Wait() // Wait for goroutines to finish
}
