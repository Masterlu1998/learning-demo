package mergesort

// Sort 归并排序
// 时间复杂度 最坏：O(nlogn) 最好：O(nlogn) / 空间复杂度 O(n)
func Sort(sli []int) []int {
	tempSli := make([]int, len(sli), cap(sli))
	copy(tempSli, sli)
	return mergeSort(tempSli)
}

func merge(left []int, right []int) []int {
	// 合并，按照顺序将两边的数字排序成一个数组
	resultSli := make([]int, 0, cap(left)+cap(right))

	for len(left) != 0 && len(right) != 0 {
		if left[0] < right[0] {
			resultSli = append(resultSli, left[0])
			left = left[1:]
		} else {
			resultSli = append(resultSli, right[0])
			right = right[1:]
		}
	}

	for len(left) != 0 {
		resultSli = append(resultSli, left[0])
		left = left[1:]
	}

	for len(right) != 0 {
		resultSli = append(resultSli, right[0])
		right = right[1:]
	}
	return resultSli
}

func mergeSort(sli []int) []int {
	// 拆分，直到左右长度都为2为止
	length := len(sli)
	if length < 2 {
		return sli
	}
	middle := length / 2
	leftSli := sli[:middle]
	rightSli := sli[middle:]
	return merge(mergeSort(leftSli), mergeSort(rightSli))
}
