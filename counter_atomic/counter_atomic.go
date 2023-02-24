package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	const nG = 10
	var count int32 = 0

	var wg sync.WaitGroup

	wg.Add(nG)

	for i := 0; i < nG; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				atomic.AddInt32(&count, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println(count)
}
