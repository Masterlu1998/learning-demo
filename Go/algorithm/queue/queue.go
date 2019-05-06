package queue

type Queue struct {
	length int
	queue []int
}

// 新建一个队列
func NewQueue() *Queue {
	return &Queue{length: 0, queue: make([]int, 0)}
}

// Enqueue - 入队
func (this *Queue)Enqueue(i int) {
	this.queue = append(this.queue, i)
	this.length++
}

// Outputqueue - 出队
func (this *Queue)Outputqueue() int {
	result := this.queue[0]
	this.queue = this.queue[1:]
	return result
}