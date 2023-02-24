package main

import (
	"fmt"
	"time"
)

/*
	|-----------|---------------|-----------------------|
	| Operation | Channel State |      Result           |
	|-----------|---------------|-----------------------|
	|  send     |   open        | block until a receive |
	|  receive  |   open        | block until a send    |
	|  close    |   open        | closed                |
	|  send     |   closed      | panic                 |
	|  receive  |   closed      | zero value wo block   |
	|  close    |   closed      | panic                 |
	|  send     |   nil         | block forever         |
	|  receive  |   nil         | block forever         |
	|  close    |   nil         | panic                 |
	|-----------|---------------|-----------------------|
*/

func main() {
	go fmt.Println("goroutine")
	fmt.Println("main")

	for i := 0; i < 3; i++ {
		go func(num int) {
			fmt.Println(num)
		}(i)
	}

	time.Sleep(10 * time.Millisecond)
	shadowExample()

	// ------------------------------------------------------------

	ch := make(chan string)
	go func() {
		ch <- "hi"
	}()

	msg := <-ch
	fmt.Println(msg)

	go func() {
		for i := 0; i < 3; i++ {
			msg = fmt.Sprintf("message #%d", i)
			ch <- msg
		}
		close(ch) // without will be deadlock all goroutines are sleep
	}()

	for m := range ch {
		fmt.Println("got:", m)
	}

	msg, ok := <-ch
	fmt.Printf("closed:%#v, isOK: %v\n", msg, ok) // zero values, false

	// ------------------------------------------------------------
	values := []int{15, 8, 42, 16, 4, 23}
	fmt.Println(sleepSort(values))
	// ------------------------------------------------------------
}

func shadowExample() {
	n := 7
	{
		n := 2
		fmt.Println("inner", n)
	}
	fmt.Println("outer", n)
}

func sleepSort(values []int) []int {
	ch := make(chan int)
	for _, v := range values {
		v := v
		go func() {
			time.Sleep(time.Duration(v) * time.Millisecond)
			ch <- v
		}()
	}

	var out []int
	for range values {
		n := <-ch
		out = append(out, n)
	}

	return out
}
