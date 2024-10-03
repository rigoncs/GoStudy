package divideandconquer

func dfs(nums []int, target, i, j int) int {
	if i > j {
		return -1
	}
	m := i + ((j - i) >> 1)
	if nums[m] > target {
		return dfs(nums, target, i, m-1)
	} else if nums[m] < target {
		return dfs(nums, target, m+1, j)
	} else {
		return m
	}
}

func binarySearch(nums []int, target int) int {
	n := len(nums)
	return dfs(nums, target, 0, n-1)
}
