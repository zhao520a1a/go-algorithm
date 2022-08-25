### 题目

[LeetCode 第 230 题](https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/)：给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

### 解题思路

这道题考察了两个知识点：

二叉搜索树的性质

二叉搜索树的遍历

二叉搜索树的性质：对于每个节点来说，该节点的值比左孩子大，比右孩子小，而且一般来说，二叉搜索树里不出现重复的值。

二叉搜索树的中序遍历是高频考察点，节点被遍历到的顺序是按照节点数值大小的顺序排列好的。即，中序遍历当中遇到的元素都是按照从小到大的顺序出现。

因此，我们只需要对这棵树进行中序遍历的操作，当访问到第 k 个元素的时候返回结果就好。

![img](https://gitee.com/zhaojinxin_golden/article-images/raw/master/typora/CgoB5l2IRaOAag5tAHlWAofWh6A551.gif)

注意：这道题可以变成求解第 K 大的元素，方法就是对这个二叉搜索树进行反向的中序遍历，那么数据的被访问顺序就是由大到小了。
