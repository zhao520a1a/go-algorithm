package 识别山脉数组

func validMountainArray(arr []int) bool {
	i, n := 0, len(arr)

	// 递增扫描
	for ; i+1 < n && arr[i] < arr[i+1]; i++ {
	}

	// 最高点不能是数组的第一个位置或最后一个位置
	if i == 0 || i == n-1 {
		return false
	}

	// 递减扫描
	for ; i+1 < n && arr[i] > arr[i+1]; i++ {
	}

	return i == n-1
}
