package heapsort

func Sort(sli []int) []int {
	tempSli := make([]int, len(sli), cap(sli))
	copy(tempSli, sli)
	buildMaxHeap(tempSli)
	length := len(sli)
	for i := len(tempSli) - 1; i >= 0; i-- {
		tempSli[i], tempSli[0] = tempSli[0], tempSli[i]
		length--
		heapify(tempSli, 0, length)
	}
	return tempSli
}

func buildMaxHeap(sli []int) {
	length := len(sli) / 2 - 1
	for i := length; i >= 0; i-- {
		heapify(sli, i, len(sli))
	}
}

func heapify(sli []int, node, length int) {
	left := node * 2 + 1
	right := node * 2 + 2
	largest := node
	if left < length && sli[largest] < sli[left] {
		largest = left
	} 
	if right < length && sli[largest] < sli[right] {
		largest = right
	}
	if largest != node {
		sli[largest], sli[node] = sli[node], sli[largest]
		heapify(sli, largest, length)
	}
}