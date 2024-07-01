
![Image](https://img.shields.io/badge/language-muti-brightgreen.svg)
![Image](https://img.shields.io/badge/leetcode-50%2B-orange.svg)

> 引言 ： 数据结构是为算法服务的，算法要作用在特定的数据结构之上。

![Image](images/banner.png)


# 二叉树框架思维
二叉树这种结构无非就是二叉链表

## 思维模式
- 「遍历」的思维模式：
- 「分解问题」的思维模式：通过子问题的答案推导出原问题的答案；

## 观点认知

### 重新理解前中后序
前中后序是遍历二叉树过程中处理每一个节点的三个特殊时间点。
- 前序位置的代码在刚刚进入一个二叉树节点的时候执行；
- 后序位置的代码在将要离开一个二叉树节点的时候执行；
- 中序位置的代码在一个二叉树节点左子树都遍历完，即将开始遍历右子树的时候执行。
![image](https://github.com/zhao520a1a/go-algorithm/assets/18511674/66cfa393-c6f8-44b3-8eb5-e16b97609835)

### 快速排序就是个二叉树的前序遍历，归并排序就是个二叉树的后序遍历
```
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
```
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

## 



## 其他

### 文档格式化

```
npm run format
```

引用：https://prettier.io/docs/en/install.html

### 推荐项目
[fucking-algorithm](https://github.com/labuladong/fucking-algorithm)
[duke-git/lancet](https://github.com/duke-git/lancet/blob/main/docs/algorithm_zh-CN.md)
