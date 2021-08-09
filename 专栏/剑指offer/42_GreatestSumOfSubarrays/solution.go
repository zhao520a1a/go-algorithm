package _2_GreatestSumOfSubarrays

/*
题目：https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/

复杂度
时间复杂度：O(n)，其中 n 为 nums 数组的长度。
空间复杂度：O(1)，只需要常数空间存放若干变量。

*/

func maxSubArray(nums []int) int {
	var preSum int
	maxSum := nums[0]

	for i := 1; i < len(nums)-1; i++ {
		preSum = max(preSum+nums[i], nums[i])
		maxSum = max(preSum, maxSum)
	}

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
