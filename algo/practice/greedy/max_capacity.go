package greedy

import "math"

func maxCapacity(ht []int) int {
	i, j := 0, len(ht)-1
	maxCap := 0
	for i < j {
		cap := (j - i) * int(math.Min(float64(ht[i]), float64(ht[j])))
		maxCap = int(math.Max(float64(maxCap), float64(cap)))
		if ht[i] < ht[j] {
			i++
		} else {
			j--
		}
	}
	return maxCap
}
