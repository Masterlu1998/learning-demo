package design

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
    Origin []int
    Now []int
}


func Constructor1(nums []int) Solution {
    return Solution{Origin: nums, Now: make([]int, len(nums), len(nums))}
}


/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	copy(this.Now, this.Origin)
	fmt.Println(this.Now)
    return this.Now
}

func (this *Solution) Get() {
	fmt.Println(this.Now, this.Origin)
}


/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
    copy(this.Now, this.Origin)
    for i := len(this.Now) - 1; i >= 1; i-- {
		time.Sleep(1*time.Millisecond)
        rand.Seed(time.Now().UnixNano())
		swtichIndex := rand.Intn(i)
        this.Now[swtichIndex], this.Now[i] = this.Now[i], this.Now[swtichIndex]
	}
	fmt.Println(this.Now)
    return this.Now
}



