package bubblesort

// Sort 冒泡排序法
// 时间复杂度 最坏：O(n^2) 最好：O(n-1) / 空间复杂度：O(1)
func Sort(sli []int) []int {
	tempArr := make([]int, len(sli), cap(sli))
	copy(tempArr, sli)
	length := len(sli)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if tempArr[j] > tempArr[j+1] {
				tempArr[j], tempArr[j+1] = tempArr[j+1], tempArr[j]
			}
		}
	}
	return tempArr
}
