package array

import (
	"fmt"
)

// -------- 数组类别算法 --------

// maxProfit - 买卖股票的最佳时机 II
func maxProfit(prices []int) int {
	profix := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profix += prices[i] - prices[i-1]
		}
	}
	return profix
}

// MaxProfitTest - 测试
func TestMaxProfit() {
	arr := []int{7, 6, 4, 3, 1}
	profix := maxProfit(arr)
	fmt.Println(profix)
}

// rotate - 旋转数组
func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		current := nums[len(nums) - 1]
		for j := len(nums) - 1; j >= 1; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = current 
	}
}

func TestRotate() {
	arr := []int{1,2,3,4,5,6,7}
	rotate(arr, 3)
	fmt.Println(arr)
}

// singleNumber - 只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for _, val := range nums {
		result = result^val 
	}
	return result
}

func TestSingleNumber() {
	sli := []int{1, 3, 4, 3, 4}
	fmt.Println(singleNumber(sli))
}


