package src

func binarySearchLeftEdge(nums []int, target int) int {
	i := binarySearchInsertion(nums, target)
	if i == len(nums) || nums[i] != target {
		return -1
	}
	return i
}

func binarySearchRightEdge(nums []int, target int) int {
	i := binarySearchLeftEdge(nums, target+1)
	j := i - 1
	if j == -1 || nums[j] != target {
		return -1
	}
	return j
}
