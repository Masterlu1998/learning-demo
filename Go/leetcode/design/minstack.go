package design

import (
	"fmt"
)

type MinStack struct {
	stack []int
	// min   int
	minstack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), minstack: []int{int(^uint(0) >> 1)}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	if x <= this.minstack[len(this.minstack) - 1] {
		this.minstack = append(this.minstack, x)
	}
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack) - 1] == this.minstack[len(this.minstack) - 1] {
		this.minstack = this.minstack[:len(this.minstack) - 1]
	}
	this.stack = this.stack[:len(this.stack) - 1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack) - 1]
}

func Test() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param_1 := obj.GetMin()
	obj.Pop()
	param_2 := obj.Top()
	param_3 := obj.GetMin()
	fmt.Println(param_1, param_2, param_3)


}
