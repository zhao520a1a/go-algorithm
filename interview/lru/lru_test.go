package lru

import (
	"fmt"
	"testing"
)

func TestLRU(t *testing.T) {
	cache := NewCache(2)
	cache.Put("1", "one")       // 放入元素one，此时one在队列头部
	fmt.Println(cache.Get("1")) // 此处输出“one”，此时one在队列头部
	cache.Put("2", "two")       // 放入元素two，此时two在队列头部
	fmt.Println(cache.Get("1")) // 此处输出“one”，此时one在队列头部
	cache.Put("3", "three")     // 放入元素three，总元素个数为3，因此最近最少使用的元素“2”会被删除
	fmt.Println(cache.Get("2")) // 此处输出nil
	fmt.Println(cache.Get("3")) // 此处输出“three”
	fmt.Println(cache.Get("3")) // 此处输出“three”
	fmt.Println(cache.Get("1")) // 此处输出“one”，此时最近最少使用的元素为“3”
	cache.Put("2", "two")       // 放入元素three，总元素个数为3，因此最近最少使用的元素“3”会被删除
	fmt.Println(cache.Get("3")) // 此处输出nil
	fmt.Println(cache.Get("1")) // 此处输出one
}
