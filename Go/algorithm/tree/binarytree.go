package tree

import (
	"fmt"
	"container/list"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

// BuildTree - 构造一颗完全二叉树，返回根节点
func BuildTree(valSli []int) *Node {
	tempSli := make([]int, len(valSli), cap(valSli))
	copy(tempSli, valSli)
	// 创建队列
	q := NewQueue()

	// 创建根节点并入队
	root := &Node{}
	q.Enqueue(root)
	for len(tempSli) > 0 {
		currentNode := q.Outputqueue()
		currentNode.val = tempSli[0]
		tempSli = tempSli[1:]
		tempLeftNode := &Node{}
		tempRightNode := &Node{}
		currentNode.left = tempLeftNode
		currentNode.right = tempRightNode
		q.Enqueue(tempLeftNode)
		q.Enqueue(tempRightNode)
	}
	return root
}

// 深度优先遍历二叉树（前序）
func FindAll(root *Node) {
	fmt.Println(root.val, 1)
	if root.left != nil {
		FindAll(root.left)
	}
	if root.right != nil {
		FindAll(root.right)
	}
}

// 深度优先遍历二叉树（中序非递归）
func FindAllMiddle(root *Node) {
	stack := make([]*Node, 0)
	node := root

	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.left
		}

		node = stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		fmt.Println(node.val)
		node = node.right
	}

}

// 深度优先遍历二叉树（前序非递归）
func FrontFind(root *Node) {
    
    stack := make([]*Node, 0)
    
    stack = append(stack, root)
    for len(stack) != 0 {
        node := stack[len(stack) - 1]
       	fmt.Println(node.val)
        
        if node.left != nil {
            stack = append(stack, node.left)
        } else if node.right != nil {
            stack = stack[:len(stack) - 1]
            stack = append(stack, node.right)
        } else {
            stack = stack[:len(stack) - 1]
            
            for len(stack) != 0 {
				pre := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]
                if pre.right != nil {
					stack = append(stack, pre.right)
					break
                }
            }
        }
    }
}
func MaxDepth(root *Node) {
    if root == nil {
        return
    }
    if (root.left == nil) && (root.right == nil) {
        return
    }
    l := list.New()
    l.PushBack(root)
    outloop:
    for  {
        cruE := l.Back()
        if cruE == nil {
            break
        }
        cru := cruE.Value.(*Node)
        fmt.Println("节点", cru.val)
        if cru.left != nil {
            l.PushBack(cru.left)
        } else if cru.right != nil {
            l.PushBack(cru.right)
            l.Remove(cruE)
        } else {
            l.Remove(cruE)
            for {
                preE := l.Back()
                if preE == nil {
                    break outloop
                }
                pre := preE.Value.(*Node)
                l.Remove(preE)
                if pre.right != nil {
                    l.PushBack(pre.right)
                    break
                }
            }
        }
    }
}

// 广度遍历二叉树
func FindAllBreath(root *Node) {
	myQueue := NewQueue()
	myQueue.Enqueue(root)
	for myQueue.length != 0 {
		fmt.Println("长度", myQueue.length)
		currentNode := myQueue.Outputqueue()
		fmt.Println(currentNode.val)
		if currentNode.left != nil {
			myQueue.Enqueue(currentNode.left)
		}
		if currentNode.right != nil {
			myQueue.Enqueue(currentNode.right)
		}
	}
}

