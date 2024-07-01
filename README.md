
![Image](https://img.shields.io/badge/language-muti-brightgreen.svg)
![Image](https://img.shields.io/badge/leetcode-50%2B-orange.svg)

> 引言 ： 数据结构是为算法服务的，算法要作用在特定的数据结构之上。

<img src="images/banner.png" alt="前中后序" style="zoom:50%;" />


# 两种思维模式
- 「遍历」的思维模式：
- 「分解问题」的思维模式：通过子问题的答案推导出原问题的答案；


# 二叉树

## 观点认知

### 重新理解前中后序
前中后序是遍历二叉树过程中处理每一个节点的三个特殊时间点，绝不仅仅是三个顺序不同的 List；
- 前序位置的代码在刚刚进入一个二叉树节点的时候执行；
- 后序位置的代码在将要离开一个二叉树节点的时候执行；
- 中序位置的代码在一个二叉树节点左子树都遍历完，即将开始遍历右子树的时候执行。

<img src="https://github.com/zhao520a1a/go-algorithm/assets/18511674/66cfa393-c6f8-44b3-8eb5-e16b97609835" alt="前中后序" style="zoom:50%;" />

二叉树遍历框架:
``` go
// 基本的二叉树节点
type TreeNode struct {
    val int
    left *TreeNode
    right *TreeNode
}
func traverse(root *TreeNode) {
    if root == nil {
    		return
    }
    // 前序位置
    traverse(root.left)
    // 中序位置
    traverse(root.right)
    // 后序位置
}
```

#### 二叉树这种结构无非就是二叉链表
```
// 递归遍历单链表
func traverse(head *ListNode) {
    if head == nil {
        return
    }
    // 前序位置
    traverse(head.Next)
    // 后序位置
}
```
#### 快速排序就是个二叉树的前序遍历，归并排序就是个二叉树的后序遍历
``` go
// 快速排序
func sort(nums []int, lo, hi int) {
    /****** 前序遍历位置 ******/
    // 通过交换元素构建分界点 p
    p := partition(nums, lo, hi)
    /************************/

    sort(nums, lo, p - 1)
    sort(nums, p + 1, hi)
}
```
``` go
// 归并排序
func sort(nums []int, lo int, hi int) {
    mid := (lo + hi) / 2
    // 排序 nums[lo..mid]
    sort(nums, lo, mid)
    // 排序 nums[mid+1..hi]
    sort(nums, mid + 1, hi)

    /****** 后序位置 ******/
    // 合并 nums[lo..mid] 和 nums[mid+1..hi]
    merge(nums, lo, mid, hi)
    /*********************/
}
```


## 解题思路
二叉树的所有问题，就是让你在前中后序位置注入巧妙的代码逻辑，去达到自己的目的，你只需要单独思考每一个节点应该做什么，其他的不用你管，抛给二叉树遍历框架，递归会在所有节点上做相同的操作。

## 两种解题思路
二叉树题目的递归解法可以分两类思路:
第一类是遍历一遍二叉树得出答案 -> 回溯算法核心框架

第二类是通过分解问题计算出答案 -> 动态规划核心框架



# 推荐项目
[fucking-algorithm](https://github.com/labuladong/fucking-algorithm)
[duke-git/lancet](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md)

## 文档格式化
引用：https://prettier.io/docs/en/install.html
```
npm run format
```
