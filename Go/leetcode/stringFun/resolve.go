package stringfun

import (
	"strconv"
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

func myAtoi(str string) int {
	// 过滤开头的空格
	startIndex := 0
	calFlag := 1
	for index, letter := range str {
		if letter != ' ' {
			isValidate := (letter >= '0' && letter <= '9') || letter == '+' || letter == '-'
			if isValidate {
				startIndex = index
				if letter == '-' {
					calFlag = -1
					startIndex++
				}
				if letter == '+' {
					calFlag = 1
					startIndex++
				}
				break;
			} else {
				return 0
			}
		}
	}

	// 开始转换
	result := 0
	for i := startIndex; i < len(str); i++ {
		isValidate := (str[i] >= '0' && str[i] <= '9')
		if isValidate {
			if str[i] >= '0' && str[i] <= '9' {
				result = result * 10 + int(str[i] - 48) * calFlag
				// 边界检测
				if result > math.MaxInt32 {
					return math.MaxInt32
				}
	
				if result < math.MinInt32 {
					return math.MinInt32
				}
			} 
		} else {
			break
		}
	}
	return result
}

func TestMyAtoi() {
	b := "    0-1"
	fmt.Println(myAtoi(b))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	outloop:
	for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		if haystack[i] == needle[0] {
			for j := 1; j < len(needle); j++ {
				if haystack[i+j] != needle[j] {
					continue outloop
				}
			}
			return i
		}
	}
	return -1
}

func TestStrStr() {
	a := "hell"
	b := "ll"
	fmt.Println(strStr(a, b))
}

// 递归方法
func countAndSay(n int) string {
    if n == 1 {
		return "1"
	}

	var result string
	pre := countAndSay(n-1)
	count := 1
	preLetter := string(pre[0])
	for i := 1; i < len(pre); i++ {
		if string(pre[i]) == preLetter {
			count++
		} else {
			result = result + strconv.Itoa(count) + preLetter
			count = 1
		}
		preLetter = string(pre[i])
	}
	result = result + strconv.Itoa(count) + preLetter

	return result
}

func TestCountAndSay() {
	fmt.Println(countAndSay(6))
}

func minIndex(strs []string) int {
	min := math.MaxInt64
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}
	return min
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 	{
		return ""
	}
	minIndex := minIndex(strs)
	result := ""
	outloop:
	for i := 0; i < minIndex; i++ {
		common := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if common != strs[j][i] {
				break outloop
			}
		}
		result += string(common)
	}
	return result
}

func TestLongestCommonPrefix() {
	fmt.Println(longestCommonPrefix([]string{}))
}