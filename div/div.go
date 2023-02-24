package main

import (
	"fmt"
	"log"
)

func div(a, b int) int {
	return a / b
}

func safeDiv(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide by zero")
	}

	return a / b, nil
}

func safeDivRecover(a, b int) (int, error) {
	defer func() {
		if e := recover(); e != nil {
			log.Println("[recover] error:", e)
		}
	}()

	return a / b, nil
}

func main() {
	//fmt.Println(div(1, 0)) // panic divide by zero
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(1, 1))
	fmt.Println(safeDivRecover(1, 0))
}
