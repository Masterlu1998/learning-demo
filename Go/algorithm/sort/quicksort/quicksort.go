package quicksort

// Sort 快速排序
// 时间复杂度 最坏：O(nlogn) 最好：O(n^2) / 空间复杂度：O(nlogn)
func Sort(sli []int) []int {
	tempSli := make([]int, len(sli), cap(sli))
	copy(tempSli, sli)
	resultSli := quicksort(tempSli, 0, len(tempSli))
	return resultSli
}

func quicksort(sli []int, left, right int) []int {
	if left < right {
		index := findIndex(sli, left, right)
		quicksort(sli, left, index - 1)
		quicksort(sli, index + 1, right)
	}
	return sli
}

func findIndex(sli []int, left, right int) int {
	// 定义标准
	pivot := left
	index := pivot + 1

	for i := index; i < right; i++ {
		if sli[i] < sli[pivot] {
			sli[i], sli[index] = sli[index], sli[i]
			index++
		}
	}
	sli[index - 1], sli[pivot] = sli[pivot], sli[index - 1]
	return index - 1
}