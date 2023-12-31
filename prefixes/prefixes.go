package prefixes

import (
	"bufio"
	"os"
	"sync"
)

var prefixMap map[string]bool
var mutex = &sync.Mutex{}

func LoadPrefixes(filename string, maxGoroutines int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	prefixMap = make(map[string]bool)
	scanner := bufio.NewScanner(file)

	semaphore := make(chan struct{}, maxGoroutines)
	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		prefix := scanner.Text()
		semaphore <- struct{}{}
		go func(p string) {
			defer func() {
				wg.Done()
				<-semaphore
			}()
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
	for i := len(input); i > 0; i-- {
		substr := input[:i]
		mutex.Lock()
		_, ok := prefixMap[substr]
		mutex.Unlock()
		if ok {
			longestPrefix = substr
			return longestPrefix
		}
	}
	return "sorry! no records found :("
}
