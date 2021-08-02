# 无重复字符的最长子串-leetcode-003

* 贡献者：毛瑞磊
* 涉及知识点：数组，滑动窗口

## 题目

给定一个字符串，请你找出其中不含有重复字符的，最长子串的长度。

输入/输出示例：

Example 1:

```
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```


Example 2:

```
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

Example 3:

```
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
```

Example 4:

```
输入: s = ""
输出: 0
```
提示：
0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成


## 思路描述
## 思路 1
详细描述：两层循环字符串。

* 时间复杂度：O(n^2)
* 空间复杂度：O(Σ)，其中Σ 表示字符集

代码：

```go
func lengthOfLongestSubstring(s string) int {
    if "" == s {
        return 0
    }
    max := 1
    for i := 0; i < len(s); i++ {
        dict := make(map[uint8]bool)
        dict[s[i]] = true
        for j := i + 1; j < len(s); j++ {
            // 判断字符是否在map中，如果存在就直接跳出循环
            if _, ok := dict[s[j]]; ok {
                break
            }
            // 将字符放入map中
            dict[s[j]] = true
            // 每次都计算最大值
            cMax := j - i + 1
            if cMax > max {
                max = cMax
            }
        }
    }
    return max
}

```

### 思路 2
详细描述：滑动窗口。

* 时间复杂度：o(n)
* 空间复杂度：O(Σ)，其中Σ 表示字符集

代码：
```go
func lengthOfLongestSubstring(s string) int {
var (
		max  int
		l, r int
		sl   = len(s)
		dict = make(map[byte]int)
	)
	getMax := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for r = 0; r < sl; r++ {

		if t,ok:=dict[s[r]] ;ok{
			l = getMax(l,t)
		}
		max = getMax(max,r-l+1)
		dict[s[r]] = r+1
	}
	return max
}


```

### 思路 3
详细描述：在思路二的基础上用string代替map。

* 时间复杂度：O(n)
* 空间复杂度：O(Σ)，其中Σ 表示字符集

代码：
```go
func lengthOfLongestSubstring(s string) int {
	var (
		max int
		s1     string
		l, r   int
	)

	s1 = s[l:r]
	for ; r < len(s); r++ {
		if index := strings.IndexByte(s1, s[r]); index != -1 {
			l += index + 1
		}
		s1 = s[l : r+1]
		if len(s1) > max {
			max = len(s1)
		}
	}
	return max
}
```


#### 评分
* 提出思路1，能正确分析时空复杂度(2分)
* 提出思路2，能正确分析时空复杂度(3分)
* 提出思路3，能正确分析时空复杂度(4分)
* bugfree且代码简洁(5分)