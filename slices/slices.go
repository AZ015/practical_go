package main

import (
	"fmt"
	"sort"
)

func main() {
	var s []int
	fmt.Printf("[s] len:%d\n", len(s))
	if s == nil {
		fmt.Println("slice is nil")
	}

	// ------------------------------------------------------------

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("[s2]=%+v\n", s2)

	s3 := s2[1:4]
	fmt.Printf("[s3]=%+v\n", s3)

	//fmt.Println(s2[:100])  // panic: out of range

	s3 = append(s3, 100)
	fmt.Printf("[s3] (append) = %+v\n", s3)
	fmt.Printf("[s2] (after append to s3) = %+v\n", s2) // s2 is changed!
	fmt.Printf("[s2] len=%d, cap=%d\n", len(s2), cap(s2))
	fmt.Printf("[s3] len=%d, cap=%d\n", len(s3), cap(s3))

	// ------------------------------------------------------------

	fmt.Println("-------------------- CONCAT ---------------------------")
	fmt.Println(concat(nil, nil))
	fmt.Println(concat(nil, []string{}))
	fmt.Println(concat([]string{}, nil))
	fmt.Println(concat([]string{"a", "b", "c"}, nil))
	fmt.Println(concat(nil, []string{"a", "b", "c"}))
	fmt.Println(concat([]string{}, []string{"a", "b", "c"}))
	fmt.Println(concat([]string{"a", "b", "c"}, []string{}))
	fmt.Println(concat([]string{"a", "b", "c"}, []string{"d", "e", "f"}))

	// ------------------------------------------------------------

	fmt.Println("-------------------- MEDIAN ---------------------------")
	fmt.Println(median([]float64{}))
	fmt.Println(median([]float64{1.0, 2.0, 3.0}))
	fmt.Println(median([]float64{2.0, 1.0, 3.0, 4.0}))

}

func concat(s1, s2 []string) []string {
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)

	return s
}

func median(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("in slice is empty")
	}

	nums := make([]float64, len(values))
	copy(nums, values)

	sort.Float64s(nums)

	i := len(nums) / 2
	if len(nums)%2 == 1 {
		return nums[i], nil
	}

	return (nums[i-1] + nums[i]) / 2, nil
}
