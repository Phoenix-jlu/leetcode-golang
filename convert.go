package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if (len(s) == 1 || numRows == 1){
		return s
	}
	var ans = ""
	sTmp := strings.Split(s,"")
	var num = numRows + numRows -2
	var tmp = make([][]string,numRows)
	for i:=0;i< len(s);i++  {
		if( (i % num) <= numRows-1){
			tmp[i%num] = append(tmp[i%num],sTmp[i])
			continue
		}
		tmp[num-i%num] = append(tmp[num-i%num],sTmp[i])
	}
	for i:=0;i<len(tmp);i++{
		ans = ans + strings.Join(tmp[i],"")
	}
	return ans
}

func main()  {
	//var test = "PAYPALISHIRING"
	var test = "AB"
	fmt.Println(convert(test,2))
}
