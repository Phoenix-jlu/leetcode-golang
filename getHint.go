package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getHint(secret string, guess string) string {
	var bullsnum int //公牛
	var cowsnum int //奶牛
	var indexBullScli = make(map[int]int,0)
	secretScli := strings.Split(secret,"")
	guessScli := strings.Split(guess,"")
	for i:=0;i< len(secretScli);i++{
		if(secretScli[i] == guessScli[i]){
			bullsnum++
			tmp,_ := strconv.Atoi(secretScli[i])
			indexBullScli[i] = tmp
		}
	}
	for i:=0;i<len(guessScli);i++{
		//排除公牛已经在map中的情况
		if _,ok := indexBullScli[i];ok&&guessScli[i]==strconv.Itoa(indexBullScli[i]){
			continue
		}
		for j:=0;j<len(secretScli);j++ {
			//排除奶牛已经在map中的情况
			if _,ok := indexBullScli[j];ok{
				continue
			}
			if(guessScli[i] == secretScli[j]){
				cowsnum++
				tmp,_ := strconv.Atoi(secretScli[j])
				indexBullScli[j]=tmp
				break
			}
		}
	}
	return strconv.Itoa(bullsnum)+"A"+strconv.Itoa(cowsnum)+"B"
}

func main(){
	secret := "1"
	guess :=  "0"
	//"1807789"
	//"7810906"

	ans := getHint(secret,guess)
	fmt.Print(ans)
}
