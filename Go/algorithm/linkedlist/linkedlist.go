package linkedlist

import (
	"fmt"
)

// Node - 节点结构
type Node struct {
	Val  int
	Next *Node
}

// 头结点
var headNode = Node{Val: 0, Next: nil}

// FindAll - 遍历整个链表
func FindAll() {
	nextNode := &headNode
	for nextNode.Next != nil {
		nextNode = nextNode.Next
		fmt.Println(nextNode.Val)
	}
}

// DeleteNode - 删除第index个节点
func DeleteNode(index int) {
	nextNode := &headNode
	for i := 0; i < index; i++ {
		nextNode = nextNode.Next
	}
	beforeNode := nextNode
	beforeNode.Next = beforeNode.Next.Next

}

// HeadAdd - 链表头节点之后添加节点
func HeadAdd(i *Node) {
	thirdNode := headNode.Next
	headNode.Next = i
	i.Next = thirdNode
}

// TailAdd - 在链表尾部添加结点
func TailAdd(i *Node) {
	nextNode := &headNode
	for nextNode.Next != nil {
		nextNode = nextNode.Next
	}
	finalNode := nextNode
	finalNode.Next = i
}

// MiddleAdd - 在第index个节点后添加一个节点
func MiddleAdd(i *Node, index int) {
	nextNode := &headNode
	for i := 0; i <= index; i++ {
		nextNode = nextNode.Next
	}
	beforeNode := nextNode
	tempNode := beforeNode.Next
	beforeNode.Next = i
	i.Next = tempNode
}
