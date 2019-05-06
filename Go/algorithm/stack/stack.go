package stack

type Stack struct {
	length int
	stack []int
}

// Push - 入栈
func (this *Stack)Push(i int) {
	this.length++
	this.stack = append(this.stack, i)
}

// Pop - 弹栈
func (this *Stack)Pop() int {
	lastIndex := this.length - 1
	val := this.stack[lastIndex]
	this.stack = this.stack[:lastIndex]
	this.length--
	return val
}