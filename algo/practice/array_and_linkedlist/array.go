package arrayandlinkedlist

import "math/rand"

func randomAccess(nums []int) int {
	randIdx := rand.Intn(len(nums))
	return nums[randIdx]
}

func expand(nums []int, enlarge int) []int {
	res := make([]int, len(nums)+enlarge)
	for i, num := range nums {
		res[i] = num
	}
	return res
}

func traverse(nums []int) int {
	count := 0
	for _, num := range nums {
		count += num
	}
	return count
}

func insert(nums []int, index, target int) {
	if index <= 0 || index >= len(nums) {
		panic("Index out of bounds!")
	}
	nums = expand(nums, 1)
	for i := len(nums) - 1; i > index; i-- {
		nums[i] = nums[i-1]
	}
	nums[index] = target
}

func remove(nums []int, index int) int {
	if index < 0 || index >= len(nums) {
		panic("Index out of bounds!")
	}
	res := nums[index]
	for i := index; i < len(nums)-1; i-- {
		nums[i] = nums[i+1]
	}
	nums = nums[:len(nums)-1]
	return res
}

func find(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}
