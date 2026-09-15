package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for currIDX, num := range nums {
		diff := target - num
		if prevIDX, ok := seen[diff]; ok {
			return []int{prevIDX, currIDX}
		}
		seen[num] = currIDX
	}

	return nil
}

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	var target int
	fmt.Scan(&target)

	result := twoSum(nums, target)

	if result != nil {
		fmt.Println(result)
	} else {
		fmt.Println(1)
	}
}
