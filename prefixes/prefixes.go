package prefixes

import (
	"bufio"
	"os"
)

var prefixMap map[string]bool

func LoadPrefixes(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	prefixMap = make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prefix := scanner.Text()
		prefixMap[prefix] = true
	}
	return scanner.Err()
}

func FindLongestPrefix(input string) string {
	longestPrefix := ""
	for i := 1; i <= len(input); i++ {
		substr := input[:i]
		if _, ok := prefixMap[substr]; ok {
			longestPrefix = substr
		}
	}
	return longestPrefix
}
