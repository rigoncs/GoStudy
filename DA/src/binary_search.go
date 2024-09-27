package src

func binarySearch(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] < val {
			i = mid + 1
		} else if nums[mid] > val {
			j = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func binarySearchLCRO(nums []int, target int) int {
	i, j := 0, len(nums)
	for i < j {
		m := i + (j-i)/2
		if nums[m] < target {
			i = m + 1
		} else if nums[m] > target {
			j = m
		} else {
			return m
		}
	}
	return -1
}
