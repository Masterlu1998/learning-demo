package list

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    list := make([]*ListNode, 1)
    list[0] = head
    currentNode := head
    for currentNode.Next != nil {
        currentNode = currentNode.Next
        list = append(list, currentNode)
    }
    
    length := len(list)
    
    if length == 1 {
        return nil
    }
    
    if n == 1 {
        list[length - 2].Next = nil
    } else if length == n {
        return list[1]
    } else {
        deleteIndex := len(list) - n
        list[deleteIndex-1].Next = list[deleteIndex+1]
    }
    return list[0]
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    if head.Next == nil {
        return head
    }
    
    preNode := head
    currentNode := head.Next
    preNode.Next = nil
    for currentNode.Next != nil {
        tempNode := currentNode.Next
        currentNode.Next = preNode
        preNode = currentNode
        currentNode = tempNode
    }
    currentNode.Next = preNode
    return currentNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    cur1, cur2, cur := l1, l2, head
    for cur1 != nil && cur2 != nil {
        if cur1.Val < cur2.Val {
            newNode := &ListNode{cur1.Val, nil}
            cur.Next = newNode
            cur = newNode
            cur1 = cur1.Next
        } else {
            newNode := &ListNode{cur2.Val, nil}
            cur.Next = newNode
            cur = newNode 
            cur2 = cur2.Next
        }
    }
    
    for cur1 != nil {
            newNode := &ListNode{cur1.Val, nil}
            cur.Next = newNode
            cur = newNode
            cur1 = cur1.Next
    }
    
    for cur2 != nil {
            newNode := &ListNode{cur2.Val, nil}
            cur.Next = newNode
            cur = newNode
            cur2 = cur2.Next
    }
    return head.Next
}

func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    cur := head
    stack := make([]int, 0)
    for cur != nil {
        stack = append(stack, cur.Val)
        cur = cur.Next
    }
    
    cur = head
    i := len(stack) - 1
    for cur != nil {
        if cur.Val != stack[i] {
            return false
        }
        i--
        cur = cur.Next
    }
    return true
}

func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }
    fast, slow := head, head
    for {
        if fast.Next == nil || fast.Next.Next == nil || slow.Next == nil {
            break
        } else {
            fast = fast.Next.Next
            slow = slow.Next
            if fast == slow {
                return true
            }
        }
    }
    return false
}