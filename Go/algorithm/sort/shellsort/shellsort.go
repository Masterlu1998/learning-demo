package shellsort

// Sort 希尔排序
// 时间复杂度 最坏：O(nlog^2n)  最好：O(nlog^2n) / 空间复杂度
func Sort(sli []int) []int {
	tempSli := make([]int, len(sli), cap(sli))
	copy(tempSli, sli)
	length := len(tempSli)
	// 增量递减
	for gep := length / 2; gep > 0; gep /= 2 {
		// 给每个子序列进行插值排序
		for i := gep; i < length; i++ {
			for j := i - gep; j >=0; j-=gep {
				if (tempSli[j] > tempSli[j + gep]) {
					tempSli[j], tempSli[j + gep] = tempSli[j + gep], tempSli[j]
				}
			}
		}
	}
	return tempSli
}