package prefixes

import (
	"bufio"
	"os"
	"sync"
)

var prefixMap map[string]bool
var mutex = &sync.Mutex{}

func LoadPrefixes(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	prefixMap = make(map[string]bool)
	scanner := bufio.NewScanner(file)

	var wg sync.WaitGroup
	for scanner.Scan() {
		wg.Add(1)
		prefix := scanner.Text()
		go func(p string) {
			defer wg.Done()
			mutex.Lock()
			prefixMap[p] = true
			mutex.Unlock()
		}(prefix)
	}
	wg.Wait()
	return scanner.Err()
}

func FindLongestPrefix(input string) string {
	longestPrefix := ""
	for i := 1; i <= len(input); i++ {
		substr := input[:i]
		mutex.Lock()
		_, ok := prefixMap[substr]
		mutex.Unlock()
		if ok {
			longestPrefix = substr
		}
	}
	return longestPrefix
}
