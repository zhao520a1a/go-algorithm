package _3_DuplicationInArray

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
