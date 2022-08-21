package interview

/*
10000以内的所有质数的输出
实现：
*/

func primeNumSum1(n int) int {
	sum := 0
	for i := 2; i < n; i++ {
		if isPrime1(i) {
			sum += i
		}
	}
	return sum
}

// 使用"埃氏筛"求质数,时间复杂度O(n log log n) 接近于 O(n)
func primeNumSum2(n int) int {
	isPrimeList := make([]bool, n)
	sum := 0

	for i := range isPrimeList {
		isPrimeList[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrimeList[i] {
			sum += i
			for j := i * i; j < n; j += i {
				isPrimeList[j] = false
			}
		}
	}
	return sum
}
