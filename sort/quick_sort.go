package sort

// Comparator is for comparing two values
type Comparator interface {
	// Compare v1 and v2
	// Ascending order: should return 1 -> v1 > v2, 0 -> v1 = v2, -1 -> v1 < v2
	// Descending order: should return 1 -> v1 < v2, 0 -> v1 = v2, -1 -> v1 > v2
	Compare(v1, v2 any) int
}

// QuickSort quick sorting for slice, lowIndex is 0 and highIndex is len(slice)-1
func QuickSort[T any](slice []T, lowIndex, highIndex int, comparator Comparator) {
	if lowIndex < highIndex {
		p := partition(slice, lowIndex, highIndex, comparator)
		QuickSort(slice, lowIndex, p-1, comparator)
		QuickSort(slice, p+1, highIndex, comparator)
	}
}

// partition split slice into two parts
func partition[T any](slice []T, lowIndex, highIndex int, comparator Comparator) int {
	marker := slice[highIndex]

	// 1. 比大小后，找到 marker 坐在的坐标
	i := lowIndex
	for j := lowIndex; j < highIndex; j++ {
		if comparator.Compare(slice[j], marker) == -1 { // slice[j] < p
			swap(slice, i, j) // 令坐标 i 之前的值都比 maker 小， 坐标 i 之后的值都比 maker 大
			i++
		}
	}

	// 2. 将坐标i上的值更新为value
	swap(slice, i, highIndex)

	return i
}

// swap two slice value at index i and j
func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
