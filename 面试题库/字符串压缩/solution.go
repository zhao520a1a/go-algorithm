package 字符串压缩

import (
	"bytes"
	"strconv"
)

func compressString(str string) string {
	if len(str) == 0 {
		return str
	}

	res := bytes.NewBufferString("")
	count := 1
	ch := str[0]
	for i := 1; i < len(str); i++ {
		if ch == str[i] {
			count++
		} else {
			res.WriteByte(ch)
			res.WriteString(strconv.Itoa(count))
			ch = str[i]
			count = 1
		}
	}
	res.WriteByte(ch)
	res.WriteString(strconv.Itoa(count))

	if len(res.String()) >= len(str) {
		return str
	}
	return res.String()
}
