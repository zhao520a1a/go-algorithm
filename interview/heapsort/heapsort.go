package heapsort

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

var _ heap.Interface = (*hp)(nil)

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	v := h.IntSlice[len(h.IntSlice)-1]
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return v
}
