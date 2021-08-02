package 找出数组中单独除出现的数字

// 异或位运算
func findAloneNum(arr []int) int {
	num := 0
	for _, item := range arr {
		num ^= item
	}
	return num
}

func findAloneTwoNum(arr []int) (num1 int, num2 int) {
	ret := findAloneNum(arr)
	var i uint
	// //先找到一位为1的bit位
	for i = 0; i < 32; i++ {
		if ((ret >> i) & 1) == 1 {
			break
		}
	}

	// //按照这一位将数组分成两部分分别异或
	for _, item := range arr {
		if (item>>i)&1 == 0 {
			num1 ^= item
		} else {
			num2 ^= item
		}
	}

	if num1 > num2 {
		num1, num2 = num2, num1
	}

	return
}
