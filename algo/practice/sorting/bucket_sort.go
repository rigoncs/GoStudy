package sorting

import "sort"

func bucketSort(nums []float64) {
	k := len(nums) / 2
	buckets := make([][]float64, k)
	for i := range buckets {
		buckets[i] = make([]float64, 0)
	}
	for _, num := range nums {
		i := int(num * float64(k))
		buckets[i] = append(buckets[i], num)
	}
	for i := 0; i < k; i++ {
		sort.Float64s(buckets[i])
	}
	i := 0
	for _, bucket := range buckets {
		for _, num := range bucket {
			nums[i] = num
			i++
		}
	}
}
