package sorting

func countingSortNative(nums []int) {
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	for i, num := 0, 0; num < m+1; num++ {
		for j := 0; j < counter[num]; j++ {
			nums[i] = num
			i++
		}
	}
}

func countingSort(nums []int) {
	m := 0
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	counter := make([]int, m+1)
	for _, num := range nums {
		counter[num]++
	}
	for i := 0; i < m; i++ {
		counter[i+1] += counter[i]
	}
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		num := nums[i]
		res[counter[num]-1] = num
		counter[num]--
	}
	copy(nums, res)
}
