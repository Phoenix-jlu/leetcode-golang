package main

import (
	"fmt"
	"strings"
	"unicode"
)

func shortestCompletingWord(licensePlate string, words []string) string {
	licensePlateTmp := strings.ToLower(licensePlate)
	var tmp = make(map[string]int)
	var ans = ""
	for i:=0;i<len(licensePlateTmp) ;i++ {
		if (!unicode.IsNumber(rune(licensePlateTmp[i])) && licensePlateTmp[i] != ' '){
			tmp[string(licensePlateTmp[i])]++
		}
	}
	for i:=0;i<len(words);i++ {
		var tmpnum =0
		for k,v := range tmp{
			if(strings.Count(words[i],k) >= v){
				tmpnum++
			}
		}
		if(len(tmp)==tmpnum && len(ans) == 0){
			ans = words[i]
			continue
		}
		if (len(tmp)==tmpnum && len(words[i]) < len(ans)) {
			ans = words[i]
		}
	}
	return ans
}

func main(){
	test := "AN87005"
	words := []string{"participant","individual","start","exist","above","already","easy","attack","player","important"}
	fmt.Println(shortestCompletingWord(test,words))
}
