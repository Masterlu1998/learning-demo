package tree

type Queue struct {
	length int
	queue  []*Node
}

// 新建一个队列
func NewQueue() *Queue {
	return &Queue{length: 0, queue: make([]*Node, 0)}
}

// Enqueue - 入队
func (this *Queue) Enqueue(i *Node) {
	this.queue = append(this.queue, i)
	this.length++
}

// Outputqueue - 出队
func (this *Queue) Outputqueue() *Node {
	result := this.queue[0]
	this.queue = this.queue[1:]
	this.length--
	return result
}
