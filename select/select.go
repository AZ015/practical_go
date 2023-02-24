package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- 1
		close(ch1)
	}()

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch2 <- 2
		close(ch2)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)
	defer cancel()

	select {
	case v := <-ch1:
		fmt.Printf("[ch1] <- %d\n", v)
	case v := <-ch2:
		fmt.Printf("[ch2] <- %d\n", v)
	case <-ctx.Done():
		fmt.Println("after 5 milliseconds")

	}
}
