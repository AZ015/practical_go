package main

import "fmt"

func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}

	return max
}

func maxFloats(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}

	return max
}

type Numbers interface {
	int | float64
}

func maxGenerics[T Numbers](nums []T) T {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	// ------------------------------------------------------------

	fmt.Println("-------------------- TYPE CASTING ---------------------------")
	var y any // don't use any!!!

	y = 7
	fmt.Println(y)

	y = "hello"
	fmt.Println(y) // hello

	s := y.(string) // type assertion
	fmt.Println(s)  // hello

	//n := y.(int) // type assertion - panic string not int
	//fmt.Println(n)

	n2, ok := y.(int)   // type assertion - panic string not int
	fmt.Println(n2, ok) // 0, false

	switch y.(type) { // switch type
	case int:
		fmt.Println("y is int")
	case string:
		fmt.Println("y is string")
	default:
		fmt.Printf("y is unknown type: %T", y)
	}

	// ------------------------------------------------------------

	fmt.Println("-------------------- GENERICS ---------------------------")
	fmt.Println(maxInts([]int{1, 2, 3, 4}))
	fmt.Println(maxFloats([]float64{1.4, 1.45, 3.123, 4.456}))
	fmt.Println(maxGenerics([]int{1, 2, 3, 4}))
	fmt.Println(maxGenerics([]float64{1.4, 1.45, 3.123, 4.456}))
}
