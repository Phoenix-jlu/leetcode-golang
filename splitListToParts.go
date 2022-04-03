package main

func splitListToParts(head *ListNode, k int) []*ListNode {
	var listNodeSlic = make([]*ListNode,0)
	var ansSlic = make([]*ListNode,0)
	for (head != nil){
		tmpNode := head.Next
		head.Next = nil
		listNodeSlic = append(listNodeSlic,head)
		head = tmpNode
	}
	if len(listNodeSlic) / k < 1{
		for i:=0;i<k;i++{
			if i<len(listNodeSlic){
				ansSlic = append(ansSlic,listNodeSlic[i])
				continue
			}
			ansSlic = append(ansSlic,nil)
		}
		return ansSlic
	}
	a := len(listNodeSlic) / k
	b := len(listNodeSlic) % k
	lenSlic := make([]int,0)
	for i:=1;i<=k;i++{
		if b > 0 {
			lenSlic = append(lenSlic,a+1)
			b--
			continue
		}
		lenSlic = append(lenSlic,a)
	}
	tmp := 0
	for i:=0;i<len(lenSlic);i++ {
		ansSlic = append(ansSlic,listNodeSlic[tmp])
		for j:=0 + tmp;j<lenSlic[i]-1+tmp;j++{
			listNodeSlic[j].Next = listNodeSlic[j+1]
		}
		tmp += lenSlic[i]
	}
	return ansSlic
}
