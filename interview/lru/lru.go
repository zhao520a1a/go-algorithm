package lru

type Entry struct {
	pre   *Entry
	next  *Entry
	Key   string
	Value interface{}
}

type Cache struct {
	capacity int
	cache    map[string]*Entry
	head     *Entry
	tail     *Entry
}

func NewCache(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
		cache:    make(map[string]*Entry, capacity),
		head:     nil,
		tail:     nil,
	}
}

func (c *Cache) Put(key string, val interface{}) {
	e := &Entry{
		Key:   key,
		Value: val,
	}
	entry, ok := c.cache[key]
	if !ok {
		// 元素不存在:把head设置为当前元素
		if c.head != nil {
			c.head.pre = e
			e.next = c.head
		}
		if c.tail == nil {
			c.tail = e
		}
		c.head = e
		c.cache[key] = e

		// -- 当放入新元素之后，缓存元素个数达到最大容量，需要把队列尾部的元素删掉
		if len(c.cache) > c.capacity {
			if c.tail != nil {
				removeEntry := c.tail
				c.tail = c.tail.pre
				removeEntry.next = nil
				c.tail.next = nil
				delete(c.cache, removeEntry.Key)
			}
		}
		return
	}
	// 元素已经存在:提到队列头部
	c.moveToHead(entry)
	return
}

func (c *Cache) moveToHead(entry *Entry) {
	if entry == c.head {
		return
	}
	// - 剥离元素: 当前节点可能是 tail 节点
	entry.pre.next = entry.next
	if entry != c.tail {
		entry.next.pre = entry.pre
	} else {
		c.tail = entry.pre
	}
	// - 放到头部
	entry.pre = nil
	entry.next = c.head
	c.head.pre = entry
	c.head = entry
	return
}

func (c *Cache) Get(key string) interface{} {
	if entry, ok := c.cache[key]; ok {
		// 元素存在:提到队列头部
		c.moveToHead(entry)
		return entry.Value
	}
	// 元素不存在，返回nil
	return nil
}
