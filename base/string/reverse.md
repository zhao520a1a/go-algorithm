### 题目

翻转字符串“algorithm”。



### 解法

用两个指针，一个指向字符串的第一个字符 a，一个指向它的最后一个字符 m，然后互相交换。交换之后，两个指针向中央一步步地靠拢并相互交换字符，直到两个指针相遇。这是一种比较快速和直观的方法。

注意：由于无法直接修改字符串里的字符，所以必须先把字符串变换为数组，然后再运用这个算法。

动态算法演示：

 ![img](https://gitee.com/zhaojinxin_golden/article-images/raw/master/typora/CgoB5l2IRiCATj5LAGJa69BtQRA357.gif)

