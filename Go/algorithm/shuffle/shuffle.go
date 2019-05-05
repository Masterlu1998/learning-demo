package shuffle

import (
	"math/rand"
	"time"
)

// Exec - 洗牌算法
func Exec(sli []int) []int {
	// 设置随机种子
	rand.Seed(int64(time.Now().UnixNano()))
	tempSli := make([]int, len(sli), cap(sli))
	copy(tempSli, sli)
	length := len(sli)
	for i := length - 1; i >= 0; i-- {
		randomIndex := rand.Intn(length)
		tempSli[randomIndex], tempSli[i] = tempSli[i], tempSli[randomIndex]
		length--
	}
	return tempSli
}