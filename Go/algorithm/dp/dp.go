package dp

// 动态规划算法
// 问题：有一个国家发现了5座金矿，每座金矿的黄金储量不同，需要参与挖掘的工人数也不同。参与挖矿工人的总数是10人。
// 每座金矿要么全挖，要么不挖，不能派出一半人挖取一半金矿。要求用程序求解出，要想得到尽可能多的黄金，
// 应该选择挖取哪几座金矿？

import (
	// "fmt"
	"math"
)

// 每座金矿产量g与所需工人数量p数组如下
var (
	g = []int{500, 400, 350, 300, 200}
	p = []int{5, 5, 3, 4, 3}
)

func MaxInt(a, b int) int {
	tempA := float64(a)
	tempB := float64(b)
	result := math.Max(tempA, tempB)
	return int(result)
}

// RecruitCal - 递归算法 n 金矿数量 w工人数量
func RecruitCal(n, w int) int {
	arrIndex := n - 1
	if n <= 1 {
		if w < p[0] {
			return 0
		}
		return g[0]
	}

	if w < p[arrIndex] {
		return RecruitCal(n-1, w)
	}
	return MaxInt(RecruitCal(n-1, w), g[arrIndex]+RecruitCal(n-1, w-p[arrIndex]))
}

// 备忘录算法
func MemoCal(n, w int, memo map[[2]int]int) int {
	// 如果命中缓存，直接返回
	if val, ok:= memo[[2]int{n, w}]; ok {
		return val
	}
	var result int
	arrIndex := n - 1
	if n <= 1 {
		if w < p[arrIndex] {
			result = 0
		} else {
			result = g[arrIndex]
		}
	} else {
		if w < p[arrIndex] {
			result = RecruitCal(n-1, w)
		} else {
			result = MaxInt(RecruitCal(n-1, w), g[arrIndex] + RecruitCal(n-1, w-p[arrIndex]))
		}
	}
	memo[[2]int{n, w}] = result
	return result
}

// DPCal - 动态规划算法
func DPCal(n, w int) int {
	preresults := make([]int, w+1)
	results := make([]int, w+1)

	// 填充第一行数据
	for i := 0; i <= w; i++ {
		if i < p[0] {
			preresults[i] = 0
		} else {
			preresults[i] = g[0]
		}
	}

	// 动态规划
	for i := 1; i < n; i++ {
		for j := 0; j <= w; j++ {
			if j < p[i] {
				results[j] = preresults[j]
			} else {
				results[j] = MaxInt(preresults[j], preresults[j-p[i]]+g[i])
			}
		}
		copy(preresults, results)
	}
	return results[w]
}
