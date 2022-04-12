package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	strSlic := strings.Split(strs[0],"")
	var ans string
	for i:=0;i<len(strs[0]);i++{
		num := 0
		for j:=1;j<len(strs);j++{
			tmpSlic := strings.Split(strs[j],"")
			if (len(tmpSlic) <= i){
				break
			}
			if (strSlic[i] == tmpSlic[i]){
				num++
			}
		}
		if num == len(strs) -1{
			ans = ans + strSlic[i]
		}else {
			break
		}
	}
	return ans
}

func main(){
	var test = []string{"aaaaaaaaaaaaa","aaaaa","aaaaaaaa"}

	fmt.Println(longestCommonPrefix(test))
}