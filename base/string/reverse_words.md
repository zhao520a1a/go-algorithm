### 题目

[LeetCode 第 151 题](https://leetcode-cn.com/problems/reverse-words-in-a-string/)：给你一个字符串 s ，逐个翻转字符串中的所有 单词 。

单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。

说明：
输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
翻转后单词间应当仅用一个空格分隔。
翻转后的字符串中不应包含额外的空格。

```
示例 1
输入：s = "a good   example"
输出："example good a"
解释：如果两个单词间有多余的空格，将翻转后单词间的空格减少到只含一个。

示例 2
输入：s = "  hello world  "
输出："world hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是翻转后的字符不能包括。

```

### 解题思路

1. 移除多余空格
2. 将整个字符串反转
3. 将每个单词反转

复杂度分析：
时间复杂度：O(n) ，其中 n 为输入字符串的长度。
空间复杂度：O(n) ，双端队列存储单词需要 O(n) 的空间。

> 💡 此题还有利用双端队列的解法，在学习双端队列时会讲解到
