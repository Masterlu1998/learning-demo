package queue

type CircleQueue struct {
	Capacity int
	queue [4]int
	head int
	tail int
}

// New - 创建循环队列
func New(cap int) CircleQueue {
	return CircleQueue{Capacity: cap, head: 0, tail: 0}
}

// EnQueue - 入队
func (this *CircleQueue) EnQueue(i int) {
	if ((this.tail + 1) % this.Capacity == this.head) {
		panic("队列已满，插入失败")
	}
	this.queue[this.tail] = i
	this.tail++
	this.tail = this.tail % this.Capacity
}

// OutputQueue - 出队
func (this *CircleQueue) OutputQueue() int {
	if (this.head % this.Capacity == this.tail) {
		panic("队列为空")
	}
	result := this.queue[this.head]
	this.head++
	this.head = this.head % this.Capacity
	return result
}

