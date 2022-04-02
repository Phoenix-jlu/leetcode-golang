package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
//剑指 Offer 24. 反转链表
func reverseList(head *ListNode) *ListNode {
	var next = new(ListNode)
	var pre *ListNode = nil
	for(head != nil){
		next = head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}

func main()  {
	var l1  = new(ListNode)
	var l2 = new(ListNode)
	var l3 = new(ListNode)
	l1.Val = 1
	l1.Next = l2
	l2.Val = 2
	l2.Next = l3
	l3.Val = 3
	l3.Next = nil
	ans11 := reverseList(l1)
	for (ans11 != nil){
		fmt.Println(ans11.Val)
		ans11 = ans11.Next
	}

}