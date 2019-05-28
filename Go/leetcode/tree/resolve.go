package tree

import (
	"container/list"
	"fmt"
	"reflect"
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
    if root == nil || (root.Left == nil && root.Right == nil) {
        return true
    }
    
    if root.Left == nil || root.Right == nil {
        return false
    }
    // 置换右子树
    switchTree(root.Right)
    return middleFind(root.Left, root.Right)
    
}

func switchTree(root *TreeNode) {
    root.Left, root.Right = root.Right, root.Left
    if root.Left != nil {
        switchTree(root.Left)
    }
    
    if root.Right != nil {
        switchTree(root.Right)
    }
}

func middleFind(root1, root2 *TreeNode) bool {
    
    stack1 := make([]*TreeNode, 0)
    stack2 := make([]*TreeNode, 0)
    
    node1 := root1
    node2 := root2
    
    
    for node1 != nil || node2 != nil || len(stack1) != 0 || len(stack2) != 0 {
        for node1 != nil {
            stack1 = append(stack1, node1)
            node1 = node1.Left  
        }
        
        for node2 != nil {
            stack2 = append(stack2, node2)
            node2 = node2.Left  
        }

        node1 = stack1[len(stack1) - 1]
        stack1 = stack1[:len(stack1) - 1]
        
        node2 = stack2[len(stack2) - 1]
        stack2 = stack2[:len(stack2) - 1]
        
        if !reflect.DeepEqual(node1, node2) || !reflect.DeepEqual(stack1, stack2) {
            return false
        }

        node1 = node1.Right
        node2 = node2.Right
        
    }
    return true
  
}
