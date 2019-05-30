package dp

import (
	"time"
	"math/rand"
)

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    
    if n == 2 {
        return 2
    }
    
    a, b := 1, 2
    var tmp int
    for i := 3; i <= n; i++ {
        tmp = a + b
        a = b
        b = tmp
    }
    return tmp
    
}

func maxProfit(prices []int) int {
    maxResult := 0
    if len(prices) < 2 {
        return 0
    }
    
    min := prices[0]
    
    for i := 0; i < len(prices); i++ {
        cur := prices[i]
        if cur < min {
            min = cur
            continue
        }
        
        if cur - min > maxResult {
            maxResult = cur - min
        }
    }
    return maxResult
}

func MaxInt(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxSubArray(nums []int) int {
    length := len(nums)
    
    first := nums[0]
    max := first
    for i := 1; i < length; i++ {
        first = MaxInt(first + nums[i], nums[i])
        max = MaxInt(first, max)
    }
    return max
}

func rob(nums []int) int {
	copy(nums, nums)
	rand.Seed(time.Now().UnixNano())
    if len(nums) == 0 {
        return 0
	}
	cache := make(map[[2]int]int)
    return _rob(len(nums), 0, nums, cache)
}

func max(a, b int) int {
    if a > b {
        return a
    }
	return b
}

func _rob(n, flag int, nums []int, cache map[[2]int]int) int {
	if val, ok := cache[[2]int{n, flag}]; ok {
		return val
	}
	var result int
    if flag == 0 {
        // 没有偷相邻的金库
        // 最后一个结束规划
        if n == 1 {
            result = nums[0]
        } else {
			result = max(nums[0]+_rob(n-1, 1, nums[1:], cache), _rob(n-1, 0, nums[1:], cache))
		}
    } else {
		// 偷相邻的金库
		if n == 1 {
			result = 0
		} else {
			result = _rob(n-1, 0, nums[1:], cache)
		}
	}
	cache[[2]int{n, flag}] = result
	return result
}
