package src

import "math"

func maxCapacity(ht []int) int {
	i, j := 0, len(ht)-1
	res := 0
	for i < j {
		capacity := int(math.Min(float64(ht[i]), float64(ht[j]))) * (j - i)
		res = int(math.Max(float64(res), float64(capacity)))
		if ht[i] < ht[j] {
			i++
		} else {
			j--
		}
	}
	return res
}
