package array

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

// singleNumber - 只出现一次的数字
func singleNumber(nums []int) int {
	result := 0
	for _, val := range nums {
		result = result^val 
	}
	return result
}


// plusOne - 加一
func plusOne(digits []int) []int {
	length := len(digits)
	lastIndex := length - 1
	for i := lastIndex; i >= 0; i-- {
		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			return digits
		}
	}
	newDigits := make([]int, 1, length + 1)
	newDigits[0] = 1
	newDigits = append(newDigits, digits...)
	return newDigits
}

// moveZeroes - 移动0
func moveZeroes(nums []int) {
	current := 0
	for detect := 0; detect < len(nums); detect++ {
		if nums[detect] != 0 {
			nums[current] = nums[detect]
			current++
		}
	}

	for i := current; i < len(nums); i++ {
		nums[i] = 0
	}
}

// twoSum - 两数之和
func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cache[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		remain := target - nums[i]
		if index, ok := cache[remain]; ok {
			if index == i {
				continue
			} else {
				return []int{i, index}
			}
		}
	}
	return []int{}
}

func rotateMa(matrix [][]int)  {
	length := len(matrix)
	for i := 0; i < length / 2; i++ {
		for j := i; j < length - i - 1; j++ {
			// 内层方阵的第n个元素
			n := j-i
			// 内层方阵的阶层-1
			x := length - 2*i - 1
			matrix[j][i], matrix[j+x-n][i+n], matrix[j+x-2*n][i+x], matrix[j-n][i+x-n] = matrix[j+x-n][i+n], matrix[j+x-2*n][i+x], matrix[j-n][i+x-n], matrix[j][i]  
		} 
	}
}

