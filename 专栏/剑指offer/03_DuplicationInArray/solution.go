package _3_DuplicationInArray

import "sort"

//解法1: 哈希表 / Set 利用map的特性，判断重复元素；
func findRepeatNumber1(nums []int) int {
	itemMap := make(map[int]bool)
	for _, num := range nums {
		if ok := itemMap[num]; ok {
			return num
		} else {
			itemMap[num] = true
		}
	}
	return -1
}

//解法2: 先排序，在依次遍历查找
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}

	return -1
}

//解法3：原地交换
func findRepeatNumber3(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for i != nums[i] {
			if nums[i] == nums[nums[i]] { //交换前校验
				return nums[i]
			} else {
				nums[i], nums[nums[i]] = nums[nums[i]], nums[i] //交换
			}
		}
	}
	return -1
}
