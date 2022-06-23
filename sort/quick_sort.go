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
	p := slice[highIndex]
	i := lowIndex
	for j := lowIndex; j < highIndex; j++ {
		if comparator.Compare(slice[j], p) == -1 { // slice[j] < p
			swap(slice, i, j)
			i++
		}
	}

	swap(slice, i, highIndex)

	return i
}

// swap two slice value at index i and j
func swap[T any](slice []T, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
