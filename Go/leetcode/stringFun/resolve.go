package stringfun

import (
	"strings"
	"math"
	"fmt"
)

// -------- 字符串类别算法 --------

// reverseString - 反转字符串
func reverseString(s []byte) {
	length := len(s)
	for i := 0; i < length/2; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}

func TestReverseString() {
	s := []byte{'8', '2', '3'}
	reverseString(s)
	fmt.Println(string(s))
}

func reverse(x int) int {
	handleX := x
	if x < 0 {
		handleX = -x
	}
	var result int
	for handleX > 0 {
		num := handleX % 10
		result = result * 10 + num
		handleX /= 10 
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	if x < 0 {
		result = -result
	}
	return result
}

func TestReverse() {
	a := -3320
	fmt.Println(reverse(a))
}

func firstUniqChar(s string) int {
	cache := make(map[string][]int)
	for index, letter := range s {
		cache[string(letter)] = append(cache[string(letter)], index)
	}
	const MAX = int(^uint(0) >> 1)
	minIndex := MAX
	for _, value := range cache {
		if len(value) == 1 && value[0] < minIndex {
			minIndex = value[0]
		}
	}

	if minIndex == MAX {
		return -1
	}
	return minIndex
}

func TestFirstUniqChar() {
	fmt.Println(firstUniqChar("loveleetcode"))
}


func isAnagram(s string, t string) bool {
		cache := make(map[string]int)
		for _, letter := range s {
			cache[string(letter)]++
		}

		for _, letter := range t {
			if _, ok := cache[string(letter)]; ok {
				cache[string(letter)]--
				continue
			} else {
				return false
			}
		}

		for _, value := range cache {
			if value != 0 {
				return false
			}
		}
	
		return true
}

func TestIsAnagram() {
	fmt.Println(isAnagram("anagm", "nagaram"))
}

func isValidate(letter byte) bool {
	isLowercase := letter >= 97 && letter <= 122
	isUppercase := letter >= 65 && letter <= 90
	isNumber := letter >= 48 && letter <= 57
	return isLowercase || isUppercase || isNumber
}

func isPalindrome(s string) bool {
	tempS := strings.Replace(strings.ToLower(s), " ", "", -1)

	var pureS []byte
	for _, letter := range tempS {
		if isValidate(byte(letter)) {
			pureS = append(pureS, byte(letter))
		}
	}
	fmt.Println(string(pureS))

	length := len(pureS)
	middle := length / 2
	var stack []byte
	for i := 0; i < middle; i++ {
		stack = append(stack, pureS[i])
	}
	if remain := length % 2; remain != 0 {
		for i := middle + 1 ; i < length; i++ {
			length := len(stack)
			last := stack[length - 1]
			if pureS[i] != last {
				return false
			}
			stack = stack[:length - 1]
		}
	} else {
		for i := middle; i < length; i++ {
			length := len(stack)
			last := stack[length - 1]
			if pureS[i] != last {
				return false
			}
			stack = stack[:length - 1]
		}
	}

	return true
}

func TestIsPalindrome() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}