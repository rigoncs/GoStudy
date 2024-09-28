package src

import "math"

func maxProductCutting(n int) int {
	if n <= 3 {
		return 1 * (n - 1)
	}
	a := n / 3
	b := n % 3
	if b == 1 {
		return int(math.Pow(3, float64(a-1))) * 2 * 2
	}
	if b == 2 {
		return int(math.Pow(3, float64(a))) * 2
	}
	return int(math.Pow(3, float64(a)))
}
