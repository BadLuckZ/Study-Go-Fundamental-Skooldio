package main

import "fmt"

func maxInt(nums ...int) int {
	var max int
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	// == Array ==

	var fourNums [4]int
	// [0, 0, 0, 0]

	fourNums[0] = 5
	// [5, 0, 0, 0]

	fourNums[2] = 3
	// [5, 0, 3, 0]

	for i := 0; i < len(fourNums); i++ {
		value := fourNums[i]
		fmt.Printf("value at idx = %d: %d\n", i, value)
	}

	fmt.Println("=======================")

	// == Slices ==
	var nums []int = make([]int, 4)
	// [0, 0, 0, 0]

	nums[0] = 3
	// [3, 0, 0, 0]

	nums = append(nums, 20)
	// [3, 0, 0, 0, 20]

	for i := 0; i < len(nums); i++ {
		value := nums[i]
		fmt.Printf("value at idx = %d: %d\n", i, value)
	}

	fmt.Println("=======================")

	// ====================================

	var slicedNums []int = fourNums[1:3]
	// [0, 3]

	slicedNums = append(slicedNums, 1, 4, 2)
	// [0, 3, 1, 4, 2]

	for i := 0; i < len(slicedNums); i++ {
		value := slicedNums[i]
		fmt.Printf("value at idx = %d: %d\n", i, value)
	}

	fmt.Println("=======================")

	// ====================================

	maxNum := maxInt(slicedNums...)
	fmt.Printf("Max number in slicedNums: %d\n", maxNum)
}
