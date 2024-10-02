package sorting

func selectionSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = j
			}
		}
		nums[min], nums[i] = nums[i], nums[min]
	}
}
