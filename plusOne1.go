package main

import "fmt"

func plusOne1(digits []int) []int {
	ans := make([]int,0)
	if(len(digits)== 1 && digits[0] == 0){
		return []int{1}
	}
	for i:=len(digits)-1;i>=0;i--{
		if(digits[i] != 9){
			digits[i] = digits[i] + 1
			return digits
		}
		if(digits[i] == 9 && i != 0){
			digits[i] = 0
			continue
		}
		if(digits[i] == 9 && i ==0){
			digits[i] = 0
			tmp := append([]int{},digits[0:]...)
			ans=append(digits[0:0],1)
			ans=append(ans,tmp...)
		}
	}
	return ans
}

func main(){
	//testInt := []int{7,2,8,5,0,9,1,2,9,5,3,6,6,7,3,2,8,4,3,7,9,5,7,7,4,7,4,9,4,7,0,1,1,1,7,9,9,9,9}
	testInt := []int{9}
	//testInt := []int{1,0}
	ans := plusOne1(testInt)
	fmt.Print(ans)
}