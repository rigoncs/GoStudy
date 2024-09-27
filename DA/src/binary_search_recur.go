package src

func dfsSearch(nums []int, target, i, j int) int {
	if i > j {
		return -1
	}
	m := i + ((j - i) >> 1)
	if nums[m] < target {
		return dfsSearch(nums, target, m+1, j)
	} else if nums[m] > target {
		return dfsSearch(nums, target, i, m-1)
	} else {
		return m
	}
}

func binarySearchRecur(nums []int, target int) int {
	n := len(nums)
	return dfsSearch(nums, target, 0, n-1)
}
