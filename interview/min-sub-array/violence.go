package min_sub_array

import "math"

func minSubArrayLen(s int, num []int) (minLen int) {
	// verify
	if len(num) == 0 {
		return
	}

	minLen = math.MaxInt32
	for i := 0; i < len(num); i++ {
		sum := num[i]
		for j := i + 1; j < len(num); j++ {
			sum += num[j]
			if sum >= s {
				minLen = min(minLen, j-i+1)
				break
			}
		}
	}
	if minLen == math.MaxInt32 {
		minLen = 0
	}

	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
