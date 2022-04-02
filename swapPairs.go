package main

import "fmt"

//两两交换链表中的节点
type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if(head == nil || head.Next == nil){
		return head
	}
	var newHead = new(ListNode)
	newHead = head.Next
	head.Next = swapPairs(newHead.Next)
	newHead.Next = head
	return newHead
}

func main()  {
	var l1  = new(ListNode)
	var l2 = new(ListNode)
	var l3 = new(ListNode)
	var l4 = new(ListNode)
	l1.Val = 1
	l1.Next = l2
	l2.Val = 2
	l2.Next = l3
	l3.Val = 3
	l3.Next = l4
	l4.Val = 4
	l4.Next = nil
	ans11 := swapPairs(l1)
	for (ans11 != nil){
		fmt.Println(ans11.Val)
		ans11 = ans11.Next
	}
}