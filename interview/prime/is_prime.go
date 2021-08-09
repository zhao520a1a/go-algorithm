package interview

func isPrime1(x int) bool {
	// 时间复杂度 O(n)
	//for j := 2; j < i; j++ {
	//	if i%j == 0 { // 可以整除
	//		return false
	//	}
	//}

	// 优化后：时间复杂度 O(sqrt{n})
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
