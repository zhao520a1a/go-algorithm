package interview

import "fmt"

/*
10000以内的所有质数的输出
实现：
*/

func primeNumSum() int {
	var flag bool
	var sum int
	limit := 10000
	for i := 2; i < limit; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 { // 可以整除
				flag = true
				break
			}
		}
		if flag {
			sum += i
		}
	}
	fmt.Println(sum)

	return sum
}
