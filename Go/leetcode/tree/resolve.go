package tree

import (
	"container/list"
	"fmt"
	"math"
)

 type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    level := 1
    levelSize := 1
    l := list.New()
    l.PushBack(root)
    for l.Len() != 0 {
        nodeE := l.Front()
        node := nodeE.Value.(*TreeNode)
        fmt.Println(node.Val)
        l.Remove(nodeE)
        if node.Left != nil {
            l.PushBack(node.Left) 
        }
        if node.Right != nil {
            l.PushBack(node.Right) 
        }

        levelSize--
        
        if levelSize == 0 {
            levelSize = l.Len()
            level++
        }  
    }
    return level-1
}


func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Left != nil && root.Right != nil {
        if root.Left.Val < root.Val && root.Val < root.Right.Val {
            return IsValid(root.Left, root.Val, math.MinInt64) && IsValid(root.Right, math.MaxInt64, root.Val)
        }
    } else if root.Left != nil {
        if root.Left.Val < root.Val {
            return IsValid(root.Left, root.Val, math.MinInt64)
        }
    } else if root.Right != nil {
        if root.Val < root.Right.Val {
            return IsValid(root.Right, math.MaxInt64, root.Val)
        }
    } else {
        return true
    }
    return false
}



func IsValid(node *TreeNode, max, min int) bool {
    if node.Val >= max || node.Val <= min {
        return false
    }
    if node.Left != nil && node.Right != nil {
        if node.Left.Val < node.Val && node.Val < node.Right.Val {
            return IsValid(node.Left, node.Val, min) && IsValid(node.Right, max,node.Val)
        }
    } else if node.Left != nil {
        if node.Left.Val < node.Val {
            return IsValid(node.Left, node.Val, min)
        }
    } else if node.Right != nil {
        if node.Val < node.Right.Val {
            return IsValid(node.Right, max,node.Val)
        }
    } else {
        return true
    }
    return false
}

func isSymmetric(root *TreeNode) bool {
    return _isSymmetric(root, root)
}

func _isSymmetric(node1 *TreeNode, node2 *TreeNode) bool {
    if node1 == nil && node2 == nil {
        return true
    } else if node1 == nil || node2 == nil {
        return false
    }
    
    return node1.Val == node2.Val && _isSymmetric(node1.Left, node2.Right) && _isSymmetric(node1.Right, node2.Left)
}


func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    
    result := make([][]int, 0)
    
    // 创建队列
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)
    
    // 初始化每行容量
    levelSize := 1
    levelResult := make([]int, 0)
    
    for len(queue) != 0 {
        // 出队
        node := queue[0]
        queue = queue[1:]
        levelResult = append(levelResult, node.Val)
        
        // 左右子节点入队
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        
        if node.Right != nil {
            queue = append(queue, node.Right)
        }

        levelSize--
        
        if levelSize == 0 {
            // 一层遍历完
            levelSize = len(queue)
            result = append(result, levelResult)
            levelResult = make([]int, 0)
        }
    }
    
    return result
}


func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    middle := (len(nums) - 1) / 2
    node := &TreeNode{Val: nums[middle]}
    node.Left = sortedArrayToBST(nums[:middle])
    node.Right = sortedArrayToBST(nums[middle+1:])
    return node
}