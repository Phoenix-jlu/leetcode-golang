package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if (n == 1){
		return "1"
	}
	if(n == 2){
		return "11"
	}
	tmp := "11"
	if(n>2){
		for i:=2;i<n;i++{
			num := 1
			itoaS  := ""
			for j,k:=0,1;k<= len(tmp)-1; {
				if(tmp[j] == tmp[k] && k != len(tmp) -1 ){
					num++
					k++
				}
				if(tmp[j] != tmp[k]  && k != len(tmp) -1 ){
					itoaS = itoaS + strconv.Itoa(num) + string(tmp[j])
					j = k
					num = 1
					k++
				}
				if( tmp[j] == tmp[k]  && k == len(tmp) -1){
					num++
					itoaS = itoaS + strconv.Itoa(num) + string(tmp[j])
					break
				}
				if(tmp[j] != tmp[k] && k == len(tmp) -1){
					itoaS = itoaS + strconv.Itoa(num) + string(tmp[j])
					num = 1
					itoaS = itoaS + strconv.Itoa(num) + string(tmp[k])
					break
				}
			}
			tmp = itoaS
		}
	}
	return tmp
}

func main (){
	fmt.Println(countAndSay(5))
}