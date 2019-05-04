package selectionsort

// Sort 选择排序法
// 时间复杂度 最坏：O(n^2) 最好：O(n^2) / 空间复杂度：O(1)
func Sort(sli []int) []int {
	tempSli := make([]int, len(sli), cap(sli))
	copy(tempSli, sli)
	length := len(sli)
	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if tempSli[j] < tempSli[i] {
				tempSli[i], tempSli[j] = tempSli[j], tempSli[i]
			}
		}
	}
	return tempSli
}
