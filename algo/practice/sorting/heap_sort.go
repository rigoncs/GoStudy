package sorting

func siftDown(nums *[]int, n, i int) {
	for true {
		l := 2*i + 1
		r := 2*i + 2
		ma := i
		if l < n && (*nums)[l] > (*nums)[ma] {
			ma = l
		}
		if r < n && (*nums)[r] > (*nums)[ma] {
			ma = r
		}
		if ma == i {
			break
		}
		(*nums)[i], (*nums)[ma] = (*nums)[ma], (*nums)[i]
		i = ma
	}
}

func heapSort(nums *[]int) {
	for i := len(*nums)/2 - 1; i >= 0; i-- {
		siftDown(nums, len(*nums), i)
	}
	for i := len(*nums) - 1; i > 0; i-- {
		(*nums)[i], (*nums)[0] = (*nums)[0], (*nums)[i]
		siftDown(nums, i, 0)
	}
}
