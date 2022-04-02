package main

import "fmt"

type ListNode struct {
	     Val int
	     Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans = new(ListNode)
	var head = new(ListNode)
	tmpNum := 0
	if(l1 != nil && l2 != nil){
		if ((l1.Val + l2.Val) >= 10){
			ans.Val = l1.Val + l2.Val - 10
			tmpNum = 1
			l1 = l1.Next
			l2 = l2.Next
			head = ans
		}else {
			ans.Val = l1.Val + l2.Val
			l1 = l1.Next
			l2 = l2.Next
			head = ans
		}
	}
	for (l1 != nil || l2 != nil || tmpNum != 0){
		var tmp = new(ListNode)
		if ((l1 != nil) && (l2 != nil) &&(l1.Val + l2.Val + tmpNum) >= 10){
			tmp.Val = l1.Val + l2.Val + tmpNum - 10
			tmpNum = 1
			l1 = l1.Next
			l2 = l2.Next
			ans.Next = tmp
			ans = ans.Next
			continue
		}
		if ((l1 != nil) && (l2 != nil) &&(l1.Val + l2.Val + tmpNum) < 10){
			tmp.Val = l1.Val + l2.Val + tmpNum
			tmpNum = 0
			l1 = l1.Next
			l2 = l2.Next
			ans.Next = tmp
			ans = ans.Next
			continue
		}
		if ((l1 != nil) && (l2 == nil) && (l1.Val  + tmpNum) >= 10){
			tmp.Val = l1.Val  + tmpNum -10
			tmpNum = 1
			l1 = l1.Next
			ans.Next = tmp
			ans = ans.Next
			continue
		}
		if ((l1 != nil) && (l2 == nil) && (l1.Val  + tmpNum) < 10){
			tmp.Val = l1.Val  + tmpNum
			tmpNum = 0
			l1 = l1.Next
			ans.Next = tmp
			ans = ans.Next
			continue
		}
		if ((l1 == nil) && (l2 != nil) && (l2.Val  + tmpNum) >= 10){
			tmp.Val = l2.Val  + tmpNum -10
			tmpNum = 1
			l2 = l2.Next
			ans.Next = tmp
			ans = ans.Next
			continue
		}
		if ((l1 == nil) && (l2 != nil) && (l2.Val  + tmpNum) < 10){
			tmp.Val = l2.Val  + tmpNum
			tmpNum = 0
			l2 = l2.Next
			ans.Next = tmp
			ans = ans.Next
			continue
		}
		if ((l1 == nil) && (l2 == nil) && tmpNum != 0){
			tmp.Val = tmpNum
			tmpNum = 0
			ans.Next = tmp
			ans = ans.Next
			continue
		}
	}
	return head
}

func main(){
	//var l1  = new(ListNode)
	//var l2 = new(ListNode)
	//var l3 = new(ListNode)
	//var l4 = new(ListNode)
	//var l5 = new(ListNode)
	//var l6 = new(ListNode)
	//var l7 = new(ListNode)
	//var l9 = new(ListNode)
	//var l11 = new(ListNode)
	//var l13 = new(ListNode)
	//var l15 = new(ListNode)
	//var l8 = new(ListNode)
	//var ans = new(ListNode)
	//l1.Val = 9
	//l1.Next = l3
	//l3.Val = 9
	//l3.Next = l5
	//l5.Val = 9
	//l5.Next = l7
	//l7.Val = 9
	//l7.Next = l9
	//l9.Val = 9
	//l9.Next = l11
	//l11.Val = 9
	//l11.Next = l13
	//l13.Val = 9
	//l13.Next = l15
	//l15.Next = nil
	//l2.Val = 9
	//l2.Next = l4
	//l4.Val = 9
	//l4.Next = l6
	//l6.Val = 9
	//l6.Next = l8
	//l8.Val = 9
	//l8.Next = nil
	//ans = addTwoNumbers(l1,l2)
	//for (ans != nil){
	//	fmt.Println(ans.Val)
	//	ans = ans.Next
	//}
	var l1  = new(ListNode)
	var l2 = new(ListNode)
	var l3 = new(ListNode)
	var ans11 = new(ListNode)
	l1.Val = 1
	l1.Next = l3
	l3.Val = 1
	l3.Next = nil
	l2.Val = 0
	l2.Next = nil
	ans11 = addTwoNumbers(l1,l2)
	for (ans11 != nil){
			fmt.Println(ans11.Val)
			ans11 = ans11.Next
	}
}
