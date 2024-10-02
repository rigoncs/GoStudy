package sorting

type quickSort struct{}

type quickSortMedian struct{}

type quickSortTailCall struct{}

func (q *quickSort) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[left] = nums[left], nums[i]
	return i
}

func (q *quickSort) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
}

func (q *quickSortMedian) medianThree(nums []int, left, right, mid int) int {
	l, r, m := nums[left], nums[right], nums[mid]
	if (l <= m && m <= r) || (r <= m && m <= l) {
		return mid
	} else if (l <= r && r <= m) || (m <= r && r <= l) {
		return right
	} else {
		return left
	}
}

func (q *quickSortMedian) partition(nums []int, left, right int) int {
	med := q.medianThree(nums, left, right, (left+right)/2)
	nums[left], nums[med] = nums[med], nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

func (q *quickSortMedian) quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := q.partition(nums, left, right)
	q.quickSort(nums, left, pivot-1)
	q.quickSort(nums, pivot+1, right)
}

func (q *quickSortTailCall) partition(nums []int, left, right int) int {
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= nums[left] {
			j--
		}
		for i < j && nums[i] <= nums[left] {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[i] = nums[i], nums[left]
	return i
}

func (q *quickSortTailCall) quickSort(nums []int, left, right int) {
	for left < right {
		pivot := q.partition(nums, left, right)
		if pivot-left < right-pivot {
			q.quickSort(nums, left, pivot-1)
			left = pivot + 1
		} else {
			q.quickSort(nums, pivot+1, right)
			right = pivot - 1
		}
	}
}
