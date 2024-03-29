## 手撸一个 LRU 缓存

### 基本概念

LRU（英文:Least Recently Used），中文叫做最近最少使用缓存。也就是说，我们设计一个固定最大容量的缓存，当达到最大容量之后，我们再往缓存里放数据时，会先把最近最少使用的那个元素删除，再放入最新的元素。

### 目标

- 快速存取元素。
- 有固定的最大容量，不会无限制增长。
- 达到最大容量后，再往缓存里面放入新元素时，会先把最近最少使用的元素删除，再放入新元素。

### 思路

- 快速存取。我们使用 map（哈希表、映射）就可以了。
- 固定最大容量。我们给这个 map 增加 capacity（整数）字段来表示最大容量，当 map 中元素个数达到 capacity 时，我们要继续往里面放入新元素，则先把最近最少使用的元素删除再放入新元素。
- 删除最近最少使用的元素。怎么找到最近最少使用的元素删除？这个也很简单，参考 Java 中的 LinkedHashMap 的做法，把 map 中的元素使用双向链表串起来，把最近最少使用的元素放在队列尾部，最近最常使用的元素放在头部就可以了。

### 参考

[Go 语言进阶之路：手撸一个 LRU 缓存](https://blog.csdn.net/c315838651/article/details/105741886)
