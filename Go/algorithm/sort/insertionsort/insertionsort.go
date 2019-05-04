package insertionsort

// Sort 插值排序法
// 时间复杂度 最坏：O(n^2) 最好：O(n)  / 空间复杂度 O(1)
func Sort(sli []int) []int {
	tempSli := make([]int, len(sli), cap(sli))
	copy(tempSli, sli)
	length := len(sli)
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if tempSli[j+1] < tempSli[j] {
				tempSli[j+1], tempSli[j] = tempSli[j], tempSli[j + 1]
			}
		}
	}
	return tempSli
}
