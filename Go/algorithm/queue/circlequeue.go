package queue

type CircleQueue struct {
	capacity int
	queue    []int
	head     int
	tail     int
}

// New - 创建循环队列
func NewCircleQueue(cap int) *CircleQueue {
	return &CircleQueue{capacity: cap, head: 0, tail: 0, queue: make([]int, cap)}
}

// EnQueue - 入队
func (this *CircleQueue) EnQueue(i int) {
	if (this.tail+1)%this.capacity == this.head {
		panic("队列已满，插入失败")
	}
	this.queue[this.tail] = i
	this.tail++
	this.tail = this.tail % this.capacity
}

// OutputQueue - 出队
func (this *CircleQueue) OutputQueue() int {
	if this.head%this.capacity == this.tail {
		panic("队列为空")
	}
	result := this.queue[this.head]
	this.head++
	this.head = this.head % this.capacity
	return result
}
