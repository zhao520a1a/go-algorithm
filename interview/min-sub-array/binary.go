package min_sub_array

import (
	"math"
	"sort"
)

func minSubArrayLenBinary(s int, nums []int) (minLen int) {
	// verify
	if len(nums) == 0 {
		return
	}

	minLen = math.MaxInt32
	// 为了方便计算，令 size = n + 1
	// prefixSum[0] = 0 意味着前 0 个元素的前缀和为 0
	// prefixSum[1] = A[0] 前 1 个元素的前缀和为 A[0]
	prefixSums := make([]int, len(nums)+1)
	for i, num := range nums {
		prefixSums[i+1] = prefixSums[i] + num
	}

	for i, prefixSum := range prefixSums {
		target := prefixSum + s
		bound := sort.SearchInts(prefixSums, target)
		if bound <= len(prefixSums) {
			minLen = min(minLen, bound-i+1)
		}
	}

	if minLen == math.MaxInt32 {
		minLen = 0
	}

	return

}
