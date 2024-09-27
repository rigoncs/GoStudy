package src

func binarySearchInsertionSimple(nums []int, target int) int {
	i, j := 0, len(nums)-1
	m := i + (j-i)/2
	for i <= j {
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m - 1
		} else {
			return m
		}
	}
	return i
}

func binarySearchInsertion(nums []int, target int) int {
	i, j := 0, len(nums)-1
	m := i + (j-i)/2
	for i <= j {
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m - 1
		} else {
			j = m - 1
		}
	}
	return i
}
