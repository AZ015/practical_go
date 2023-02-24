package main

import (
	"fmt"
	"sync"
)

func main() {
	const nG = 10
	count := 0

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(nG)

	for i := 0; i < nG; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
